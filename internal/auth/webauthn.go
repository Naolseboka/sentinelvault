package auth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/crypto"
)

var (
	ErrChallengeNotFound = errors.New("webauthn challenge not found")
	ErrChallengeExpired  = errors.New("webauthn challenge expired")
	ErrChallengeMismatch = errors.New("webauthn challenge mismatch")
)

// WebAuthnCredential stores a registered passkey credential.
type WebAuthnCredential struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	CredentialID []byte `json:"credential_id"`
	PublicKey  []byte `json:"public_key"`
	SignCount  uint32 `json:"sign_count"`
	CreatedAt  time.Time `json:"created_at"`
}

// WebAuthnChallenge represents an in-flight registration or assertion challenge.
type WebAuthnChallenge struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Challenge []byte    `json:"challenge"`
	Kind      string    `json:"kind"` // enroll | assert
	ExpiresAt time.Time `json:"expires_at"`
}

// ChallengeStore persists short-lived WebAuthn challenges.
type ChallengeStore interface {
	Save(ctx context.Context, ch *WebAuthnChallenge) error
	Consume(ctx context.Context, id, userID string) (*WebAuthnChallenge, error)
}

// MemoryChallengeStore is an in-memory ChallengeStore for development.
type MemoryChallengeStore struct {
	mu    sync.Mutex
	items map[string]*WebAuthnChallenge
}

// NewMemoryChallengeStore creates an empty challenge store.
func NewMemoryChallengeStore() *MemoryChallengeStore {
	return &MemoryChallengeStore{items: map[string]*WebAuthnChallenge{}}
}

func (s *MemoryChallengeStore) Save(_ context.Context, ch *WebAuthnChallenge) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	cp := *ch
	s.items[ch.ID] = &cp
	return nil
}

func (s *MemoryChallengeStore) Consume(_ context.Context, id, userID string) (*WebAuthnChallenge, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ch, ok := s.items[id]
	if !ok {
		return nil, ErrChallengeNotFound
	}
	delete(s.items, id)
	if ch.UserID != userID {
		return nil, ErrChallengeMismatch
	}
	if time.Now().UTC().After(ch.ExpiresAt) {
		return nil, ErrChallengeExpired
	}
	cp := *ch
	return &cp, nil
}

// NewWebAuthnChallenge generates a random challenge bound to a user.
func NewWebAuthnChallenge(userID, kind string, ttl time.Duration) (*WebAuthnChallenge, error) {
	raw, err := crypto.RandomBytes(32)
	if err != nil {
		return nil, err
	}
	if ttl <= 0 {
		ttl = 5 * time.Minute
	}
	return &WebAuthnChallenge{
		ID:        common.NewID(),
		UserID:    userID,
		Challenge: raw,
		Kind:      kind,
		ExpiresAt: time.Now().UTC().Add(ttl),
	}, nil
}

// VerifyAssertionBinding compares a client-provided challenge with the stored value.
func VerifyAssertionBinding(stored, provided []byte) error {
	if len(stored) == 0 || len(provided) == 0 {
		return ErrChallengeMismatch
	}
	a := sha256.Sum256(stored)
	b := sha256.Sum256(provided)
	for i := range a {
		if a[i] != b[i] {
			return ErrChallengeMismatch
		}
	}
	return nil
}

// EncodeChallenge returns a base64url challenge for clients.
func EncodeChallenge(ch []byte) string {
	return base64.RawURLEncoding.EncodeToString(ch)
}

// DecodeChallenge parses a base64url challenge from a client.
func DecodeChallenge(s string) ([]byte, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("decode challenge: %w", err)
	}
	return b, nil
}

// CredentialStore persists registered WebAuthn credentials.
type CredentialStore interface {
	Save(ctx context.Context, cred *WebAuthnCredential) error
	ListByUser(ctx context.Context, userID string) ([]WebAuthnCredential, error)
	GetByCredentialID(ctx context.Context, userID string, credentialID []byte) (*WebAuthnCredential, error)
	UpdateSignCount(ctx context.Context, id string, signCount uint32) error
}

// MemoryCredentialStore is an in-memory CredentialStore for development.
type MemoryCredentialStore struct {
	mu    sync.RWMutex
	byID  map[string]*WebAuthnCredential
	byUser map[string][]string // userID -> credential IDs
}

// NewMemoryCredentialStore creates an empty credential store.
func NewMemoryCredentialStore() *MemoryCredentialStore {
	return &MemoryCredentialStore{
		byID:   map[string]*WebAuthnCredential{},
		byUser: map[string][]string{},
	}
}

func (s *MemoryCredentialStore) Save(_ context.Context, cred *WebAuthnCredential) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	cp := *cred
	s.byID[cred.ID] = &cp
	found := false
	for _, id := range s.byUser[cred.UserID] {
		if id == cred.ID {
			found = true
			break
		}
	}
	if !found {
		s.byUser[cred.UserID] = append(s.byUser[cred.UserID], cred.ID)
	}
	return nil
}

func (s *MemoryCredentialStore) ListByUser(_ context.Context, userID string) ([]WebAuthnCredential, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ids := s.byUser[userID]
	out := make([]WebAuthnCredential, 0, len(ids))
	for _, id := range ids {
		if c, ok := s.byID[id]; ok {
			out = append(out, *c)
		}
	}
	return out, nil
}

func (s *MemoryCredentialStore) GetByCredentialID(_ context.Context, userID string, credentialID []byte) (*WebAuthnCredential, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, id := range s.byUser[userID] {
		c, ok := s.byID[id]
		if !ok {
			continue
		}
		if len(c.CredentialID) == len(credentialID) {
			match := true
			for i := range credentialID {
				if c.CredentialID[i] != credentialID[i] {
					match = false
					break
				}
			}
			if match {
				cp := *c
				return &cp, nil
			}
		}
	}
	return nil, ErrChallengeNotFound
}

func (s *MemoryCredentialStore) UpdateSignCount(_ context.Context, id string, signCount uint32) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	c, ok := s.byID[id]
	if !ok {
		return ErrChallengeNotFound
	}
	c.SignCount = signCount
	return nil
}

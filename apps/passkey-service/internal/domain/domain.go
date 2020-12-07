package domain

import (
	"context"
	"encoding/base64"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/auth"
	"github.com/sentinelvault/sentinelvault/internal/common"
)

// PasskeyRecord is an API view of a stored WebAuthn credential.
type PasskeyRecord struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	OrgID        string    `json:"org_id"`
	Name         string    `json:"name"`
	CredentialID string    `json:"credential_id"`
	PublicKey    string    `json:"public_key"`
	SignCount    uint32    `json:"sign_count"`
	CreatedAt    time.Time `json:"created_at"`
}

// CreateChallengeInput starts passkey registration or assertion.
type CreateChallengeInput struct {
	UserID string `json:"user_id"`
	OrgID  string `json:"org_id"`
	Kind   string `json:"kind"` // enroll | assert
}

// ChallengeResult returns challenge metadata for the client.
type ChallengeResult struct {
	ChallengeID string    `json:"challenge_id"`
	Challenge   string    `json:"challenge"`
	Kind        string    `json:"kind"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// RegisterInput finishes passkey registration.
type RegisterInput struct {
	UserID          string `json:"user_id"`
	OrgID           string `json:"org_id"`
	ChallengeID     string `json:"challenge_id"`
	ClientChallenge string `json:"client_challenge"`
	Name            string `json:"name"`
	CredentialID    string `json:"credential_id"`
	PublicKey       string `json:"public_key"`
}

// AssertInput completes a passkey authentication.
type AssertInput struct {
	UserID          string `json:"user_id"`
	ChallengeID     string `json:"challenge_id"`
	ClientChallenge string `json:"client_challenge"`
	CredentialID    string `json:"credential_id"`
	SignCount       uint32 `json:"sign_count"`
}

// AssertResult confirms successful passkey authentication.
type AssertResult struct {
	CredentialID string `json:"credential_id"`
	SignCount    uint32 `json:"sign_count"`
	Verified     bool   `json:"verified"`
}

// Service manages WebAuthn passkey credentials.
type Service struct {
	mu          sync.RWMutex
	names       map[string]string // credential ID -> display name
	orgIDs      map[string]string // credential ID -> org ID
	challenges  auth.ChallengeStore
	credentials auth.CredentialStore
}

// NewService creates a passkey domain service with in-memory stores.
func NewService() *Service {
	return &Service{
		names:       map[string]string{},
		orgIDs:      map[string]string{},
		challenges:  auth.NewMemoryChallengeStore(),
		credentials: auth.NewMemoryCredentialStore(),
	}
}

// CreateChallenge creates a registration or assertion challenge.
func (s *Service) CreateChallenge(ctx context.Context, in CreateChallengeInput) (*ChallengeResult, error) {
	if in.UserID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "user_id required")
	}
	kind := in.Kind
	if kind == "" {
		kind = "enroll"
	}
	if kind != "enroll" && kind != "assert" {
		return nil, common.New(common.ErrCodeInvalidArgument, "kind must be enroll or assert")
	}
	if kind == "assert" {
		creds, err := s.credentials.ListByUser(ctx, in.UserID)
		if err != nil {
			return nil, err
		}
		if len(creds) == 0 {
			return nil, common.New(common.ErrCodeNotFound, "no passkeys registered")
		}
	}
	ch, err := auth.NewWebAuthnChallenge(in.UserID, kind, 5*time.Minute)
	if err != nil {
		return nil, err
	}
	if err := s.challenges.Save(ctx, ch); err != nil {
		return nil, err
	}
	return &ChallengeResult{
		ChallengeID: ch.ID,
		Challenge:   auth.EncodeChallenge(ch.Challenge),
		Kind:        kind,
		ExpiresAt:   ch.ExpiresAt,
	}, nil
}

// Register verifies an enrollment challenge and stores the credential.
func (s *Service) Register(ctx context.Context, in RegisterInput) (*PasskeyRecord, error) {
	ch, err := s.challenges.Consume(ctx, in.ChallengeID, in.UserID)
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid challenge", err)
	}
	if ch.Kind != "enroll" {
		return nil, common.New(common.ErrCodeInvalidArgument, "challenge is not for enrollment")
	}
	if err := verifyClientChallenge(ch.Challenge, in.ClientChallenge); err != nil {
		return nil, err
	}
	pubKey, err := base64.RawURLEncoding.DecodeString(in.PublicKey)
	if err != nil {
		return nil, common.New(common.ErrCodeInvalidArgument, "invalid public_key encoding")
	}
	credID, err := base64.RawURLEncoding.DecodeString(in.CredentialID)
	if err != nil {
		return nil, common.New(common.ErrCodeInvalidArgument, "invalid credential_id encoding")
	}
	cred := &auth.WebAuthnCredential{
		ID:           common.NewID(),
		UserID:       in.UserID,
		CredentialID: credID,
		PublicKey:    pubKey,
		SignCount:    0,
		CreatedAt:    common.NowUTC(),
	}
	if err := s.credentials.Save(ctx, cred); err != nil {
		return nil, err
	}
	s.mu.Lock()
	s.names[cred.ID] = in.Name
	s.orgIDs[cred.ID] = in.OrgID
	s.mu.Unlock()
	return toPasskeyRecord(cred, in.Name, in.OrgID), nil
}

// Assert verifies an assertion challenge against a stored credential.
func (s *Service) Assert(ctx context.Context, in AssertInput) (*AssertResult, error) {
	ch, err := s.challenges.Consume(ctx, in.ChallengeID, in.UserID)
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid challenge", err)
	}
	if ch.Kind != "assert" {
		return nil, common.New(common.ErrCodeInvalidArgument, "challenge is not for assertion")
	}
	if err := verifyClientChallenge(ch.Challenge, in.ClientChallenge); err != nil {
		return nil, err
	}
	credID, err := base64.RawURLEncoding.DecodeString(in.CredentialID)
	if err != nil {
		return nil, common.New(common.ErrCodeInvalidArgument, "invalid credential_id encoding")
	}
	cred, err := s.credentials.GetByCredentialID(ctx, in.UserID, credID)
	if err != nil {
		return nil, common.Wrap(common.ErrCodeNotFound, "credential not found", err)
	}
	if in.SignCount > 0 && in.SignCount <= cred.SignCount {
		return nil, common.New(common.ErrCodeInvalidArgument, "sign count replay detected")
	}
	newCount := cred.SignCount
	if in.SignCount > cred.SignCount {
		newCount = in.SignCount
	} else {
		newCount++
	}
	if err := s.credentials.UpdateSignCount(ctx, cred.ID, newCount); err != nil {
		return nil, err
	}
	return &AssertResult{
		CredentialID: in.CredentialID,
		SignCount:    newCount,
		Verified:     true,
	}, nil
}

// List returns passkeys for a user.
func (s *Service) List(ctx context.Context, userID string) ([]PasskeyRecord, error) {
	creds, err := s.credentials.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]PasskeyRecord, 0, len(creds))
	for _, c := range creds {
		out = append(out, *toPasskeyRecord(&c, s.names[c.ID], s.orgIDs[c.ID]))
	}
	return out, nil
}

func verifyClientChallenge(stored []byte, clientEnc string) error {
	clientCh, err := auth.DecodeChallenge(clientEnc)
	if err != nil {
		return err
	}
	if err := auth.VerifyAssertionBinding(stored, clientCh); err != nil {
		return common.Wrap(common.ErrCodeInvalidArgument, "challenge mismatch", err)
	}
	return nil
}

func toPasskeyRecord(c *auth.WebAuthnCredential, name, orgID string) *PasskeyRecord {
	return &PasskeyRecord{
		ID:           c.ID,
		UserID:       c.UserID,
		OrgID:        orgID,
		Name:         name,
		CredentialID: base64.RawURLEncoding.EncodeToString(c.CredentialID),
		PublicKey:    base64.RawURLEncoding.EncodeToString(c.PublicKey),
		SignCount:    c.SignCount,
		CreatedAt:    c.CreatedAt,
	}
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "passkey" }

// TableName returns the primary table name.
func TableName() string { return "credentials" }

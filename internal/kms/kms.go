package kms

import (
	"context"
	"encoding/hex"
	"os"
	"sync"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/crypto"
)

// Key represents KMS encryption key metadata.
type Key struct {
	ID      string `json:"id"`
	OrgID   string `json:"org_id"`
	Name    string `json:"name"`
	Version int    `json:"version"`
	Status  string `json:"status"`
}

// WrapRequest wraps a DEK with a KMS key.
type WrapRequest struct {
	KeyID string `json:"key_id"`
	DEK   []byte `json:"dek"`
	AAD   []byte `json:"aad,omitempty"`
}

// WrapResponse holds wrapped DEK material.
type WrapResponse struct {
	WrappedDEK []byte `json:"wrapped_dek"`
	WrapNonce  []byte `json:"wrap_nonce"`
	KeyVersion int    `json:"key_version"`
}

// UnwrapRequest unwraps a DEK.
type UnwrapRequest struct {
	KeyID      string `json:"key_id"`
	WrappedDEK []byte `json:"wrapped_dek"`
	WrapNonce  []byte `json:"wrap_nonce"`
	AAD        []byte `json:"aad,omitempty"`
}

// UnwrapResponse holds the plaintext DEK.
type UnwrapResponse struct {
	DEK []byte `json:"dek"`
}

type keyMaterial struct {
	meta Key
	kek  []byte
}

// Registry is an in-memory KMS key registry with software master key.
type Registry struct {
	mu        sync.RWMutex
	masterKey []byte
	keys      map[string]*keyMaterial
}

// NewRegistry loads or generates the software master key.
func NewRegistry() (*Registry, error) {
	mk, err := loadMasterKey()
	if err != nil {
		return nil, err
	}
	return &Registry{
		masterKey: mk,
		keys:      map[string]*keyMaterial{},
	}, nil
}

// NewRegistryWithMasterKey creates a registry with an explicit master key (tests).
func NewRegistryWithMasterKey(masterKey []byte) (*Registry, error) {
	if len(masterKey) != crypto.AESKeySize {
		return nil, common.New(common.ErrCodeInvalidArgument, "master key must be 32 bytes")
	}
	cp := make([]byte, len(masterKey))
	copy(cp, masterKey)
	return &Registry{masterKey: cp, keys: map[string]*keyMaterial{}}, nil
}

func loadMasterKey() ([]byte, error) {
	hexKey := os.Getenv("KMS_MASTER_KEY")
	if hexKey != "" {
		key, err := hex.DecodeString(hexKey)
		if err != nil {
			return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid KMS_MASTER_KEY hex", err)
		}
		if len(key) != crypto.AESKeySize {
			return nil, common.New(common.ErrCodeInvalidArgument, "KMS_MASTER_KEY must decode to 32 bytes")
		}
		return key, nil
	}
	return crypto.RandomBytes(crypto.AESKeySize)
}

// Service implements software KMS operations.
type Service struct {
	reg *Registry
}

// NewService creates a KMS service backed by an in-memory registry.
func NewService(reg *Registry) *Service {
	return &Service{reg: reg}
}

// CreateKey registers a new org encryption key.
func (s *Service) CreateKey(ctx context.Context, orgID, name string) (*Key, error) {
	if err := common.ValidateRequired("name", name); err != nil {
		return nil, err
	}
	if err := common.ValidateRequired("org_id", orgID); err != nil {
		return nil, err
	}
	kek, err := crypto.GenerateDEK()
	if err != nil {
		return nil, err
	}
	key := Key{
		ID:      common.NewID(),
		OrgID:   orgID,
		Name:    name,
		Version: 1,
		Status:  "active",
	}
	s.reg.mu.Lock()
	defer s.reg.mu.Unlock()
	cp := make([]byte, len(kek))
	copy(cp, kek)
	s.reg.keys[key.ID] = &keyMaterial{meta: key, kek: cp}
	return &key, nil
}

// GetKey returns key metadata.
func (s *Service) GetKey(ctx context.Context, keyID string) (*Key, error) {
	s.reg.mu.RLock()
	defer s.reg.mu.RUnlock()
	km, ok := s.reg.keys[keyID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "key not found")
	}
	meta := km.meta
	return &meta, nil
}

// ListKeys returns keys for an org.
func (s *Service) ListKeys(ctx context.Context, orgID string) ([]Key, error) {
	s.reg.mu.RLock()
	defer s.reg.mu.RUnlock()
	out := make([]Key, 0)
	for _, km := range s.reg.keys {
		if km.meta.OrgID == orgID && km.meta.Status == "active" {
			out = append(out, km.meta)
		}
	}
	return out, nil
}

// WrapDEK encrypts a DEK under the named KMS key.
func (s *Service) WrapDEK(ctx context.Context, req WrapRequest) (*WrapResponse, error) {
	if len(req.DEK) != crypto.AESKeySize {
		return nil, common.New(common.ErrCodeInvalidArgument, "dek must be 32 bytes")
	}
	s.reg.mu.RLock()
	km, ok := s.reg.keys[req.KeyID]
	s.reg.mu.RUnlock()
	if !ok || km.meta.Status != "active" {
		return nil, common.New(common.ErrCodeNotFound, "key not found or inactive")
	}
	aead, err := crypto.NewAEAD(km.kek)
	if err != nil {
		return nil, err
	}
	wrapNonce, err := crypto.RandomBytes(aead.NonceSize())
	if err != nil {
		return nil, err
	}
	wrapped, err := aead.Encrypt(wrapNonce, req.DEK, req.AAD)
	if err != nil {
		return nil, err
	}
	return &WrapResponse{
		WrappedDEK: wrapped,
		WrapNonce:  wrapNonce,
		KeyVersion: km.meta.Version,
	}, nil
}

// UnwrapDEK decrypts a wrapped DEK.
func (s *Service) UnwrapDEK(ctx context.Context, req UnwrapRequest) (*UnwrapResponse, error) {
	s.reg.mu.RLock()
	km, ok := s.reg.keys[req.KeyID]
	s.reg.mu.RUnlock()
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "key not found")
	}
	aead, err := crypto.NewAEAD(km.kek)
	if err != nil {
		return nil, err
	}
	dek, err := aead.Decrypt(req.WrapNonce, req.WrappedDEK, req.AAD)
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "unwrap failed", err)
	}
	return &UnwrapResponse{DEK: dek}, nil
}

// RotateKey creates a new KEK version for an existing key id.
func (s *Service) RotateKey(ctx context.Context, keyID string) (*Key, error) {
	s.reg.mu.Lock()
	defer s.reg.mu.Unlock()
	km, ok := s.reg.keys[keyID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "key not found")
	}
	newKEK, err := crypto.GenerateDEK()
	if err != nil {
		return nil, err
	}
	km.meta.Version++
	copy(km.kek, newKEK)
	meta := km.meta
	return &meta, nil
}

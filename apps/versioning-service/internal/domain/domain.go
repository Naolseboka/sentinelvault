package domain

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/sentinelvault/sentinelvault/internal/common"
	encdomain "github.com/sentinelvault/sentinelvault/internal/encryption"
	"github.com/sentinelvault/sentinelvault/internal/vault"
)

// ItemRestorer applies a version snapshot back to live storage.
type ItemRestorer interface {
	RestoreFromSnapshot(ctx context.Context, snap *vault.VersionSnapshot) error
}

// Repository stores version snapshots.
type Repository interface {
	Save(ctx context.Context, snap *vault.VersionSnapshot) error
	Get(ctx context.Context, id string) (*vault.VersionSnapshot, error)
	ListByItem(ctx context.Context, itemID string) ([]vault.VersionSnapshot, error)
}

// Service manages vault item versions.
type Service struct {
	repo      Repository
	encryptor vault.Encryptor
	restorer  ItemRestorer
}

// NewService creates a versioning service.
func NewService(repo Repository, enc vault.Encryptor, restorer ItemRestorer) *Service {
	return &Service{repo: repo, encryptor: enc, restorer: restorer}
}

// RecordVersion persists an immutable snapshot.
func (s *Service) RecordVersion(ctx context.Context, snap *vault.VersionSnapshot) (*vault.VersionSnapshot, error) {
	if snap.ID == "" {
		snap.ID = common.NewID()
	}
	if snap.CreatedAt.IsZero() {
		snap.CreatedAt = common.NowUTC()
	}
	cp := *snap
	cp.Ciphertext = append([]byte(nil), snap.Ciphertext...)
	if err := s.repo.Save(ctx, &cp); err != nil {
		return nil, err
	}
	return &cp, nil
}

// Get returns a version snapshot metadata (ciphertext included).
func (s *Service) Get(ctx context.Context, orgID, versionID string) (*vault.VersionSnapshot, error) {
	snap, err := s.repo.Get(ctx, versionID)
	if err != nil {
		return nil, err
	}
	if snap == nil || snap.OrgID != orgID {
		return nil, common.New(common.ErrCodeNotFound, "version not found")
	}
	return snap, nil
}

// List returns versions for an item.
func (s *Service) List(ctx context.Context, orgID, itemID string) ([]vault.VersionSnapshot, error) {
	items, err := s.repo.ListByItem(ctx, itemID)
	if err != nil {
		return nil, err
	}
	out := make([]vault.VersionSnapshot, 0)
	for _, item := range items {
		if item.OrgID == orgID {
			out = append(out, item)
		}
	}
	return out, nil
}

// Restore applies a historical version to the live item via restorer.
func (s *Service) Restore(ctx context.Context, orgID, versionID string) (*vault.VersionSnapshot, error) {
	snap, err := s.Get(ctx, orgID, versionID)
	if err != nil {
		return nil, err
	}
	if s.restorer == nil {
		return nil, common.New(common.ErrCodeNotImplemented, "restore not configured")
	}
	if err := s.restorer.RestoreFromSnapshot(ctx, snap); err != nil {
		return nil, err
	}
	return snap, nil
}

// DecryptPayload decrypts a version snapshot payload for authorized callers.
func (s *Service) DecryptPayload(ctx context.Context, snap *vault.VersionSnapshot) (vault.ItemPayload, error) {
	decrypted, err := s.encryptor.EnvelopeDecrypt(ctx, encdomain.DecryptRequest{
		Ciphertext: snap.Ciphertext,
		AAD:        []byte(snap.OrgID),
	})
	if err != nil {
		return nil, err
	}
	var payload vault.ItemPayload
	if err := json.Unmarshal(decrypted.Plaintext, &payload); err != nil {
		return nil, err
	}
	return payload, nil
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "versioning" }

// TableName returns the primary table name.
func TableName() string { return "item_versions" }

// MemoryRepository stores versions in memory.
type MemoryRepository struct {
	mu   sync.RWMutex
	data map[string]*vault.VersionSnapshot
}

// NewMemoryRepository creates an empty repository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{data: map[string]*vault.VersionSnapshot{}}
}

func (r *MemoryRepository) Save(_ context.Context, snap *vault.VersionSnapshot) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *snap
	cp.Ciphertext = append([]byte(nil), snap.Ciphertext...)
	r.data[snap.ID] = &cp
	return nil
}

func (r *MemoryRepository) Get(_ context.Context, id string) (*vault.VersionSnapshot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	snap, ok := r.data[id]
	if !ok {
		return nil, nil
	}
	cp := *snap
	cp.Ciphertext = append([]byte(nil), snap.Ciphertext...)
	return &cp, nil
}

func (r *MemoryRepository) ListByItem(_ context.Context, itemID string) ([]vault.VersionSnapshot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]vault.VersionSnapshot, 0)
	for _, snap := range r.data {
		if snap.ItemID == itemID {
			out = append(out, *snap)
		}
	}
	return out, nil
}

package domain

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/backupx"
)

// Snapshot re-exports backupx.Snapshot.
type Snapshot = backupx.Snapshot

// BackupInput holds backup creation fields.
type BackupInput struct {
	OrgID   string `json:"org_id"`
	Label   string `json:"label"`
	Payload any    `json:"payload"`
}

// Service wraps the backup store.
type Service struct {
	store *backupx.Store
}

// NewService creates a backup domain service.
func NewService(store *backupx.Store) *Service {
	return &Service{store: store}
}

// NewStore creates a backup store with optional KEK.
func NewStore(kek []byte) (*backupx.Store, error) {
	return backupx.NewStore(kek)
}

// Backup compresses and encrypts a payload.
func (s *Service) Backup(_ context.Context, in BackupInput) (*Snapshot, error) {
	return s.store.Backup(in.OrgID, in.Label, in.Payload)
}

// Restore decrypts a snapshot into dest.
func (s *Service) Restore(_ context.Context, id string, dest any) (*Snapshot, error) {
	return s.store.Restore(id, dest)
}

// List returns snapshots for an org.
func (s *Service) List(_ context.Context, orgID string) []Snapshot {
	return s.store.List(orgID)
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "backup" }

// TableName returns the primary table name.
func TableName() string { return "snapshots" }

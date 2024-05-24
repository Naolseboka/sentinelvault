package domain

import (
	"context"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/sshx"
)

// KeyRecord re-exports sshx.KeyRecord.
type KeyRecord = sshx.KeyRecord

// SignedCert re-exports sshx.SignedCert.
type SignedCert = sshx.SignedCert

// GenerateKeyInput holds SSH key generation fields.
type GenerateKeyInput struct {
	OrgID string `json:"org_id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

// SignCertInput holds SSH certificate signing fields.
type SignCertInput struct {
	OrgID      string        `json:"org_id"`
	KeyID      string        `json:"key_id"`
	Principals []string      `json:"principals"`
	TTL        time.Duration `json:"ttl"`
}

// Service wraps the SSH manager.
type Service struct {
	manager *sshx.Manager
}

// NewService creates an SSH domain service.
func NewService(m *sshx.Manager) *Service {
	return &Service{manager: m}
}

// NewManager creates an SSH key manager.
func NewManager() (*sshx.Manager, error) {
	return sshx.NewManager()
}

// GenerateKey creates a new SSH keypair.
func (s *Service) GenerateKey(_ context.Context, in GenerateKeyInput) (*KeyRecord, error) {
	return s.manager.GenerateKey(in.OrgID, in.Name, in.Type)
}

// RotateKey replaces key material for an existing key.
func (s *Service) RotateKey(_ context.Context, id string) (*KeyRecord, error) {
	return s.manager.RotateKey(id)
}

// SignUserCert issues an SSH user certificate.
func (s *Service) SignUserCert(_ context.Context, in SignCertInput) (*SignedCert, error) {
	return s.manager.SignUserCert(in.OrgID, in.KeyID, in.Principals, in.TTL)
}

// GetKey returns key metadata.
func (s *Service) GetKey(_ context.Context, id string) (*KeyRecord, error) {
	return s.manager.GetKey(id, false)
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "ssh" }

// TableName returns the primary table name.
func TableName() string { return "keys" }

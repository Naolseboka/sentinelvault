package domain

import "github.com/sentinelvault/sentinelvault/internal/kms"

// Service wraps the shared software KMS.
type Service = kms.Service

// Key re-exports KMS key type.
type Key = kms.Key

// WrapRequest re-exports KMS wrap request.
type WrapRequest = kms.WrapRequest

// WrapResponse re-exports KMS wrap response.
type WrapResponse = kms.WrapResponse

// UnwrapRequest re-exports KMS unwrap request.
type UnwrapRequest = kms.UnwrapRequest

// UnwrapResponse re-exports KMS unwrap response.
type UnwrapResponse = kms.UnwrapResponse

// NewRegistry creates a KMS registry.
func NewRegistry() (*kms.Registry, error) { return kms.NewRegistry() }

// NewRegistryWithMasterKey creates a test registry.
func NewRegistryWithMasterKey(masterKey []byte) (*kms.Registry, error) {
	return kms.NewRegistryWithMasterKey(masterKey)
}

// NewService creates a domain service.
func NewService(reg *kms.Registry) *Service { return kms.NewService(reg) }

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "kms" }

// TableName returns the primary table name.
func TableName() string { return "keys" }

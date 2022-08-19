package domain

import "github.com/sentinelvault/sentinelvault/internal/encryption"

// Service wraps shared envelope encryption.
type Service = encryption.Service

// EncryptRequest re-exports encryption request type.
type EncryptRequest = encryption.EncryptRequest

// EncryptResponse re-exports encryption response type.
type EncryptResponse = encryption.EncryptResponse

// DecryptRequest re-exports decrypt request type.
type DecryptRequest = encryption.DecryptRequest

// DecryptResponse re-exports decrypt response type.
type DecryptResponse = encryption.DecryptResponse

// InProcessKMS re-exports in-process KMS adapter.
type InProcessKMS = encryption.InProcessKMS

// InProcessEncryptor re-exports encryptor adapter.
type InProcessEncryptor = encryption.InProcessEncryptor

// NewService creates an encryption domain service.
func NewService(kms encryption.KMSClient) *Service {
	return encryption.NewService(kms)
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "encryption" }

// TableName returns the primary table name.
func TableName() string { return "operations" }

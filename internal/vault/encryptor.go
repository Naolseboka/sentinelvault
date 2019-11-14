package vault

import (
	"context"

	enc "github.com/sentinelvault/sentinelvault/internal/encryption"
)

// Encryptor encrypts and decrypts vault payloads.
type Encryptor interface {
	EnvelopeEncrypt(ctx context.Context, req enc.EncryptRequest) (*enc.EncryptResponse, error)
	EnvelopeDecrypt(ctx context.Context, req enc.DecryptRequest) (*enc.DecryptResponse, error)
}

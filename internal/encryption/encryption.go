package encryption

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/crypto"
	"github.com/sentinelvault/sentinelvault/internal/kms"
)

const blobVersion byte = 1

// KMSClient wraps/unwarps DEKs via KMS.
type KMSClient interface {
	WrapDEK(ctx context.Context, req kms.WrapRequest) (*kms.WrapResponse, error)
	UnwrapDEK(ctx context.Context, req kms.UnwrapRequest) (*kms.UnwrapResponse, error)
}

// InProcessKMS adapts a KMS service for in-process use.
type InProcessKMS struct {
	Svc *kms.Service
}

func (c *InProcessKMS) WrapDEK(ctx context.Context, req kms.WrapRequest) (*kms.WrapResponse, error) {
	return c.Svc.WrapDEK(ctx, req)
}

func (c *InProcessKMS) UnwrapDEK(ctx context.Context, req kms.UnwrapRequest) (*kms.UnwrapResponse, error) {
	return c.Svc.UnwrapDEK(ctx, req)
}

// Service performs envelope encryption using KMS for DEK wrapping.
type Service struct {
	kms KMSClient
}

// NewService creates an encryption service.
func NewService(kmsClient KMSClient) *Service {
	return &Service{kms: kmsClient}
}

// EncryptRequest holds plaintext input.
type EncryptRequest struct {
	Plaintext []byte `json:"plaintext"`
	KeyID     string `json:"key_id"`
	AAD       []byte `json:"aad,omitempty"`
}

// EncryptResponse returns a serialized ciphertext blob.
type EncryptResponse struct {
	Ciphertext []byte `json:"ciphertext"`
	KeyID      string `json:"key_id"`
}

// DecryptRequest holds ciphertext input.
type DecryptRequest struct {
	Ciphertext []byte `json:"ciphertext"`
	AAD        []byte `json:"aad,omitempty"`
}

// DecryptResponse holds recovered plaintext.
type DecryptResponse struct {
	Plaintext []byte `json:"plaintext"`
}

// EnvelopeEncrypt generates a DEK, encrypts plaintext, wraps DEK via KMS.
func (s *Service) EnvelopeEncrypt(ctx context.Context, req EncryptRequest) (*EncryptResponse, error) {
	if err := common.ValidateRequired("key_id", req.KeyID); err != nil {
		return nil, err
	}
	if len(req.Plaintext) == 0 {
		return nil, common.New(common.ErrCodeInvalidArgument, "plaintext required")
	}
	dek, err := crypto.GenerateDEK()
	if err != nil {
		return nil, err
	}
	aead, err := crypto.NewAEAD(dek)
	if err != nil {
		return nil, err
	}
	nonce, err := crypto.RandomBytes(aead.NonceSize())
	if err != nil {
		return nil, err
	}
	ciphertext, err := aead.Encrypt(nonce, req.Plaintext, req.AAD)
	if err != nil {
		return nil, err
	}
	wrapped, err := s.kms.WrapDEK(ctx, kms.WrapRequest{
		KeyID: req.KeyID,
		DEK:   dek,
		AAD:   req.AAD,
	})
	if err != nil {
		return nil, err
	}
	blob, err := crypto.MarshalBlob(&crypto.CiphertextBlob{
		Version:    blobVersion,
		KeyID:      req.KeyID,
		Nonce:      nonce,
		WrapNonce:  wrapped.WrapNonce,
		WrappedDEK: wrapped.WrappedDEK,
		Ciphertext: ciphertext,
	})
	if err != nil {
		return nil, err
	}
	return &EncryptResponse{Ciphertext: blob, KeyID: req.KeyID}, nil
}

// EnvelopeDecrypt unwraps DEK via KMS and decrypts ciphertext blob.
func (s *Service) EnvelopeDecrypt(ctx context.Context, req DecryptRequest) (*DecryptResponse, error) {
	blob, err := crypto.UnmarshalBlob(req.Ciphertext)
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid ciphertext blob", err)
	}
	unwrapped, err := s.kms.UnwrapDEK(ctx, kms.UnwrapRequest{
		KeyID:      blob.KeyID,
		WrappedDEK: blob.WrappedDEK,
		WrapNonce:  blob.WrapNonce,
		AAD:        req.AAD,
	})
	if err != nil {
		return nil, err
	}
	aead, err := crypto.NewAEAD(unwrapped.DEK)
	if err != nil {
		return nil, err
	}
	plain, err := aead.Decrypt(blob.Nonce, blob.Ciphertext, req.AAD)
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "decrypt failed", err)
	}
	return &DecryptResponse{Plaintext: plain}, nil
}

// InProcessEncryptor adapts this service for vault consumers.
type InProcessEncryptor struct {
	Svc *Service
}

func (e *InProcessEncryptor) EnvelopeEncrypt(ctx context.Context, req EncryptRequest) (*EncryptResponse, error) {
	return e.Svc.EnvelopeEncrypt(ctx, req)
}

func (e *InProcessEncryptor) EnvelopeDecrypt(ctx context.Context, req DecryptRequest) (*DecryptResponse, error) {
	return e.Svc.EnvelopeDecrypt(ctx, req)
}

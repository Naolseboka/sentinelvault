package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

const AESKeySize = 32

var (
	ErrInvalidKeySize  = errors.New("invalid AES key size")
	ErrInvalidNonce    = errors.New("invalid nonce size")
	ErrCiphertextShort = errors.New("ciphertext too short")
)

// AEAD wraps AES-256-GCM encrypt/decrypt operations.
type AEAD struct {
	gcm cipher.AEAD
}

// NewAEAD creates an AES-256-GCM instance from a 32-byte key.
func NewAEAD(key []byte) (*AEAD, error) {
	if len(key) != AESKeySize {
		return nil, fmt.Errorf("%w: got %d", ErrInvalidKeySize, len(key))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &AEAD{gcm: gcm}, nil
}

// NonceSize returns the required nonce size for GCM.
func (a *AEAD) NonceSize() int { return a.gcm.NonceSize() }

// Overhead returns authentication tag size.
func (a *AEAD) Overhead() int { return a.gcm.Overhead() }

// Encrypt seals plaintext with optional associated data.
func (a *AEAD) Encrypt(nonce, plaintext, additionalData []byte) ([]byte, error) {
	if len(nonce) != a.gcm.NonceSize() {
		return nil, fmt.Errorf("%w: expected %d", ErrInvalidNonce, a.gcm.NonceSize())
	}
	return a.gcm.Seal(nil, nonce, plaintext, additionalData), nil
}

// Decrypt opens ciphertext with optional associated data.
func (a *AEAD) Decrypt(nonce, ciphertext, additionalData []byte) ([]byte, error) {
	if len(nonce) != a.gcm.NonceSize() {
		return nil, fmt.Errorf("%w: expected %d", ErrInvalidNonce, a.gcm.NonceSize())
	}
	plain, err := a.gcm.Open(nil, nonce, ciphertext, additionalData)
	if err != nil {
		return nil, err
	}
	return plain, nil
}

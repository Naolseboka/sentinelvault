package crypto

import (
	"crypto/ed25519"
	"errors"
)

var (
	ErrInvalidSignature = errors.New("invalid ed25519 signature")
	ErrInvalidKeyLength = errors.New("invalid ed25519 key length")
)

// GenerateEd25519KeyPair creates a new signing key pair.
func GenerateEd25519KeyPair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(RandomReader())
}

// SignEd25519 signs message with a private key.
func SignEd25519(privateKey ed25519.PrivateKey, message []byte) ([]byte, error) {
	if len(privateKey) != ed25519.PrivateKeySize {
		return nil, ErrInvalidKeyLength
	}
	return ed25519.Sign(privateKey, message), nil
}

// VerifyEd25519 verifies a signature against a public key.
func VerifyEd25519(publicKey ed25519.PublicKey, message, signature []byte) error {
	if len(publicKey) != ed25519.PublicKeySize {
		return ErrInvalidKeyLength
	}
	if !ed25519.Verify(publicKey, message, signature) {
		return ErrInvalidSignature
	}
	return nil
}

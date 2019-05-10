package crypto

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/hkdf"
)

// DeriveKey performs HKDF-SHA256 key derivation.
func DeriveKey(secret, salt, info []byte, length int) ([]byte, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be positive")
	}
	reader := hkdf.New(sha256.New, secret, salt, info)
	out := make([]byte, length)
	if _, err := reader.Read(out); err != nil {
		return nil, err
	}
	return out, nil
}

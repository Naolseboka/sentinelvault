package auth

import (
	"github.com/sentinelvault/sentinelvault/internal/crypto"
)

// HashPassword derives an Argon2id hash suitable for credential storage.
func HashPassword(password string) (string, error) {
	return crypto.HashPassword(password)
}

// VerifyPassword compares a plaintext password against a stored hash.
func VerifyPassword(password, encoded string) error {
	return crypto.VerifyPassword(password, encoded)
}

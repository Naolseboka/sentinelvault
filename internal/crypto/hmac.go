package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
)

// HMACSHA256 computes HMAC-SHA256(key, message).
func HMACSHA256(key, message []byte) ([]byte, error) {
	m := hmac.New(sha256.New, key)
	if _, err := m.Write(message); err != nil {
		return nil, err
	}
	return m.Sum(nil), nil
}

// MustSHA256 returns SHA-256 digest (never fails).
func MustSHA256(data []byte) []byte {
	sum := sha256.Sum256(data)
	out := make([]byte, len(sum))
	copy(out, sum[:])
	return out
}

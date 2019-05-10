package crypto

import (
	"crypto/rand"
	"io"
)

// RandomReader returns the process CSPRNG.
func RandomReader() io.Reader { return rand.Reader }

// RandomBytes reads n cryptographically secure random bytes.
func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return nil, err
	}
	return b, nil
}

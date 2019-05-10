package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256Hex returns the hex-encoded SHA-256 digest of data.
func SHA256Hex(data []byte) (string, error) {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

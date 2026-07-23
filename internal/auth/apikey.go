package auth

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/crypto"
)

const apiKeyPrefix = "sv_"

// APIKeyMaterial holds a generated API key and its storage hash.
type APIKeyMaterial struct {
	Key    string `json:"key"`
	KeyID  string `json:"key_id"`
	Hash   string `json:"hash"`
	Prefix string `json:"prefix"`
}

// GenerateAPIKey creates a new API key with sv_ prefix and hash for storage.
func GenerateAPIKey() (*APIKeyMaterial, error) {
	raw, err := crypto.RandomBytes(32)
	if err != nil {
		return nil, err
	}
	key := apiKeyPrefix + base64.RawURLEncoding.EncodeToString(raw)
	hash, err := HashAPIKey(key)
	if err != nil {
		return nil, err
	}
	prefix := key
	if len(prefix) > 12 {
		prefix = prefix[:12]
	}
	return &APIKeyMaterial{
		Key:    key,
		KeyID:  hex.EncodeToString(raw[:8]),
		Hash:   hash,
		Prefix: prefix,
	}, nil
}

// HashAPIKey returns a SHA-256 hex digest of the API key for storage.
func HashAPIKey(key string) (string, error) {
	if !strings.HasPrefix(key, apiKeyPrefix) {
		return "", fmt.Errorf("invalid api key prefix")
	}
	sum := sha256.Sum256([]byte(key))
	return hex.EncodeToString(sum[:]), nil
}

// VerifyAPIKey compares a presented key against a stored hash.
func VerifyAPIKey(key, storedHash string) (bool, error) {
	hash, err := HashAPIKey(key)
	if err != nil {
		return false, err
	}
	if subtle.ConstantTimeCompare([]byte(hash), []byte(storedHash)) == 1 {
		return true, nil
	}
	return false, nil
}

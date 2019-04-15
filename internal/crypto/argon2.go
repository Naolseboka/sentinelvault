package crypto

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	argonTime    = 3
	argonMemory  = 64 * 1024
	argonThreads = 4
	argonKeyLen  = 32
	argonSaltLen = 16
)

var ErrPasswordMismatch = errors.New("password mismatch")

// HashPassword derives an Argon2id hash encoded as base64(salt)$base64(hash).
func HashPassword(password string) (string, error) {
	salt, err := RandomBytes(argonSaltLen)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, argonTime, argonMemory, argonThreads, argonKeyLen)
	return fmt.Sprintf("%s$%s",
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	), nil
}

// VerifyPassword compares a password against an Argon2id hash.
func VerifyPassword(password, encoded string) error {
	parts := strings.Split(encoded, "$")
	if len(parts) != 2 {
		return fmt.Errorf("invalid hash format")
	}
	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return err
	}
	expected, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return err
	}
	actual := argon2.IDKey([]byte(password), salt, argonTime, argonMemory, argonThreads, uint32(len(expected)))
	if subtle.ConstantTimeCompare(actual, expected) != 1 {
		return ErrPasswordMismatch
	}
	return nil
}

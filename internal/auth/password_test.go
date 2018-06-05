package auth_test

import (
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/auth"
)

func TestHashAndVerifyPassword(t *testing.T) {
	t.Parallel()
	hash, err := auth.HashPassword("correct horse battery staple")
	if err != nil {
		t.Fatal(err)
	}
	if err := auth.VerifyPassword("correct horse battery staple", hash); err != nil {
		t.Fatalf("verify: %v", err)
	}
	if err := auth.VerifyPassword("wrong", hash); err == nil {
		t.Fatal("expected mismatch")
	}
}

func TestVerifyPasswordInvalidHash(t *testing.T) {
	t.Parallel()
	if err := auth.VerifyPassword("pw", "not-a-valid-hash"); err == nil {
		t.Fatal("expected error for invalid hash")
	}
}

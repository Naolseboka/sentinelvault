package auth_test

import (
	"strings"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/auth"
)

func TestGenerateVerifyAPIKey(t *testing.T) {
	t.Parallel()
	mat, err := auth.GenerateAPIKey()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(mat.Key, "sv_") {
		t.Fatalf("prefix = %q", mat.Key[:4])
	}
	ok, err := auth.VerifyAPIKey(mat.Key, mat.Hash)
	if err != nil || !ok {
		t.Fatalf("verify own key: ok=%v err=%v", ok, err)
	}
	ok, err = auth.VerifyAPIKey("sv_invalid", mat.Hash)
	if err != nil || ok {
		t.Fatalf("verify wrong key: ok=%v err=%v", ok, err)
	}
}

func TestHashAPIKeyInvalidPrefix(t *testing.T) {
	t.Parallel()
	if _, err := auth.HashAPIKey("bad_prefix"); err == nil {
		t.Fatal("expected invalid prefix error")
	}
}

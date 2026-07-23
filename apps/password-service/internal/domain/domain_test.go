package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"

	encdomain "github.com/sentinelvault/sentinelvault/internal/encryption"
	kmsdomain "github.com/sentinelvault/sentinelvault/internal/kms"
	"github.com/sentinelvault/sentinelvault/internal/vault"
)

func setupPasswordService(t *testing.T) (*Service, *MemoryRepository, string) {
	t.Helper()
	reg, err := kmsdomain.NewRegistryWithMasterKey(bytes.Repeat([]byte{0x42}, 32))
	if err != nil {
		t.Fatal(err)
	}
	kmsSvc := kmsdomain.NewService(reg)
	ctx := context.Background()
	key, err := kmsSvc.CreateKey(ctx, "org-1", "vault-key")
	if err != nil {
		t.Fatal(err)
	}
	encSvc := encdomain.NewService(&encdomain.InProcessKMS{Svc: kmsSvc})
	repo := NewMemoryRepository()
	svc := NewService(repo, &encdomain.InProcessEncryptor{Svc: encSvc})
	return svc, repo, key.ID
}

func TestCreateStoresCiphertextNotPlaintext(t *testing.T) {
	svc, repo, keyID := setupPasswordService(t)
	ctx := context.Background()
	secret := "super-secret-password"
	view, err := svc.Create(ctx, CreateRequest{
		OrgID:    "org-1",
		KeyID:    keyID,
		Type:     vault.ItemTypePassword,
		Name:     "Example Login",
		Payload:  vault.ItemPayload{"password": secret, "username": "alice"},
		Metadata: vault.ItemMetadata{URL: "https://example.com"},
	})
	if err != nil {
		t.Fatal(err)
	}
	stored, err := repo.Get(ctx, view.ID)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(stored.Ciphertext, []byte(secret)) {
		t.Fatal("plaintext password found in stored ciphertext")
	}
	raw, _ := json.Marshal(stored)
	if strings.Contains(string(raw), secret) {
		t.Fatal("plaintext password found in stored record json")
	}
	got, err := svc.Get(ctx, "org-1", view.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Payload["password"] != secret {
		t.Fatal("decrypted payload mismatch")
	}
}

func TestUpdateVersionsItem(t *testing.T) {
	svc, _, keyID := setupPasswordService(t)
	ctx := context.Background()
	view, err := svc.Create(ctx, CreateRequest{
		OrgID:   "org-1",
		KeyID:   keyID,
		Type:    vault.ItemTypeSecureNote,
		Name:    "Note",
		Payload: vault.ItemPayload{"note": "v1"},
	})
	if err != nil {
		t.Fatal(err)
	}
	updated, err := svc.Update(ctx, view.ID, UpdateRequest{
		OrgID:   "org-1",
		Payload: vault.ItemPayload{"note": "v2"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if updated.Version != 2 {
		t.Fatalf("expected version 2, got %d", updated.Version)
	}
}

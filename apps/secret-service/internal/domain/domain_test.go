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

func TestSecretCreateNoPlaintextAtRest(t *testing.T) {
	reg, _ := kmsdomain.NewRegistryWithMasterKey(bytes.Repeat([]byte{0x99}, 32))
	kmsSvc := kmsdomain.NewService(reg)
	ctx := context.Background()
	key, _ := kmsSvc.CreateKey(ctx, "org-1", "secret-key")
	encSvc := encdomain.NewService(&encdomain.InProcessKMS{Svc: kmsSvc})
	repo := NewMemoryRepository()
	svc := NewService(repo, &encdomain.InProcessEncryptor{Svc: encSvc})

	token := "sk-live-abc123"
	view, err := svc.Create(ctx, CreateRequest{
		OrgID:   "org-1",
		KeyID:   key.ID,
		Type:    vault.ItemTypeAPIKey,
		Name:    "Stripe",
		Payload: vault.ItemPayload{"api_key": token},
	})
	if err != nil {
		t.Fatal(err)
	}
	stored, _ := repo.Get(ctx, view.ID)
	if bytes.Contains(stored.Ciphertext, []byte(token)) {
		t.Fatal("token found in ciphertext")
	}
	raw, _ := json.Marshal(stored)
	if strings.Contains(string(raw), token) {
		t.Fatal("token found in stored json")
	}
}

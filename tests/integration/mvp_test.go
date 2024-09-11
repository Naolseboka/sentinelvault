package integration_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/audit"
	"github.com/sentinelvault/sentinelvault/internal/auth"
	encdomain "github.com/sentinelvault/sentinelvault/internal/encryption"
	kmsdomain "github.com/sentinelvault/sentinelvault/internal/kms"
	"github.com/sentinelvault/sentinelvault/internal/risk"
	"github.com/sentinelvault/sentinelvault/internal/vault"
)

// TestIdentityVaultAuditLoop covers Phase 2–4 core shared libraries end-to-end.
func TestIdentityVaultAuditLoop(t *testing.T) {
	ctx := context.Background()

	hash, err := auth.HashPassword("Str0ng!Pass#2026")
	if err != nil {
		t.Fatal(err)
	}
	if err := auth.VerifyPassword("Str0ng!Pass#2026", hash); err != nil {
		t.Fatal(err)
	}

	tm := auth.NewTokenManager([]byte("integration-jwt-secret-32bytes!!"), 15*time.Minute)
	access, err := tm.IssueAccessToken(auth.Claims{
		Subject: "user-1", OrgID: "org-1", SessionID: "sess-1", Roles: []string{"admin"}, MFAVerified: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	claims, err := tm.ParseAccessToken(access)
	if err != nil || string(claims.Subject) != "user-1" {
		t.Fatalf("claims=%+v err=%v", claims, err)
	}
	if _, err := auth.NewRefreshTokenID(); err != nil {
		t.Fatal(err)
	}

	score := risk.NewEngine().Evaluate(risk.Factors{NewDevice: true, DeviceTrust: "pending"})
	if score.Value <= 0 {
		t.Fatal("expected risk score")
	}

	reg, err := kmsdomain.NewRegistryWithMasterKey(bytes.Repeat([]byte{9}, 32))
	if err != nil {
		t.Fatal(err)
	}
	kmsSvc := kmsdomain.NewService(reg)
	key, err := kmsSvc.CreateKey(ctx, "org-1", "vault")
	if err != nil {
		t.Fatal(err)
	}
	encSvc := encdomain.NewService(&encdomain.InProcessKMS{Svc: kmsSvc})
	plain := []byte(`{"password":"s3cret"}`)
	enc, err := encSvc.EnvelopeEncrypt(ctx, encdomain.EncryptRequest{KeyID: key.ID, Plaintext: plain})
	if err != nil {
		t.Fatal(err)
	}
	dec, err := encSvc.EnvelopeDecrypt(ctx, encdomain.DecryptRequest{Ciphertext: enc.Ciphertext})
	if err != nil {
		t.Fatal(err)
	}
	if string(dec.Plaintext) != string(plain) {
		t.Fatal("roundtrip mismatch")
	}
	if err := vault.ValidatePayload(vault.ItemTypePassword, vault.ItemPayload{"password": "s3cret", "username": "u"}); err != nil {
		t.Fatal(err)
	}

	chain, err := audit.NewChain()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := chain.Append("org-1", "user-1", "vault.item.create", "password", "item-1", "info", nil); err != nil {
		t.Fatal(err)
	}
	if err := chain.VerifyIntegrity(); err != nil {
		t.Fatal(err)
	}
}

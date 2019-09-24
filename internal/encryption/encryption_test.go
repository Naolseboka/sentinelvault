package encryption

import (
	"bytes"
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/kms"
)

func TestEnvelopeEncryptDecrypt(t *testing.T) {
	reg, err := kms.NewRegistryWithMasterKey(bytes.Repeat([]byte{0x11}, 32))
	if err != nil {
		t.Fatal(err)
	}
	kmsSvc := kms.NewService(reg)
	ctx := context.Background()
	key, err := kmsSvc.CreateKey(ctx, "org-1", "enc-key")
	if err != nil {
		t.Fatal(err)
	}

	enc := NewService(&InProcessKMS{Svc: kmsSvc})
	plain := []byte(`{"password":"hunter2"}`)
	aad := []byte("item-aad")
	encrypted, err := enc.EnvelopeEncrypt(ctx, EncryptRequest{
		Plaintext: plain,
		KeyID:     key.ID,
		AAD:       aad,
	})
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(encrypted.Ciphertext, plain) {
		t.Fatal("plaintext must not appear in ciphertext blob")
	}
	decrypted, err := enc.EnvelopeDecrypt(ctx, DecryptRequest{
		Ciphertext: encrypted.Ciphertext,
		AAD:        aad,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(decrypted.Plaintext, plain) {
		t.Fatal("decrypted plaintext mismatch")
	}
}

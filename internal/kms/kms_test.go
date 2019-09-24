package kms

import (
	"bytes"
	"context"
	"testing"
)

func TestKMSWrapUnwrapRoundTrip(t *testing.T) {
	reg, err := NewRegistryWithMasterKey(bytes.Repeat([]byte{0xAB}, 32))
	if err != nil {
		t.Fatal(err)
	}
	svc := NewService(reg)
	ctx := context.Background()

	key, err := svc.CreateKey(ctx, "org-1", "default")
	if err != nil {
		t.Fatal(err)
	}
	dek := bytes.Repeat([]byte{0xCD}, 32)
	wrapped, err := svc.WrapDEK(ctx, WrapRequest{KeyID: key.ID, DEK: dek, AAD: []byte("aad")})
	if err != nil {
		t.Fatal(err)
	}
	unwrapped, err := svc.UnwrapDEK(ctx, UnwrapRequest{
		KeyID:      key.ID,
		WrappedDEK: wrapped.WrappedDEK,
		WrapNonce:  wrapped.WrapNonce,
		AAD:        []byte("aad"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(unwrapped.DEK, dek) {
		t.Fatal("dek mismatch after unwrap")
	}
}

func TestKMSRotateKey(t *testing.T) {
	reg, err := NewRegistryWithMasterKey(bytes.Repeat([]byte{0x01}, 32))
	if err != nil {
		t.Fatal(err)
	}
	svc := NewService(reg)
	ctx := context.Background()
	key, _ := svc.CreateKey(ctx, "org-1", "rotate-me")
	dek := bytes.Repeat([]byte{0x02}, 32)
	w1, _ := svc.WrapDEK(ctx, WrapRequest{KeyID: key.ID, DEK: dek, AAD: nil})
	rotated, err := svc.RotateKey(ctx, key.ID)
	if err != nil {
		t.Fatal(err)
	}
	if rotated.Version != 2 {
		t.Fatalf("expected version 2, got %d", rotated.Version)
	}
	if _, err := svc.UnwrapDEK(ctx, UnwrapRequest{KeyID: key.ID, WrappedDEK: w1.WrappedDEK, WrapNonce: w1.WrapNonce}); err == nil {
		t.Fatal("expected unwrap failure after rotation")
	}
}

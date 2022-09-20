package domain

import (
	"bytes"
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/kms"
)

func TestKMSWrapUnwrapRoundTrip(t *testing.T) {
	reg, err := kms.NewRegistryWithMasterKey(bytes.Repeat([]byte{0xAB}, 32))
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

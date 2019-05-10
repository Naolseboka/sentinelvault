package crypto

import (
	"bytes"
	"testing"
)

func TestDeriveKey(t *testing.T) {
	k1, err := DeriveKey([]byte("secret"), []byte("salt"), []byte("info"), 32)
	if err != nil {
		t.Fatal(err)
	}
	k2, err := DeriveKey([]byte("secret"), []byte("salt"), []byte("info"), 32)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(k1, k2) {
		t.Fatal("deterministic derivation failed")
	}
	k3, err := DeriveKey([]byte("other"), []byte("salt"), []byte("info"), 32)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(k1, k3) {
		t.Fatal("expected different keys")
	}
}

func TestDeriveKeyInvalidLength(t *testing.T) {
	if _, err := DeriveKey([]byte("s"), nil, nil, 0); err == nil {
		t.Fatal("expected error")
	}
}

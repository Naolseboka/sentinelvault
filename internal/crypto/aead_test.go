package crypto

import (
	"bytes"
	"testing"
)

func TestAEADEncryptDecrypt(t *testing.T) {
	key, err := RandomBytes(AESKeySize)
	if err != nil {
		t.Fatal(err)
	}
	aead, err := NewAEAD(key)
	if err != nil {
		t.Fatal(err)
	}
	nonce, err := RandomBytes(aead.NonceSize())
	if err != nil {
		t.Fatal(err)
	}
	plain := []byte("secret payload")
	aad := []byte("context")
	ct, err := aead.Encrypt(nonce, plain, aad)
	if err != nil {
		t.Fatal(err)
	}
	out, err := aead.Decrypt(nonce, ct, aad)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(out, plain) {
		t.Fatal("plaintext mismatch")
	}
}

func TestAEADInvalidKey(t *testing.T) {
	if _, err := NewAEAD([]byte("short")); err == nil {
		t.Fatal("expected key size error")
	}
}

func TestAEADTamper(t *testing.T) {
	key, _ := RandomBytes(AESKeySize)
	aead, _ := NewAEAD(key)
	nonce, _ := RandomBytes(aead.NonceSize())
	ct, _ := aead.Encrypt(nonce, []byte("x"), nil)
	ct[0] ^= 0xff
	if _, err := aead.Decrypt(nonce, ct, nil); err == nil {
		t.Fatal("expected auth failure")
	}
}

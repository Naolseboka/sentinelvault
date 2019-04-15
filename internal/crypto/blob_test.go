package crypto

import (
	"bytes"
	"testing"
)

func TestBlobRoundTrip(t *testing.T) {
	b := &CiphertextBlob{
		Version:    1,
		KeyID:      "key-abc",
		Nonce:      bytes.Repeat([]byte{1}, 12),
		WrapNonce:  bytes.Repeat([]byte{2}, 12),
		WrappedDEK: []byte("wrapped-dek-bytes"),
		Ciphertext: []byte("secret ciphertext"),
	}
	data, err := MarshalBlob(b)
	if err != nil {
		t.Fatal(err)
	}
	out, err := UnmarshalBlob(data)
	if err != nil {
		t.Fatal(err)
	}
	if out.KeyID != b.KeyID || !bytes.Equal(out.Ciphertext, b.Ciphertext) {
		t.Fatal("blob round trip mismatch")
	}
}

func TestGenerateDEK(t *testing.T) {
	dek, err := GenerateDEK()
	if err != nil {
		t.Fatal(err)
	}
	if len(dek) != AESKeySize {
		t.Fatalf("expected %d bytes, got %d", AESKeySize, len(dek))
	}
}

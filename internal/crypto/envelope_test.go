package crypto

import (
	"bytes"
	"testing"
)

func TestEnvelopeRoundTrip(t *testing.T) {
	kek, err := RandomBytes(AESKeySize)
	if err != nil {
		t.Fatal(err)
	}
	plain := []byte("envelope secret data")
	aad := []byte("item-123")
	env, err := EnvelopeEncrypt(plain, kek, aad)
	if err != nil {
		t.Fatal(err)
	}
	out, err := EnvelopeDecrypt(env, kek, aad)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(out, plain) {
		t.Fatal("decrypt mismatch")
	}
	data, err := MarshalJSONEnvelope(env)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := UnmarshalJSONEnvelope(data)
	if err != nil {
		t.Fatal(err)
	}
	out2, err := EnvelopeDecrypt(parsed, kek, aad)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(out2, plain) {
		t.Fatal("json round trip failed")
	}
}

func TestEnvelopeWrongKEK(t *testing.T) {
	kek, _ := RandomBytes(AESKeySize)
	env, _ := EnvelopeEncrypt([]byte("x"), kek, nil)
	wrong, _ := RandomBytes(AESKeySize)
	if _, err := EnvelopeDecrypt(env, wrong, nil); err == nil {
		t.Fatal("expected failure")
	}
}

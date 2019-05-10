package crypto

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"testing"
)

func TestAEADEdgeCases(t *testing.T) {
	key, _ := RandomBytes(AESKeySize)
	aead, _ := NewAEAD(key)

	if aead.Overhead() == 0 {
		t.Fatal("expected non-zero overhead")
	}
	if _, err := aead.Encrypt([]byte("bad"), []byte("x"), nil); err == nil {
		t.Fatal("expected invalid nonce error")
	}
	if _, err := aead.Decrypt([]byte("bad"), []byte("x"), nil); err == nil {
		t.Fatal("expected invalid nonce error")
	}
}

func TestSignVerifyInvalidKeys(t *testing.T) {
	if _, err := SignEd25519(ed25519.PrivateKey([]byte("short")), []byte("x")); err != ErrInvalidKeyLength {
		t.Fatalf("got %v", err)
	}
	if err := VerifyEd25519(ed25519.PublicKey([]byte("short")), []byte("x"), []byte("sig")); err != ErrInvalidKeyLength {
		t.Fatalf("got %v", err)
	}
}

func TestEnvelopeErrors(t *testing.T) {
	kek, _ := RandomBytes(AESKeySize)
	if _, err := EnvelopeEncrypt([]byte("x"), []byte("short"), nil); err == nil {
		t.Fatal("expected key error")
	}
	if _, err := EnvelopeDecrypt(nil, kek, nil); err == nil {
		t.Fatal("expected nil envelope error")
	}
	env := &Envelope{Version: 2, Algorithm: "AES-256-GCM"}
	if _, err := EnvelopeDecrypt(env, kek, nil); err == nil {
		t.Fatal("expected version error")
	}
	env = &Envelope{Version: 1, Algorithm: "legacy"}
	if _, err := EnvelopeDecrypt(env, kek, nil); err == nil {
		t.Fatal("expected algorithm error")
	}
	for _, field := range []struct {
		name string
		env  *Envelope
	}{
		{"nonce", &Envelope{Version: 1, Algorithm: "AES-256-GCM", Nonce: "!!!", Ciphertext: "QQ", WrappedDEK: "QQ", WrapNonce: "QQ"}},
		{"ciphertext", &Envelope{Version: 1, Algorithm: "AES-256-GCM", Nonce: "AQID", Ciphertext: "!!!", WrappedDEK: "QQ", WrapNonce: "QQ"}},
		{"wrapped", &Envelope{Version: 1, Algorithm: "AES-256-GCM", Nonce: "AQID", Ciphertext: "QQ", WrappedDEK: "!!!", WrapNonce: "QQ"}},
		{"wrapnonce", &Envelope{Version: 1, Algorithm: "AES-256-GCM", Nonce: "AQID", Ciphertext: "QQ", WrappedDEK: "QQ", WrapNonce: "!!!"}},
	} {
		if _, err := EnvelopeDecrypt(field.env, kek, nil); err == nil {
			t.Fatalf("expected decode error for %s", field.name)
		}
	}
	if _, err := UnmarshalJSONEnvelope([]byte("{")); err == nil {
		t.Fatal("expected json error")
	}
}

func TestEnvelopeWrongAAD(t *testing.T) {
	kek, _ := RandomBytes(AESKeySize)
	env, _ := EnvelopeEncrypt([]byte("secret"), kek, []byte("item-1"))
	if _, err := EnvelopeDecrypt(env, kek, []byte("item-2")); err == nil {
		t.Fatal("expected AAD mismatch")
	}
}

func TestEnvelopeTamperedCiphertext(t *testing.T) {
	kek, _ := RandomBytes(AESKeySize)
	env, _ := EnvelopeEncrypt([]byte("secret"), kek, nil)
	raw, _ := base64.RawStdEncoding.DecodeString(env.Ciphertext)
	raw[0] ^= 0xff
	env.Ciphertext = base64.RawStdEncoding.EncodeToString(raw)
	if _, err := EnvelopeDecrypt(env, kek, nil); err == nil {
		t.Fatal("expected tamper detection")
	}
}

func TestEnvelopeJSONRoundTrip(t *testing.T) {
	kek, _ := RandomBytes(AESKeySize)
	env, _ := EnvelopeEncrypt([]byte("data"), kek, []byte("aad"))
	data, err := MarshalJSONEnvelope(env)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := UnmarshalJSONEnvelope(data)
	if err != nil {
		t.Fatal(err)
	}
	if parsed.Version != 1 {
		t.Fatal("bad version")
	}
	_ = json.Valid(data)
}

func TestEnvelopeTamperedWrappedDEK(t *testing.T) {
	kek, _ := RandomBytes(AESKeySize)
	env, _ := EnvelopeEncrypt([]byte("secret"), kek, nil)
	raw, _ := base64.RawStdEncoding.DecodeString(env.WrappedDEK)
	raw[0] ^= 0xff
	env.WrappedDEK = base64.RawStdEncoding.EncodeToString(raw)
	if _, err := EnvelopeDecrypt(env, kek, nil); err == nil {
		t.Fatal("expected unwrap failure")
	}
}

func TestAEADWithAssociatedData(t *testing.T) {
	key, _ := RandomBytes(AESKeySize)
	aead, _ := NewAEAD(key)
	nonce, _ := RandomBytes(aead.NonceSize())
	ct, _ := aead.Encrypt(nonce, []byte("data"), []byte("ctx"))
	if _, err := aead.Decrypt(nonce, ct, []byte("wrong")); err == nil {
		t.Fatal("expected AAD failure")
	}
}

func TestMultiplePasswordHashes(t *testing.T) {
	h1, err := HashPassword("password-one")
	if err != nil {
		t.Fatal(err)
	}
	h2, err := HashPassword("password-two")
	if err != nil {
		t.Fatal(err)
	}
	if h1 == h2 {
		t.Fatal("expected unique salts")
	}
}
func TestArgon2DecodeErrors(t *testing.T) {
	if err := VerifyPassword("pw", "bad$hash"); err == nil {
		t.Fatal("expected decode error")
	}
	if err := VerifyPassword("pw", "!!!$!!!"); err == nil {
		t.Fatal("expected salt decode error")
	}
	if err := VerifyPassword("pw", "AQID$!!!"); err == nil {
		t.Fatal("expected hash decode error")
	}
}

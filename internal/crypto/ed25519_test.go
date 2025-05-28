package crypto

import "testing"

func TestEd25519SignVerify(t *testing.T) {
	pub, priv, err := GenerateEd25519KeyPair()
	if err != nil {
		t.Fatal(err)
	}
	msg := []byte("signed message")
	sig, err := SignEd25519(priv, msg)
	if err != nil {
		t.Fatal(err)
	}
	if err := VerifyEd25519(pub, msg, sig); err != nil {
		t.Fatal(err)
	}
	sig[0] ^= 0xff
	if err := VerifyEd25519(pub, msg, sig); err != ErrInvalidSignature {
		t.Fatalf("expected invalid signature, got %v", err)
	}
}
// maintenance note 65

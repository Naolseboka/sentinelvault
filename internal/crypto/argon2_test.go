package crypto

import "testing"

func TestArgon2HashVerify(t *testing.T) {
	hash, err := HashPassword("correct horse battery staple")
	if err != nil {
		t.Fatal(err)
	}
	if err := VerifyPassword("correct horse battery staple", hash); err != nil {
		t.Fatal(err)
	}
	if err := VerifyPassword("wrong", hash); err != ErrPasswordMismatch {
		t.Fatalf("expected mismatch, got %v", err)
	}
}

func TestArgon2InvalidFormat(t *testing.T) {
	if err := VerifyPassword("x", "bad"); err == nil {
		t.Fatal("expected error")
	}
}

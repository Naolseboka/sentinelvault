package transit

import "testing"

func TestTransitEncryptDecryptRotate(t *testing.T) {
	e := NewEngine()
	if _, err := e.CreateKey("o1", "app", KeyTypeAES256); err != nil {
		t.Fatal(err)
	}
	ct, err := e.Encrypt("o1", "app", []byte("hello-secret"))
	if err != nil {
		t.Fatal(err)
	}
	pt, err := e.Decrypt("o1", "app", ct)
	if err != nil {
		t.Fatal(err)
	}
	if string(pt) != "hello-secret" {
		t.Fatalf("%s", pt)
	}
	if _, err := e.Rotate("o1", "app"); err != nil {
		t.Fatal(err)
	}
	// old ciphertext still decrypts
	pt2, err := e.Decrypt("o1", "app", ct)
	if err != nil || string(pt2) != "hello-secret" {
		t.Fatalf("pt2=%s err=%v", pt2, err)
	}
}

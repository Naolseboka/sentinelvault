package sshx

import (
	"testing"
	"time"
)

func TestSSHManager(t *testing.T) {
	m, err := NewManager()
	if err != nil {
		t.Fatal(err)
	}
	key, err := m.GenerateKey("org1", "deploy", KeyTypeUser)
	if err != nil {
		t.Fatal(err)
	}
	if key.Fingerprint == "" || key.PublicKey == "" {
		t.Fatal("missing key material")
	}
	cert, err := m.SignUserCert("org1", key.ID, []string{"ubuntu"}, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if cert.CertString == "" {
		t.Fatal("missing cert")
	}
	rotated, err := m.RotateKey(key.ID)
	if err != nil || rotated.RotatedAt == nil {
		t.Fatalf("rotate=%+v err=%v", rotated, err)
	}
}
// maintenance note 71

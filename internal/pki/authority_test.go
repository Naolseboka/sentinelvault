package pki

import (
	"testing"
	"time"
)

func TestIssueAndRevoke(t *testing.T) {
	ca, err := NewAuthority("SentinelVault Test CA")
	if err != nil {
		t.Fatal(err)
	}
	if ca.CACertPEM() == "" {
		t.Fatal("missing CA PEM")
	}
	cert, err := ca.Issue("org1", "api.example.com", []string{"api.example.com"}, 24*time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if cert.CertPEM == "" || cert.KeyPEM == "" {
		t.Fatal("missing pem")
	}
	got, err := ca.Get(cert.ID)
	if err != nil || got.KeyPEM != "" {
		t.Fatalf("get=%+v err=%v", got, err)
	}
	rev, err := ca.Revoke(cert.ID)
	if err != nil || !rev.Revoked {
		t.Fatalf("revoke=%+v err=%v", rev, err)
	}
}

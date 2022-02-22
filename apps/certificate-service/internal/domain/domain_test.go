package domain

import (
	"context"
	"testing"
)

func TestCertificateIssueRevokeGet(t *testing.T) {
	auth, err := NewAuthority("SentinelVault Dev CA")
	if err != nil {
		t.Fatal(err)
	}
	svc := NewService(auth)
	ctx := context.Background()

	rec, err := svc.Issue(ctx, IssueInput{OrgID: "org-1", CommonName: "api.example.com", DNSNames: []string{"api.example.com"}})
	if err != nil {
		t.Fatal(err)
	}
	if rec.CertPEM == "" {
		t.Fatal("expected cert PEM")
	}

	got, err := svc.Get(ctx, rec.ID)
	if err != nil || got.CommonName != "api.example.com" {
		t.Fatal("get failed")
	}

	revoked, err := svc.Revoke(ctx, rec.ID)
	if err != nil || !revoked.Revoked {
		t.Fatal("revoke failed")
	}
	if svc.CACertPEM(ctx) == "" {
		t.Fatal("expected CA PEM")
	}
}

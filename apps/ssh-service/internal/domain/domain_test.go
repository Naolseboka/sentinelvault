package domain

import (
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/sshx"
)

func TestSSHKeyRotateAndSign(t *testing.T) {
	mgr, err := NewManager()
	if err != nil {
		t.Fatal(err)
	}
	svc := NewService(mgr)
	ctx := context.Background()

	key, err := svc.GenerateKey(ctx, GenerateKeyInput{OrgID: "org-1", Name: "deploy", Type: sshx.KeyTypeUser})
	if err != nil {
		t.Fatal(err)
	}

	rotated, err := svc.RotateKey(ctx, key.ID)
	if err != nil || rotated.Fingerprint == key.Fingerprint {
		t.Fatal("rotate should change fingerprint")
	}

	cert, err := svc.SignUserCert(ctx, SignCertInput{
		OrgID: "org-1", KeyID: key.ID, Principals: []string{"deploy@host"},
	})
	if err != nil || cert.CertString == "" {
		t.Fatal("sign cert failed")
	}
}

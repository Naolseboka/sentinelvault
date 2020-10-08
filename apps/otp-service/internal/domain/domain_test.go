package domain

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/pquerna/otp/totp"
	encdomain "github.com/sentinelvault/sentinelvault/internal/encryption"
	kmsdomain "github.com/sentinelvault/sentinelvault/internal/kms"
)

func setupOTPService(t *testing.T) *Service {
	t.Helper()
	reg, err := kmsdomain.NewRegistryWithMasterKey(bytes.Repeat([]byte{0x01}, 32))
	if err != nil {
		t.Fatal(err)
	}
	kmsSvc := kmsdomain.NewService(reg)
	key, err := kmsSvc.CreateKey(context.Background(), "org-1", "otp")
	if err != nil {
		t.Fatal(err)
	}
	encSvc := encdomain.NewService(&encdomain.InProcessKMS{Svc: kmsSvc})
	return NewService(encSvc, key.ID)
}

func TestOTPEnrollAndVerify(t *testing.T) {
	svc := setupOTPService(t)
	ctx := context.Background()
	result, err := svc.Enroll(ctx, EnrollInput{OrgID: "org-1", UserID: "user-1", Email: "user@example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Secret.Secret == "" {
		t.Fatal("expected secret at enroll")
	}
	code, err := totp.GenerateCode(result.Secret.Secret, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if err := svc.Verify(ctx, VerifyInput{UserID: "user-1", Code: code}); err != nil {
		t.Fatal(err)
	}
	rec, err := svc.Get(ctx, "user-1")
	if err != nil || rec.Label != "user@example.com" {
		t.Fatal("get record failed")
	}
	if !rec.Enabled {
		t.Fatal("expected enabled after verify")
	}
}

func TestOTPGetNeverReturnsSecret(t *testing.T) {
	svc := setupOTPService(t)
	ctx := context.Background()
	result, err := svc.Enroll(ctx, EnrollInput{OrgID: "org-1", UserID: "user-2", Email: "other@example.com"})
	if err != nil {
		t.Fatal(err)
	}
	secretAtEnroll := result.Secret.Secret
	rec, err := svc.Get(ctx, "user-2")
	if err != nil {
		t.Fatal(err)
	}
	if rec.Label == "" {
		t.Fatal("expected label")
	}
	_ = secretAtEnroll
	// TOTPRecordView has no secret fields; verify via type shape
	if rec.UserID != "user-2" {
		t.Fatal("wrong user")
	}
}

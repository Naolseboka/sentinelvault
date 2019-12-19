package domain_test

import (
	"context"
	"testing"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/sentinelvault/sentinelvault/apps/auth-service/internal/domain"
)

func newTestService() *domain.Service {
	return domain.NewService(
		domain.NewMemoryCredentialRepository(),
		domain.NewMemorySessionRepository(),
		domain.ServiceConfig{JWTSecret: []byte("test-jwt-secret-for-unit-tests!!")},
	)
}

func TestRegisterAndLogin(t *testing.T) {
	t.Parallel()
	svc := newTestService()
	ctx := context.Background()

	user, org, err := svc.Register(ctx, domain.RegisterInput{
		Email:    "alice@example.com",
		Password: "secure-password-123",
		OrgName:  "Acme Corp",
	})
	if err != nil {
		t.Fatal(err)
	}
	if user.Email != "alice@example.com" || org.Name != "Acme Corp" {
		t.Fatalf("user=%+v org=%+v", user, org)
	}

	result, err := svc.Login(ctx, domain.LoginInput{
		Email:             "alice@example.com",
		Password:          "secure-password-123",
		DeviceFingerprint: "device-1",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.MFARequired || result.Tokens == nil || result.Tokens.AccessToken == "" {
		t.Fatalf("login result=%+v", result)
	}
}

func TestLoginInvalidPassword(t *testing.T) {
	t.Parallel()
	svc := newTestService()
	ctx := context.Background()
	_, _, err := svc.Register(ctx, domain.RegisterInput{
		Email: "bob@example.com", Password: "pw", OrgName: "Org",
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = svc.Login(ctx, domain.LoginInput{Email: "bob@example.com", Password: "wrong"})
	if err == nil {
		t.Fatal("expected login failure")
	}
}

func TestMFAEnrollmentAndLogin(t *testing.T) {
	t.Parallel()
	svc := newTestService()
	ctx := context.Background()
	user, _, err := svc.Register(ctx, domain.RegisterInput{
		Email: "mfa@example.com", Password: "pw", OrgName: "MFA Org",
	})
	if err != nil {
		t.Fatal(err)
	}
	secret, err := svc.EnrollTOTP(ctx, user.ID)
	if err != nil {
		t.Fatal(err)
	}
	code, err := totp.GenerateCode(secret.Secret, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if err := svc.VerifyTOTP(ctx, user.ID, code); err != nil {
		t.Fatal(err)
	}

	pending, err := svc.Login(ctx, domain.LoginInput{
		Email: "mfa@example.com", Password: "pw", DeviceFingerprint: "dev",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !pending.MFARequired {
		t.Fatal("expected mfa required")
	}

	code, err = totp.GenerateCode(secret.Secret, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	result, err := svc.Login(ctx, domain.LoginInput{
		Email: "mfa@example.com", Password: "pw", DeviceFingerprint: "dev", TOTPCode: code,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Tokens == nil {
		t.Fatal("expected tokens after mfa")
	}
}

func TestRefreshAndLogout(t *testing.T) {
	t.Parallel()
	svc := newTestService()
	ctx := context.Background()
	user, _, _ := svc.Register(ctx, domain.RegisterInput{
		Email: "refresh@example.com", Password: "pw", OrgName: "Org",
	})
	result, err := svc.Login(ctx, domain.LoginInput{
		Email: "refresh@example.com", Password: "pw", DeviceFingerprint: "d",
	})
	if err != nil {
		t.Fatal(err)
	}
	refreshed, err := svc.Refresh(ctx, result.Tokens.RefreshToken)
	if err != nil {
		t.Fatal(err)
	}
	if refreshed.AccessToken == "" {
		t.Fatal("expected refreshed access token")
	}
	if err := svc.Logout(ctx, result.SessionID); err != nil {
		t.Fatal(err)
	}
	_ = user
}

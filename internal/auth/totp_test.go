package auth_test

import (
	"testing"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/sentinelvault/sentinelvault/internal/auth"
)

func TestGenerateAndValidateTOTP(t *testing.T) {
	t.Parallel()
	secret, err := auth.GenerateTOTPSecret("SentinelVault", "user@example.com")
	if err != nil {
		t.Fatal(err)
	}
	if secret.Secret == "" || secret.URI == "" {
		t.Fatal("expected secret and uri")
	}
	code, err := totp.GenerateCode(secret.Secret, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if !auth.ValidateTOTPCode(secret.Secret, code) {
		t.Fatal("expected valid totp code")
	}
	if auth.ValidateTOTPCode(secret.Secret, "000000") {
		t.Fatal("expected invalid code to fail")
	}
}

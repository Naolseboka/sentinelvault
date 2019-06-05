package validators

import "testing"

func TestOTPSeedValidatorHappy(t *testing.T) {
	v := OTPSeedValidator{}
	if v.Type() != "otp_seed" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"secret": "value-secret"})
	if err != nil {
		t.Fatal(err)
	}
	if len(norm) == 0 {
		t.Fatal("empty normalized")
	}
	red := v.Redact(norm)
	if red == nil {
		t.Fatal("redact nil")
	}
	if len(v.RequiredFields()) == 0 {
		t.Fatal("required empty")
	}
}

func TestOTPSeedValidatorMissing(t *testing.T) {
	v := OTPSeedValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

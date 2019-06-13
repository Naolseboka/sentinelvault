package validators

import "testing"

func TestRecoveryCodesValidatorHappy(t *testing.T) {
	v := RecoveryCodesValidator{}
	if v.Type() != "recovery_codes" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"codes": "value-codes"})
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

func TestRecoveryCodesValidatorMissing(t *testing.T) {
	v := RecoveryCodesValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

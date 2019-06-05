package validators

import "testing"

func TestPasskeyValidatorHappy(t *testing.T) {
	v := PasskeyValidator{}
	if v.Type() != "passkey" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"credential_id": "value-credential_id", "public_key": "value-public_key"})
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

func TestPasskeyValidatorMissing(t *testing.T) {
	v := PasskeyValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

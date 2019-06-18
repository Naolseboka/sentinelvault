package validators

import "testing"

func TestGenericSecretValidatorHappy(t *testing.T) {
	v := GenericSecretValidator{}
	if v.Type() != "secret" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"value": "value-value"})
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

func TestGenericSecretValidatorMissing(t *testing.T) {
	v := GenericSecretValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

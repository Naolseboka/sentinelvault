package validators

import "testing"

func TestIdentityValidatorHappy(t *testing.T) {
	v := IdentityValidator{}
	if v.Type() != "identity" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"first_name": "value-first_name", "last_name": "value-last_name"})
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

func TestIdentityValidatorMissing(t *testing.T) {
	v := IdentityValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

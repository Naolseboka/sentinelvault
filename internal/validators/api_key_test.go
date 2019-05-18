package validators

import "testing"

func TestAPIKeyValidatorHappy(t *testing.T) {
	v := APIKeyValidator{}
	if v.Type() != "api_key" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"key": "value-key"})
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

func TestAPIKeyValidatorMissing(t *testing.T) {
	v := APIKeyValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

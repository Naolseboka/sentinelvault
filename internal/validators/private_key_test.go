package validators

import "testing"

func TestPrivateKeyValidatorHappy(t *testing.T) {
	v := PrivateKeyValidator{}
	if v.Type() != "private_key" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"pem": "value-pem"})
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

func TestPrivateKeyValidatorMissing(t *testing.T) {
	v := PrivateKeyValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

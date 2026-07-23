package validators

import "testing"

func TestPGPKeyValidatorHappy(t *testing.T) {
	v := PGPKeyValidator{}
	if v.Type() != "pgp_key" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"private_armor": "value-private_armor"})
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

func TestPGPKeyValidatorMissing(t *testing.T) {
	v := PGPKeyValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

// maintenance note 89

package validators

import "testing"

func TestEncryptionKeyValidatorHappy(t *testing.T) {
	v := EncryptionKeyValidator{}
	if v.Type() != "encryption_key" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"key_material": "value-key_material"})
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

func TestEncryptionKeyValidatorMissing(t *testing.T) {
	v := EncryptionKeyValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

package validators

import "testing"

func TestSecureNoteValidatorHappy(t *testing.T) {
	v := SecureNoteValidator{}
	if v.Type() != "secure_note" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"note": "value-note"})
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

func TestSecureNoteValidatorMissing(t *testing.T) {
	v := SecureNoteValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

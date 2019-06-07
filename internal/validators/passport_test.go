package validators

import "testing"

func TestPassportValidatorHappy(t *testing.T) {
	v := PassportValidator{}
	if v.Type() != "passport" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"number": "value-number", "country": "value-country", "expiry": "value-expiry"})
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

func TestPassportValidatorMissing(t *testing.T) {
	v := PassportValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

package validators

import "testing"

func TestLicenseValidatorHappy(t *testing.T) {
	v := LicenseValidator{}
	if v.Type() != "license" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"number": "value-number", "state": "value-state", "expiry": "value-expiry"})
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

func TestLicenseValidatorMissing(t *testing.T) {
	v := LicenseValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

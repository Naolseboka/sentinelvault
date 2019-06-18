package validators

import "testing"

func TestSoftwareLicenseValidatorHappy(t *testing.T) {
	v := SoftwareLicenseValidator{}
	if v.Type() != "software_license" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"license_key": "value-license_key"})
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

func TestSoftwareLicenseValidatorMissing(t *testing.T) {
	v := SoftwareLicenseValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

package validators

import "testing"

func TestCertificateValidatorHappy(t *testing.T) {
	v := CertificateValidator{}
	if v.Type() != "certificate" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"cert_pem": "value-cert_pem"})
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

func TestCertificateValidatorMissing(t *testing.T) {
	v := CertificateValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

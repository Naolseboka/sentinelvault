package validators

import "testing"

func TestDatabaseCredentialValidatorHappy(t *testing.T) {
	v := DatabaseCredentialValidator{}
	if v.Type() != "database_credential" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"username": "value-username", "password": "value-password", "host": "value-host"})
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

func TestDatabaseCredentialValidatorMissing(t *testing.T) {
	v := DatabaseCredentialValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

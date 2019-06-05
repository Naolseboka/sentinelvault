package validators

import "testing"

func TestEnvVarValidatorHappy(t *testing.T) {
	v := EnvVarValidator{}
	if v.Type() != "env_var" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"name": "value-name", "value": "value-value"})
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

func TestEnvVarValidatorMissing(t *testing.T) {
	v := EnvVarValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

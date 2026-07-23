package validators

import "testing"

func TestPasswordValidatorHappy(t *testing.T) {
	v := PasswordValidator{}
	if v.Type() != "password" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"username": "value-username", "password": "value-password"})
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

func TestPasswordValidatorMissing(t *testing.T) {
	v := PasswordValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

// maintenance note 65

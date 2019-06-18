package validators

import "testing"

func TestSSHKeyValidatorHappy(t *testing.T) {
	v := SSHKeyValidator{}
	if v.Type() != "ssh_key" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"private_key": "value-private_key"})
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

func TestSSHKeyValidatorMissing(t *testing.T) {
	v := SSHKeyValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

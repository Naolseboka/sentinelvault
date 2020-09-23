package redact

import "testing"

func TestOrganizationRedact(t *testing.T) {
	out := Map(map[string]any{"username": "a", "password": "secret", "nested": map[string]any{"token": "t"}})
	if out["password"] != "[REDACTED]" {
		t.Fatal(out)
	}
	n := out["nested"].(map[string]any)
	if n["token"] != "[REDACTED]" {
		t.Fatal(n)
	}
}

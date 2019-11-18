package fields

import "testing"

func TestProcessors(t *testing.T) {
	if len(All()) == 0 {
		t.Fatal("empty")
	}
	out, _, err := ProcessPayload(map[string]any{"username": "  a  ", "password": "longenough"})
	if err != nil || out["username"] != "a" {
		t.Fatal(err, out)
	}
	if ByName("password").Redact("x") != "[REDACTED]" {
		t.Fatal("redact")
	}
}

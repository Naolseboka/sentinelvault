package searchx

import "testing"

func TestExtractors(t *testing.T) {
	m := ExtractPassword(map[string]any{"username": "a", "password": "secret", "url": "https://x"})
	if m["url"] != "https://x" || m["username"] != "a" {
		t.Fatalf("%v", m)
	}
	m2 := ExtractForType("api_key", map[string]any{"provider": "aws"})
	if m2["provider"] != "aws" {
		t.Fatalf("%v", m2)
	}
}

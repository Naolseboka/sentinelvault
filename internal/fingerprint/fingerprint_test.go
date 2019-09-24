package fingerprint

import "testing"

func TestFingerprint(t *testing.T) {
	h := HashUserAgent("Mozilla/5.0")
	if h == "" {
		t.Fatal("empty")
	}
	c := Composite(map[string]string{"user_agent": "a", "platform": "b"})
	if len(c) != 64 {
		t.Fatal(c)
	}
}

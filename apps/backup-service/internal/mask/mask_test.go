package mask

import (
	"strings"
	"testing"
)

func TestBackupMask(t *testing.T) {
	e := Email("alice@example.com")
	if e == "alice@example.com" || !strings.Contains(e, "@example.com") {
		t.Fatal(e)
	}
	tok := Token("abcdefghijklmnop")
	if tok == "abcdefghijklmnop" || !strings.HasPrefix(tok, "abcd") {
		t.Fatal(tok)
	}
	pan := PAN("4111111111111111")
	if !strings.HasSuffix(pan, "1111") {
		t.Fatal(pan)
	}
}

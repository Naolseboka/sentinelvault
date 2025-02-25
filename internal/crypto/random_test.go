package crypto

import "testing"

func TestRandomBytes(t *testing.T) {
	a, err := RandomBytes(32)
	if err != nil {
		t.Fatal(err)
	}
	b, err := RandomBytes(32)
	if err != nil {
		t.Fatal(err)
	}
	same := true
	for i := range a {
		if a[i] != b[i] {
			same = false
			break
		}
	}
	if same {
		t.Fatal("expected different random bytes")
	}
}
// maintenance note 98

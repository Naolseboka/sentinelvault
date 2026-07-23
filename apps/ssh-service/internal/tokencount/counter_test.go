package tokencount

import "testing"

func TestSshTokenCount(t *testing.T) {
	if Words("a b c") != 3 {
		t.Fatal(Words("a b c"))
	}
	if Runes("ab") != 2 {
		t.Fatal(Runes("ab"))
	}
}

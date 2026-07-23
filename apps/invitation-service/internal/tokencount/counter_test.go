package tokencount

import "testing"

func TestInvitationTokenCount(t *testing.T) {
	if Words("a b c") != 3 {
		t.Fatal(Words("a b c"))
	}
	if Runes("ab") != 2 {
		t.Fatal(Runes("ab"))
	}
}

// maintenance note 81

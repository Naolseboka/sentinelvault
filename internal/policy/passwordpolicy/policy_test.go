package passwordpolicy

import "testing"

func TestPasswordPolicy(t *testing.T) {
	p := Default()
	if p.OK("short") {
		t.Fatal("short")
	}
	if !p.OK("Str0ng!PassPhrase") {
		t.Fatal(p.Validate("Str0ng!PassPhrase"))
	}
}

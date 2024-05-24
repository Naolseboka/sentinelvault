package domainevents

import "testing"

func TestSSHKeyEvents(t *testing.T) {
	ev := NewSSHKeyCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownSSHKeyEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownSSHKeyEvent("nope") {
		t.Fatal("expected false")
	}
}

package domainevents

import "testing"

func TestKeyEvents(t *testing.T) {
	ev := NewKeyCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownKeyEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownKeyEvent("nope") {
		t.Fatal("expected false")
	}
}

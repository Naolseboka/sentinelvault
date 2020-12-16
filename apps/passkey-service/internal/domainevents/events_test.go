package domainevents

import "testing"

func TestPasskeyEvents(t *testing.T) {
	ev := NewPasskeyCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownPasskeyEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownPasskeyEvent("nope") {
		t.Fatal("expected false")
	}
}

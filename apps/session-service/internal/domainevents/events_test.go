package domainevents

import "testing"

func TestSessionEvents(t *testing.T) {
	ev := NewSessionCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownSessionEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownSessionEvent("nope") {
		t.Fatal("expected false")
	}
}

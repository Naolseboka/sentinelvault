package domainevents

import "testing"

func TestSignalEvents(t *testing.T) {
	ev := NewSignalCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownSignalEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownSignalEvent("nope") {
		t.Fatal("expected false")
	}
}

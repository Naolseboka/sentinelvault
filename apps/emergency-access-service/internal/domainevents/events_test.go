package domainevents

import "testing"

func TestEmergencyGrantEvents(t *testing.T) {
	ev := NewEmergencyGrantCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownEmergencyGrantEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownEmergencyGrantEvent("nope") {
		t.Fatal("expected false")
	}
}

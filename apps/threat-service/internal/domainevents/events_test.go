package domainevents

import "testing"

func TestFindingEvents(t *testing.T) {
	ev := NewFindingCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownFindingEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownFindingEvent("nope") {
		t.Fatal("expected false")
	}
}

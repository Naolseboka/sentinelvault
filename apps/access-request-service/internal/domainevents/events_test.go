package domainevents

import "testing"

func TestAccessRequestEvents(t *testing.T) {
	ev := NewAccessRequestCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownAccessRequestEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownAccessRequestEvent("nope") {
		t.Fatal("expected false")
	}
}

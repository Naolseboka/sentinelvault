package domainevents

import "testing"

func TestSCIMUserEvents(t *testing.T) {
	ev := NewSCIMUserCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownSCIMUserEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownSCIMUserEvent("nope") {
		t.Fatal("expected false")
	}
}

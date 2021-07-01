package domainevents

import "testing"

func TestUserEvents(t *testing.T) {
	ev := NewUserCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownUserEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownUserEvent("nope") {
		t.Fatal("expected false")
	}
}

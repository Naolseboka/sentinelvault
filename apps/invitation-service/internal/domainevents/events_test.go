package domainevents

import "testing"

func TestInvitationEvents(t *testing.T) {
	ev := NewInvitationCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownInvitationEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownInvitationEvent("nope") {
		t.Fatal("expected false")
	}
}

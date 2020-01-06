package domainevents

import "testing"

func TestAuthzDecisionEvents(t *testing.T) {
	ev := NewAuthzDecisionCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownAuthzDecisionEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownAuthzDecisionEvent("nope") {
		t.Fatal("expected false")
	}
}

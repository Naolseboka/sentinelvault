package domainevents

import "testing"

func TestDelegationEvents(t *testing.T) {
	ev := NewDelegationCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownDelegationEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownDelegationEvent("nope") {
		t.Fatal("expected false")
	}
}

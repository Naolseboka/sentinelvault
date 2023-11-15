package domainevents

import "testing"

func TestProvisionPlanEvents(t *testing.T) {
	ev := NewProvisionPlanCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownProvisionPlanEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownProvisionPlanEvent("nope") {
		t.Fatal("expected false")
	}
}

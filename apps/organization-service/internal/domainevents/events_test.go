package domainevents

import "testing"

func TestOrganizationEvents(t *testing.T) {
	ev := NewOrganizationCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownOrganizationEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownOrganizationEvent("nope") {
		t.Fatal("expected false")
	}
}

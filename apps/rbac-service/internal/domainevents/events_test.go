package domainevents

import "testing"

func TestRoleBindingEvents(t *testing.T) {
	ev := NewRoleBindingCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownRoleBindingEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownRoleBindingEvent("nope") {
		t.Fatal("expected false")
	}
}

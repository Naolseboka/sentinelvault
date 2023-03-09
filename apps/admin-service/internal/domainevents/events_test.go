package domainevents

import "testing"

func TestAdminActionEvents(t *testing.T) {
	ev := NewAdminActionCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownAdminActionEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownAdminActionEvent("nope") {
		t.Fatal("expected false")
	}
}

package domainevents

import "testing"

func TestAuditEventEvents(t *testing.T) {
	ev := NewAuditEventCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownAuditEventEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownAuditEventEvent("nope") {
		t.Fatal("expected false")
	}
}

package domainevents

import "testing"

func TestApprovalEvents(t *testing.T) {
	ev := NewApprovalCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownApprovalEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownApprovalEvent("nope") {
		t.Fatal("expected false")
	}
}

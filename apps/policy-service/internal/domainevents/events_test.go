package domainevents

import "testing"

func TestPolicyDocEvents(t *testing.T) {
	ev := NewPolicyDocCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownPolicyDocEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownPolicyDocEvent("nope") {
		t.Fatal("expected false")
	}
}

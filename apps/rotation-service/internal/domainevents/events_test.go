package domainevents

import "testing"

func TestRotationPolicyEvents(t *testing.T) {
	ev := NewRotationPolicyCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownRotationPolicyEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownRotationPolicyEvent("nope") {
		t.Fatal("expected false")
	}
}

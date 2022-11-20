package domainevents

import "testing"

func TestRiskScoreEvents(t *testing.T) {
	ev := NewRiskScoreCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownRiskScoreEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownRiskScoreEvent("nope") {
		t.Fatal("expected false")
	}
}

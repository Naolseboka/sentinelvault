package domainevents

import "testing"

func TestTeamEvents(t *testing.T) {
	ev := NewTeamCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownTeamEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownTeamEvent("nope") {
		t.Fatal("expected false")
	}
}

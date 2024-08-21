package domainevents

import "testing"

func TestVersionEvents(t *testing.T) {
	ev := NewVersionCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownVersionEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownVersionEvent("nope") {
		t.Fatal("expected false")
	}
}

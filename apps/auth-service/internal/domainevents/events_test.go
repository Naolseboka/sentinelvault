package domainevents

import "testing"

func TestCredentialEvents(t *testing.T) {
	ev := NewCredentialCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownCredentialEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownCredentialEvent("nope") {
		t.Fatal("expected false")
	}
}

package domainevents

import "testing"

func TestSecretItemEvents(t *testing.T) {
	ev := NewSecretItemCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownSecretItemEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownSecretItemEvent("nope") {
		t.Fatal("expected false")
	}
}

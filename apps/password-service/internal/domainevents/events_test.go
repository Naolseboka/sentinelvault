package domainevents

import "testing"

func TestVaultItemEvents(t *testing.T) {
	ev := NewVaultItemCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownVaultItemEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownVaultItemEvent("nope") {
		t.Fatal("expected false")
	}
}

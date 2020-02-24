package domainevents

import "testing"

func TestTrustRecordEvents(t *testing.T) {
	ev := NewTrustRecordCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownTrustRecordEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownTrustRecordEvent("nope") {
		t.Fatal("expected false")
	}
}

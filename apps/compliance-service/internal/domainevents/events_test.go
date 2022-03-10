package domainevents

import "testing"

func TestEvidenceEvents(t *testing.T) {
	ev := NewEvidenceCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownEvidenceEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownEvidenceEvent("nope") {
		t.Fatal("expected false")
	}
}

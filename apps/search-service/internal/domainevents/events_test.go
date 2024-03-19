package domainevents

import "testing"

func TestDocumentEvents(t *testing.T) {
	ev := NewDocumentCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownDocumentEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownDocumentEvent("nope") {
		t.Fatal("expected false")
	}
}

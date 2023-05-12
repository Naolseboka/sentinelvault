package domainevents

import "testing"

func TestAttachmentEvents(t *testing.T) {
	ev := NewAttachmentCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownAttachmentEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownAttachmentEvent("nope") {
		t.Fatal("expected false")
	}
}

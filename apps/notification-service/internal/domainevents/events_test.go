package domainevents

import "testing"

func TestMessageEvents(t *testing.T) {
	ev := NewMessageCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownMessageEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownMessageEvent("nope") {
		t.Fatal("expected false")
	}
}

package domainevents

import "testing"

func TestEncryptOpEvents(t *testing.T) {
	ev := NewEncryptOpCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownEncryptOpEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownEncryptOpEvent("nope") {
		t.Fatal("expected false")
	}
}

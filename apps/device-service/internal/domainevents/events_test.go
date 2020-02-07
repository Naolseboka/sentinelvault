package domainevents

import "testing"

func TestDeviceEvents(t *testing.T) {
	ev := NewDeviceCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownDeviceEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownDeviceEvent("nope") {
		t.Fatal("expected false")
	}
}

package domainevents

import "testing"

func TestSyncJobEvents(t *testing.T) {
	ev := NewSyncJobCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownSyncJobEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownSyncJobEvent("nope") {
		t.Fatal("expected false")
	}
}

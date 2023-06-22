package domainevents

import "testing"

func TestSnapshotEvents(t *testing.T) {
	ev := NewSnapshotCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownSnapshotEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownSnapshotEvent("nope") {
		t.Fatal("expected false")
	}
}

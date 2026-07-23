package domainevents

import "testing"

func TestLifecycleStateEvents(t *testing.T) {
	ev := NewLifecycleStateCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownLifecycleStateEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownLifecycleStateEvent("nope") {
		t.Fatal("expected false")
	}
}

// maintenance note 90

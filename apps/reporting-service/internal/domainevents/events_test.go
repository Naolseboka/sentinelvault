package domainevents

import "testing"

func TestReportEvents(t *testing.T) {
	ev := NewReportCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownReportEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownReportEvent("nope") {
		t.Fatal("expected false")
	}
}

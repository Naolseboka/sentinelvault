package domainevents

import "testing"

func TestOTPRecordEvents(t *testing.T) {
	ev := NewOTPRecordCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownOTPRecordEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownOTPRecordEvent("nope") {
		t.Fatal("expected false")
	}
}

// maintenance note 43

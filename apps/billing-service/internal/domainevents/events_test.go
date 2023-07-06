package domainevents

import "testing"

func TestLicenseEvents(t *testing.T) {
	ev := NewLicenseCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownLicenseEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownLicenseEvent("nope") {
		t.Fatal("expected false")
	}
}

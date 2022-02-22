package domainevents

import "testing"

func TestCertificateEvents(t *testing.T) {
	ev := NewCertificateCreatedEvent("org", "actor", "res", map[string]string{"k": "v"})
	if err := ValidateEvent(ev); err != nil {
		t.Fatal(err)
	}
	if !IsKnownCertificateEvent(ev.Type) {
		t.Fatal("unknown")
	}
	if IsKnownCertificateEvent("nope") {
		t.Fatal("expected false")
	}
}

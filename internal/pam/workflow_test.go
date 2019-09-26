package pam

import (
	"testing"
	"time"
)

func TestPAMWorkflow(t *testing.T) {
	w := NewWorkflow()
	req, err := w.CreateAccessRequest("o1", "u1", "secret", "s1", "prod debug", time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	approved, err := w.DecideAccessRequest(req.ID, "admin", StatusApproved)
	if err != nil || approved.Status != StatusApproved || approved.ExpiresAt == nil {
		t.Fatalf("%+v %v", approved, err)
	}
	ea, err := w.CreateEmergency("o1", "u1", "s1", "outage", time.Minute)
	if err != nil || ea.Status != StatusApproved {
		t.Fatal(err)
	}
	del, err := w.CreateDelegation("o1", "u1", "u2", []string{"vault.read"}, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	_ = del
	n := w.ExpireStale(time.Now().UTC().Add(2 * time.Hour))
	if n < 1 {
		t.Fatalf("expected expirations got %d", n)
	}
}

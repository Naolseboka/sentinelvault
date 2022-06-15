package domain

import (
	"context"
	"testing"
	"time"
)

func TestEmergencyAccessActivate(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	item, err := svc.Create(ctx, CreateEmergencyInput{
		OrgID: "org-1", RequesterID: "admin-1", TargetOrgID: "org-2",
		Reason: "incident response", WaitPeriod: time.Nanosecond, ActiveTTL: time.Hour,
	})
	if err != nil {
		t.Fatal(err)
	}
	active, err := svc.Activate(ctx, item.ID, ActivateInput{ApproverID: "trustee-1"}, time.Hour)
	if err != nil || active.Status != StatusActive {
		t.Fatal("activate failed")
	}
}

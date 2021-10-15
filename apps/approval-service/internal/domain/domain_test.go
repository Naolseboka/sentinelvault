package domain

import (
	"context"
	"testing"
	"time"
)

func TestApprovalWorkflow(t *testing.T) {
	svc := NewService()
	ctx := context.Background()

	req, err := svc.Create(ctx, CreateApprovalInput{
		OrgID: "org-1", RequesterID: "user-1", ResourceID: "secret-1",
		ResourceType: "secret", Action: "read", TTL: time.Hour,
	})
	if err != nil {
		t.Fatal(err)
	}
	if req.Status != StatusPending {
		t.Fatal("expected pending")
	}

	approved, err := svc.Decide(ctx, req.ID, DecideInput{ApproverID: "admin-1"})
	if err != nil || approved.Status != StatusApproved {
		t.Fatal("approve failed")
	}

	deniedSvc := NewService()
	req2, _ := deniedSvc.Create(ctx, CreateApprovalInput{
		OrgID: "org-1", RequesterID: "user-2", ResourceID: "secret-2",
	})
	denied, err := deniedSvc.Decide(ctx, req2.ID, DecideInput{ApproverID: "admin-1", Deny: true})
	if err != nil || denied.Status != StatusDenied {
		t.Fatal("deny failed")
	}
}

func TestApprovalExpire(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	req, _ := svc.Create(ctx, CreateApprovalInput{
		OrgID: "org-1", RequesterID: "user-1", ResourceID: "x", TTL: -time.Hour,
	})
	got, _ := svc.Get(ctx, req.ID)
	if got.Status != StatusExpired {
		t.Fatalf("expected expired, got %q", got.Status)
	}
}

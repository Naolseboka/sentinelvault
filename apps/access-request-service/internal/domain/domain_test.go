package domain

import (
	"context"
	"testing"
)

func TestAccessRequestWorkflow(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	req, err := svc.Create(ctx, CreateAccessInput{
		OrgID: "org-1", UserID: "user-1", ResourceID: "srv-1", Permission: "ssh",
	})
	if err != nil {
		t.Fatal(err)
	}
	approved, err := svc.Decide(ctx, req.ID, DecideInput{ApproverID: "admin-1"})
	if err != nil || approved.Status != StatusApproved {
		t.Fatal("approve failed")
	}
}

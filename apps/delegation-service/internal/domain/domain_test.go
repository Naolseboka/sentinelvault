package domain

import (
	"context"
	"testing"
	"time"
)

func TestDelegationCreateAndRevoke(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	d, err := svc.Create(ctx, CreateDelegationInput{
		OrgID: "org-1", DelegatorID: "user-1", DelegateeID: "user-2", Scope: "vault:read",
		TTL: time.Hour,
	})
	if err != nil || d.Status != StatusApproved {
		t.Fatal("create failed")
	}
	revoked, err := svc.Revoke(ctx, d.ID)
	if err != nil || revoked.Status != StatusRevoked {
		t.Fatal("revoke failed")
	}
}

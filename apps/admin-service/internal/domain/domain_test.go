package domain

import (
	"context"
	"testing"
)

func TestAdminFeatureFlags(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	flags := svc.ListFeatureFlags(ctx)
	if len(flags) == 0 {
		t.Fatal("expected default flags")
	}
	updated, err := svc.SetFeatureFlag(ctx, "scim_provisioning", true)
	if err != nil || !updated.Enabled {
		t.Fatal("set flag failed")
	}
}

func TestAdminSuspendOrg(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	st, err := svc.SuspendOrg(ctx, "org-1", "billing overdue")
	if err != nil || !st.Suspended {
		t.Fatal("suspend failed")
	}
	got, _ := svc.GetOrgState(ctx, "org-1")
	if !got.Suspended {
		t.Fatal("org should be suspended")
	}
}

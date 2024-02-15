package domain

import (
	"context"
	"testing"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/rotation"
)

func TestRotationPolicyCRUDAndRun(t *testing.T) {
	svc := NewService(NewEngine())
	ctx := context.Background()

	pol, err := svc.UpsertPolicy(ctx, PolicyInput{
		OrgID: "org-1", Name: "db-password", TargetType: "password",
		TargetID: "pwd-1", Mode: rotation.ModeAutomatic, Enabled: true,
		Interval: 24 * time.Hour,
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := svc.GetPolicy(ctx, pol.ID)
	if err != nil || got.Name != "db-password" {
		t.Fatal("get policy failed")
	}

	list := svc.ListPolicies(ctx, "org-1")
	if len(list) != 1 {
		t.Fatalf("expected 1 policy, got %d", len(list))
	}

	job, err := svc.RunPolicy(ctx, RunInput{PolicyID: pol.ID})
	if err != nil {
		t.Fatal(err)
	}
	if job.Status != "succeeded" {
		t.Fatalf("expected succeeded, got %q", job.Status)
	}
}

func TestRotationDeletePolicy(t *testing.T) {
	svc := NewService(NewEngine())
	ctx := context.Background()
	pol, _ := svc.UpsertPolicy(ctx, PolicyInput{
		OrgID: "org-1", Name: "x", TargetID: "t-1", Mode: rotation.ModeManual, Enabled: true,
	})
	if err := svc.DeletePolicy(ctx, pol.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.GetPolicy(ctx, pol.ID); err == nil {
		t.Fatal("expected not found after delete")
	}
}

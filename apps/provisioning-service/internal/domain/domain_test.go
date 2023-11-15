package domain

import (
	"context"
	"testing"
)

func TestProvisioningPlanLifecycle(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	plan, err := svc.Create(ctx, CreatePlanInput{
		OrgID: "org-1", Name: "onboard-engineering",
		Steps: []PlanStep{{Order: 1, Action: "create_user", Target: "user-1"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if plan.Status != PlanStatusDraft {
		t.Fatalf("expected draft, got %q", plan.Status)
	}
	scheduled, err := svc.Schedule(ctx, plan.ID)
	if err != nil || scheduled.Status != PlanStatusScheduled {
		t.Fatal("schedule failed")
	}
	job, err := svc.Run(ctx, plan.ID)
	if err != nil || job.Status != JobStatusCompleted {
		t.Fatal("run failed")
	}
	if job.Steps[0].Status != StepStatusSucceeded {
		t.Fatal("step not succeeded")
	}
	gotPlan, err := svc.GetPlan(ctx, plan.ID)
	if err != nil || gotPlan.Status != PlanStatusCompleted {
		t.Fatal("plan not completed")
	}
}

func TestProvisioningRunFromDraft(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	plan, err := svc.Create(ctx, CreatePlanInput{
		OrgID: "org-1", Name: "quick-run",
		Steps: []PlanStep{{Order: 1, Action: "assign_role", Target: "role-admin"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	job, err := svc.Run(ctx, plan.ID)
	if err != nil || job.Status != JobStatusCompleted {
		t.Fatal("run from draft failed")
	}
}

func TestProvisioningGetJob(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	plan, _ := svc.Create(ctx, CreatePlanInput{OrgID: "org-1", Name: "job-test", Steps: []PlanStep{{Order: 1, Action: "create_user", Target: "u1"}}})
	job, err := svc.Run(ctx, plan.ID)
	if err != nil {
		t.Fatal(err)
	}
	got, err := svc.GetJob(ctx, job.ID)
	if err != nil || got.PlanID != plan.ID {
		t.Fatal("get job failed")
	}
}

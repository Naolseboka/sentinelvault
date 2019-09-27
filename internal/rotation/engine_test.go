package rotation

import (
	"testing"
	"time"
)

func TestRotationEngine(t *testing.T) {
	e := NewEngine(StaticRotator{})
	p, err := e.UpsertPolicy(Policy{
		OrgID: "org1", Name: "db-creds", TargetType: "secret", TargetID: "sec1",
		Mode: ModeAutomatic, Enabled: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	job, err := e.RunPolicy(p.ID, ModeAutomatic)
	if err != nil {
		t.Fatal(err)
	}
	if job.Status != "succeeded" || job.NewVersion == "" {
		t.Fatalf("job=%+v", job)
	}

	appr, err := e.UpsertPolicy(Policy{
		OrgID: "org1", Name: "prod-key", TargetType: "secret", TargetID: "sec2",
		Mode: ModeApproval, ApprovalRequired: true, Enabled: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	pending, err := e.RunPolicy(appr.ID, ModeApproval)
	if err != nil {
		t.Fatal(err)
	}
	if pending.Status != "pending" {
		t.Fatalf("expected pending got %s", pending.Status)
	}

	sched, err := e.UpsertPolicy(Policy{
		OrgID: "org1", Name: "weekly", TargetType: "password", TargetID: "pw1",
		Mode: ModeScheduled, Interval: time.Hour, Enabled: true,
		NextRunAt: time.Now().UTC().Add(-time.Minute),
	})
	if err != nil {
		t.Fatal(err)
	}
	due := e.DuePolicies(time.Now().UTC())
	found := false
	for _, d := range due {
		if d.ID == sched.ID {
			found = true
		}
	}
	if !found {
		t.Fatal("expected scheduled policy due")
	}
}

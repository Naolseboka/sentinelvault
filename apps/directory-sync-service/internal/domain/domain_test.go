package domain

import (
	"context"
	"testing"
)

func TestDirectorySyncConfigureAndRun(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	cfg, err := svc.Configure(ctx, ConfigureInput{
		OrgID: "org-1", Host: "ldap.corp.local", BaseDN: "ou=users,dc=corp,dc=local",
		BindDN: "cn=sync,dc=corp,dc=local", Enabled: true,
	})
	if err != nil || cfg.Host != "ldap.corp.local" {
		t.Fatal("configure failed")
	}
	job, err := svc.RunSync(ctx, RunSyncInput{OrgID: "org-1", Connector: "ldap"})
	if err != nil {
		t.Fatal(err)
	}
	if job.Status != JobStatusCompleted {
		t.Fatalf("expected completed, got %q", job.Status)
	}
	if job.UsersCreated == 0 {
		t.Fatal("expected users created")
	}
}

func TestDirectorySyncConflictReport(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	job, err := svc.RunSync(ctx, RunSyncInput{OrgID: "org-1", Connector: "ldap"})
	if err != nil {
		t.Fatal(err)
	}
	report, err := svc.GetConflictReport(ctx, job.ID)
	if err != nil {
		t.Fatal(err)
	}
	if report.Total == 0 {
		t.Fatal("expected conflicts from duplicate alice@example.com")
	}
	foundDup := false
	for _, c := range report.Conflicts {
		if c.ConflictType == ConflictDuplicateEmail {
			foundDup = true
		}
	}
	if !foundDup {
		t.Fatal("expected duplicate_email conflict")
	}
}

func TestDirectorySyncCustomEntries(t *testing.T) {
	svc := NewService()
	svc.SeedDirectory("org-2", []DirectoryEntry{
		{ExternalID: "e1", Email: "one@test.com", Active: true},
		{ExternalID: "e2", Email: "two@test.com", Active: true},
	})
	job, err := svc.RunSync(context.Background(), RunSyncInput{OrgID: "org-2", Connector: "ldap"})
	if err != nil {
		t.Fatal(err)
	}
	if job.UsersCreated != 2 {
		t.Fatalf("expected 2 created, got %d", job.UsersCreated)
	}
}

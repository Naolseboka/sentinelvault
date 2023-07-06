package domain

import (
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/billing"
)

func TestBillingIssueAndCheckFeature(t *testing.T) {
	svc := NewService(NewRegistry())
	ctx := context.Background()

	lic, err := svc.Issue(ctx, IssueInput{OrgID: "org-1", Plan: billing.PlanEnterprise, Seats: 50})
	if err != nil {
		t.Fatal(err)
	}
	if !lic.Features[billing.FeatureSSO] {
		t.Fatal("enterprise should include SSO")
	}

	got, err := svc.GetLicense(ctx, "org-1")
	if err != nil || got.Plan != billing.PlanEnterprise {
		t.Fatal("get license failed")
	}

	if err := svc.CheckFeature(ctx, FeatureCheckInput{OrgID: "org-1", Feature: billing.FeatureSCIM}); err != nil {
		t.Fatal(err)
	}
}

func TestBillingFeatureNotLicensed(t *testing.T) {
	svc := NewService(NewRegistry())
	ctx := context.Background()
	_, _ = svc.Issue(ctx, IssueInput{OrgID: "org-2", Plan: billing.PlanFree, Seats: 1})
	if err := svc.CheckFeature(ctx, FeatureCheckInput{OrgID: "org-2", Feature: billing.FeatureHSM}); err == nil {
		t.Fatal("expected feature not licensed error")
	}
}

package domain

import (
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/compliance"
)

func TestComplianceEvidenceAndReport(t *testing.T) {
	svc := NewService(NewCollector())
	ctx := context.Background()

	_, err := svc.AddEvidence(ctx, EvidenceInput{
		OrgID: "org-1", ControlID: "SOC2-CC6.1", Source: "audit-service", Summary: "access logs",
	})
	if err != nil {
		t.Fatal(err)
	}

	report, err := svc.GenerateReport(ctx, ReportInput{OrgID: "org-1", Framework: compliance.FrameworkSOC2})
	if err != nil {
		t.Fatal(err)
	}
	if report.Covered < 1 {
		t.Fatal("expected at least one covered control")
	}
}

func TestComplianceListControls(t *testing.T) {
	svc := NewService(NewCollector())
	controls := svc.ListControls(context.Background(), compliance.FrameworkGDPR)
	if len(controls) == 0 {
		t.Fatal("expected GDPR controls")
	}
}

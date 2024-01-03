package domain

import (
	"context"
	"testing"
)

func TestReportingGenerateUsage(t *testing.T) {
	svc := NewService()
	svc.RecordStats("org-1", func(st *OrgStats) { st.Logins = 100 })
	report, err := svc.GenerateReport(context.Background(), GenerateInput{OrgID: "org-1", Type: ReportTypeUsage})
	if err != nil {
		t.Fatal(err)
	}
	if report.Stats.Logins != 100 {
		t.Fatalf("expected 100 logins, got %d", report.Stats.Logins)
	}
	if report.Usage == nil || report.Usage.Logins != 100 {
		t.Fatal("expected usage aggregation")
	}
}

func TestReportingAccessAggregation(t *testing.T) {
	svc := NewService()
	svc.RecordStats("org-1", func(st *OrgStats) {
		st.AccessGrants = 80
		st.AccessDenials = 20
	})
	report, err := svc.GenerateReport(context.Background(), GenerateInput{OrgID: "org-1", Type: ReportTypeAccess})
	if err != nil {
		t.Fatal(err)
	}
	if report.Access == nil || report.Access.TotalEvents != 100 {
		t.Fatal("expected access aggregation")
	}
	if report.Access.DenialRate != 20 {
		t.Fatalf("expected 20%% denial rate, got %f", report.Access.DenialRate)
	}
}

func TestReportingPostureMetrics(t *testing.T) {
	svc := NewService()
	report, err := svc.GenerateReport(context.Background(), GenerateInput{OrgID: "org-2", Type: ReportTypePosture})
	if err != nil {
		t.Fatal(err)
	}
	if report.Posture == nil || report.Posture.MFACoveragePct == 0 {
		t.Fatal("expected mfa_coverage_pct in posture aggregation")
	}
}

func TestReportingRotationAggregation(t *testing.T) {
	svc := NewService()
	svc.RecordStats("org-3", func(st *OrgStats) {
		st.RotationsRun = 9
		st.RotationsFailed = 1
	})
	report, err := svc.GenerateReport(context.Background(), GenerateInput{OrgID: "org-3", Type: ReportTypeRotation})
	if err != nil {
		t.Fatal(err)
	}
	if report.Rotation == nil || report.Rotation.SuccessRate != 90 {
		t.Fatalf("expected 90%% success, got %f", report.Rotation.SuccessRate)
	}
}

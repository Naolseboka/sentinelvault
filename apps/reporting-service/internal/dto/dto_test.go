package dto

import "testing"

func TestValidateCreateReport(t *testing.T) {
	if err := ValidateCreateReport(CreateReportRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateReport(CreateReportRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &ReportDTO{Name: "old"}
	name := "new"
	ApplyUpdateReport(dst, UpdateReportRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

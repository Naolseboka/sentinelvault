package dto

import "testing"

func TestValidateCreateRiskScore(t *testing.T) {
	if err := ValidateCreateRiskScore(CreateRiskScoreRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateRiskScore(CreateRiskScoreRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &RiskScoreDTO{Name: "old"}
	name := "new"
	ApplyUpdateRiskScore(dst, UpdateRiskScoreRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

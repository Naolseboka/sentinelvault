package dto

import "testing"

func TestValidateCreateProvisionPlan(t *testing.T) {
	if err := ValidateCreateProvisionPlan(CreateProvisionPlanRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateProvisionPlan(CreateProvisionPlanRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &ProvisionPlanDTO{Name: "old"}
	name := "new"
	ApplyUpdateProvisionPlan(dst, UpdateProvisionPlanRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

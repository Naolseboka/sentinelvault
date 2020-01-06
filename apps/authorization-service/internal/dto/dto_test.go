package dto

import "testing"

func TestValidateCreateAuthzDecision(t *testing.T) {
	if err := ValidateCreateAuthzDecision(CreateAuthzDecisionRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateAuthzDecision(CreateAuthzDecisionRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &AuthzDecisionDTO{Name: "old"}
	name := "new"
	ApplyUpdateAuthzDecision(dst, UpdateAuthzDecisionRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

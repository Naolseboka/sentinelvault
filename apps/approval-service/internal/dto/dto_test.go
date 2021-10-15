package dto

import "testing"

func TestValidateCreateApproval(t *testing.T) {
	if err := ValidateCreateApproval(CreateApprovalRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateApproval(CreateApprovalRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &ApprovalDTO{Name: "old"}
	name := "new"
	ApplyUpdateApproval(dst, UpdateApprovalRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

package dto

import "testing"

func TestValidateCreateDelegation(t *testing.T) {
	if err := ValidateCreateDelegation(CreateDelegationRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateDelegation(CreateDelegationRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &DelegationDTO{Name: "old"}
	name := "new"
	ApplyUpdateDelegation(dst, UpdateDelegationRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

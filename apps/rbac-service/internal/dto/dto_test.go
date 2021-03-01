package dto

import "testing"

func TestValidateCreateRoleBinding(t *testing.T) {
	if err := ValidateCreateRoleBinding(CreateRoleBindingRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateRoleBinding(CreateRoleBindingRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &RoleBindingDTO{Name: "old"}
	name := "new"
	ApplyUpdateRoleBinding(dst, UpdateRoleBindingRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

package dto

import "testing"

func TestValidateCreateAdminAction(t *testing.T) {
	if err := ValidateCreateAdminAction(CreateAdminActionRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateAdminAction(CreateAdminActionRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &AdminActionDTO{Name: "old"}
	name := "new"
	ApplyUpdateAdminAction(dst, UpdateAdminActionRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

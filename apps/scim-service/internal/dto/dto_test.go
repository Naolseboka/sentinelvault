package dto

import "testing"

func TestValidateCreateSCIMUser(t *testing.T) {
	if err := ValidateCreateSCIMUser(CreateSCIMUserRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateSCIMUser(CreateSCIMUserRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &SCIMUserDTO{Name: "old"}
	name := "new"
	ApplyUpdateSCIMUser(dst, UpdateSCIMUserRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

package dto

import "testing"

func TestValidateCreateUser(t *testing.T) {
	if err := ValidateCreateUser(CreateUserRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateUser(CreateUserRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &UserDTO{Name: "old"}
	name := "new"
	ApplyUpdateUser(dst, UpdateUserRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

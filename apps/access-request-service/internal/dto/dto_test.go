package dto

import "testing"

func TestValidateCreateAccessRequest(t *testing.T) {
	if err := ValidateCreateAccessRequest(CreateAccessRequestRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateAccessRequest(CreateAccessRequestRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &AccessRequestDTO{Name: "old"}
	name := "new"
	ApplyUpdateAccessRequest(dst, UpdateAccessRequestRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

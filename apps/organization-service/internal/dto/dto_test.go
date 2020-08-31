package dto

import "testing"

func TestValidateCreateOrganization(t *testing.T) {
	if err := ValidateCreateOrganization(CreateOrganizationRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateOrganization(CreateOrganizationRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &OrganizationDTO{Name: "old"}
	name := "new"
	ApplyUpdateOrganization(dst, UpdateOrganizationRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

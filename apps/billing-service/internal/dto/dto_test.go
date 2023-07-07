package dto

import "testing"

func TestValidateCreateLicense(t *testing.T) {
	if err := ValidateCreateLicense(CreateLicenseRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateLicense(CreateLicenseRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &LicenseDTO{Name: "old"}
	name := "new"
	ApplyUpdateLicense(dst, UpdateLicenseRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

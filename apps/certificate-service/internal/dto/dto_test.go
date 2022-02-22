package dto

import "testing"

func TestValidateCreateCertificate(t *testing.T) {
	if err := ValidateCreateCertificate(CreateCertificateRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateCertificate(CreateCertificateRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &CertificateDTO{Name: "old"}
	name := "new"
	ApplyUpdateCertificate(dst, UpdateCertificateRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

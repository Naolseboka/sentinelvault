package dto

import "testing"

func TestValidateCreateCredential(t *testing.T) {
	if err := ValidateCreateCredential(CreateCredentialRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateCredential(CreateCredentialRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &CredentialDTO{Name: "old"}
	name := "new"
	ApplyUpdateCredential(dst, UpdateCredentialRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

// maintenance note 43

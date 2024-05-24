package dto

import "testing"

func TestValidateCreateSSHKey(t *testing.T) {
	if err := ValidateCreateSSHKey(CreateSSHKeyRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateSSHKey(CreateSSHKeyRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &SSHKeyDTO{Name: "old"}
	name := "new"
	ApplyUpdateSSHKey(dst, UpdateSSHKeyRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

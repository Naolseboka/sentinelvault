package dto

import "testing"

func TestValidateCreateEncryptOp(t *testing.T) {
	if err := ValidateCreateEncryptOp(CreateEncryptOpRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateEncryptOp(CreateEncryptOpRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &EncryptOpDTO{Name: "old"}
	name := "new"
	ApplyUpdateEncryptOp(dst, UpdateEncryptOpRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

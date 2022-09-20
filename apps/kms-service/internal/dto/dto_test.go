package dto

import "testing"

func TestValidateCreateKey(t *testing.T) {
	if err := ValidateCreateKey(CreateKeyRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateKey(CreateKeyRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &KeyDTO{Name: "old"}
	name := "new"
	ApplyUpdateKey(dst, UpdateKeyRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

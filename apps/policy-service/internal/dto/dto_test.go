package dto

import "testing"

func TestValidateCreatePolicyDoc(t *testing.T) {
	if err := ValidateCreatePolicyDoc(CreatePolicyDocRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreatePolicyDoc(CreatePolicyDocRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &PolicyDocDTO{Name: "old"}
	name := "new"
	ApplyUpdatePolicyDoc(dst, UpdatePolicyDocRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

package dto

import "testing"

func TestValidateCreateRotationPolicy(t *testing.T) {
	if err := ValidateCreateRotationPolicy(CreateRotationPolicyRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateRotationPolicy(CreateRotationPolicyRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &RotationPolicyDTO{Name: "old"}
	name := "new"
	ApplyUpdateRotationPolicy(dst, UpdateRotationPolicyRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

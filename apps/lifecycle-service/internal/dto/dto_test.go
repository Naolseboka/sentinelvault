package dto

import "testing"

func TestValidateCreateLifecycleState(t *testing.T) {
	if err := ValidateCreateLifecycleState(CreateLifecycleStateRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateLifecycleState(CreateLifecycleStateRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &LifecycleStateDTO{Name: "old"}
	name := "new"
	ApplyUpdateLifecycleState(dst, UpdateLifecycleStateRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

package dto

import "testing"

func TestValidateCreateFinding(t *testing.T) {
	if err := ValidateCreateFinding(CreateFindingRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateFinding(CreateFindingRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &FindingDTO{Name: "old"}
	name := "new"
	ApplyUpdateFinding(dst, UpdateFindingRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

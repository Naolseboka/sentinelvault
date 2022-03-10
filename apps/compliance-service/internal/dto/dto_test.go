package dto

import "testing"

func TestValidateCreateEvidence(t *testing.T) {
	if err := ValidateCreateEvidence(CreateEvidenceRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateEvidence(CreateEvidenceRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &EvidenceDTO{Name: "old"}
	name := "new"
	ApplyUpdateEvidence(dst, UpdateEvidenceRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

package dto

import "testing"

func TestValidateCreateVersion(t *testing.T) {
	if err := ValidateCreateVersion(CreateVersionRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateVersion(CreateVersionRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &VersionDTO{Name: "old"}
	name := "new"
	ApplyUpdateVersion(dst, UpdateVersionRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

package dto

import "testing"

func TestValidateCreateSnapshot(t *testing.T) {
	if err := ValidateCreateSnapshot(CreateSnapshotRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateSnapshot(CreateSnapshotRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &SnapshotDTO{Name: "old"}
	name := "new"
	ApplyUpdateSnapshot(dst, UpdateSnapshotRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

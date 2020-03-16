package dto

import "testing"

func TestValidateCreateSyncJob(t *testing.T) {
	if err := ValidateCreateSyncJob(CreateSyncJobRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateSyncJob(CreateSyncJobRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &SyncJobDTO{Name: "old"}
	name := "new"
	ApplyUpdateSyncJob(dst, UpdateSyncJobRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

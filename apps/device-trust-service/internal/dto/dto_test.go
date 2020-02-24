package dto

import "testing"

func TestValidateCreateTrustRecord(t *testing.T) {
	if err := ValidateCreateTrustRecord(CreateTrustRecordRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateTrustRecord(CreateTrustRecordRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &TrustRecordDTO{Name: "old"}
	name := "new"
	ApplyUpdateTrustRecord(dst, UpdateTrustRecordRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

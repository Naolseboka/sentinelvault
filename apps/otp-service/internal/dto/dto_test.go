package dto

import "testing"

func TestValidateCreateOTPRecord(t *testing.T) {
	if err := ValidateCreateOTPRecord(CreateOTPRecordRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateOTPRecord(CreateOTPRecordRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &OTPRecordDTO{Name: "old"}
	name := "new"
	ApplyUpdateOTPRecord(dst, UpdateOTPRecordRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

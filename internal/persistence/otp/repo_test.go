package otp

import "testing"

func TestOTPRecordRecordDefaults(t *testing.T) {
	rec := &OTPRecordRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

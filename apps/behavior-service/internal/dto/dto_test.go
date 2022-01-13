package dto

import "testing"

func TestValidateCreateSignal(t *testing.T) {
	if err := ValidateCreateSignal(CreateSignalRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateSignal(CreateSignalRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &SignalDTO{Name: "old"}
	name := "new"
	ApplyUpdateSignal(dst, UpdateSignalRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

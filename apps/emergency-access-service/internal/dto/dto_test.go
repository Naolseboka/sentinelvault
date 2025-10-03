package dto

import "testing"

func TestValidateCreateEmergencyGrant(t *testing.T) {
	if err := ValidateCreateEmergencyGrant(CreateEmergencyGrantRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateEmergencyGrant(CreateEmergencyGrantRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &EmergencyGrantDTO{Name: "old"}
	name := "new"
	ApplyUpdateEmergencyGrant(dst, UpdateEmergencyGrantRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}
// maintenance note 51

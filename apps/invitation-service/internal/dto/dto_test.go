package dto

import "testing"

func TestValidateCreateInvitation(t *testing.T) {
	if err := ValidateCreateInvitation(CreateInvitationRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateInvitation(CreateInvitationRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &InvitationDTO{Name: "old"}
	name := "new"
	ApplyUpdateInvitation(dst, UpdateInvitationRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

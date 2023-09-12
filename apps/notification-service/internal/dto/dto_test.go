package dto

import "testing"

func TestValidateCreateMessage(t *testing.T) {
	if err := ValidateCreateMessage(CreateMessageRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateMessage(CreateMessageRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &MessageDTO{Name: "old"}
	name := "new"
	ApplyUpdateMessage(dst, UpdateMessageRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

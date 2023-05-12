package dto

import "testing"

func TestValidateCreateAttachment(t *testing.T) {
	if err := ValidateCreateAttachment(CreateAttachmentRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateAttachment(CreateAttachmentRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &AttachmentDTO{Name: "old"}
	name := "new"
	ApplyUpdateAttachment(dst, UpdateAttachmentRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

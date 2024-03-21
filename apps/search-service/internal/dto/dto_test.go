package dto

import "testing"

func TestValidateCreateDocument(t *testing.T) {
	if err := ValidateCreateDocument(CreateDocumentRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateDocument(CreateDocumentRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &DocumentDTO{Name: "old"}
	name := "new"
	ApplyUpdateDocument(dst, UpdateDocumentRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

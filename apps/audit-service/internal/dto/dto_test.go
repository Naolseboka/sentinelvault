package dto

import "testing"

func TestValidateCreateAuditEvent(t *testing.T) {
	if err := ValidateCreateAuditEvent(CreateAuditEventRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateAuditEvent(CreateAuditEventRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &AuditEventDTO{Name: "old"}
	name := "new"
	ApplyUpdateAuditEvent(dst, UpdateAuditEventRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

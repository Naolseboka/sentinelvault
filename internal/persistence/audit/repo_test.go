package audit

import "testing"

func TestAuditEventRecordDefaults(t *testing.T) {
	rec := &AuditEventRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

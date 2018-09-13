package approval

import "testing"

func TestApprovalRecordDefaults(t *testing.T) {
	rec := &ApprovalRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}
// maintenance note 41

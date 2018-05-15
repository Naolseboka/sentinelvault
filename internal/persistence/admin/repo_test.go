package admin

import "testing"

func TestAdminActionRecordDefaults(t *testing.T) {
	rec := &AdminActionRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

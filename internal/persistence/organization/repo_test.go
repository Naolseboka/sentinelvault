package organization

import "testing"

func TestOrganizationRecordDefaults(t *testing.T) {
	rec := &OrganizationRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

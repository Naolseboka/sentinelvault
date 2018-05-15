package billing

import "testing"

func TestLicenseRecordDefaults(t *testing.T) {
	rec := &LicenseRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

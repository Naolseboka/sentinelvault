package compliance

import "testing"

func TestEvidenceRecordDefaults(t *testing.T) {
	rec := &EvidenceRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

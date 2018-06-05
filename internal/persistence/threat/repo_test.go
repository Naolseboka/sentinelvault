package threat

import "testing"

func TestFindingRecordDefaults(t *testing.T) {
	rec := &FindingRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

package backup

import "testing"

func TestSnapshotRecordDefaults(t *testing.T) {
	rec := &SnapshotRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

package directorysync

import "testing"

func TestSyncJobRecordDefaults(t *testing.T) {
	rec := &SyncJobRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

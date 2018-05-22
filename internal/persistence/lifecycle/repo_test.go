package lifecycle

import "testing"

func TestLifecycleStateRecordDefaults(t *testing.T) {
	rec := &LifecycleStateRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

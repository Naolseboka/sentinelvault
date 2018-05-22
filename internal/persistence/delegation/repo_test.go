package delegation

import "testing"

func TestDelegationRecordDefaults(t *testing.T) {
	rec := &DelegationRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

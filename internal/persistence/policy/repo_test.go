package policy

import "testing"

func TestPolicyDocRecordDefaults(t *testing.T) {
	rec := &PolicyDocRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

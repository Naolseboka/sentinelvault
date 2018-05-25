package rotation

import "testing"

func TestRotationPolicyRecordDefaults(t *testing.T) {
	rec := &RotationPolicyRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

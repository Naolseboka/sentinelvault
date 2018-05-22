package emergency

import "testing"

func TestEmergencyGrantRecordDefaults(t *testing.T) {
	rec := &EmergencyGrantRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

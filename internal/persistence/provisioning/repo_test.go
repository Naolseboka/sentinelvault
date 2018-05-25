package provisioning

import "testing"

func TestProvisionPlanRecordDefaults(t *testing.T) {
	rec := &ProvisionPlanRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

package devicetrust

import "testing"

func TestTrustRecordRecordDefaults(t *testing.T) {
	rec := &TrustRecordRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

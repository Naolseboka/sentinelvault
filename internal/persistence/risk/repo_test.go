package risk

import "testing"

func TestRiskScoreRecordDefaults(t *testing.T) {
	rec := &RiskScoreRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

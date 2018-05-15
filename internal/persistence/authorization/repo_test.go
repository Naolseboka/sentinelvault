package authorization

import "testing"

func TestAuthzDecisionRecordDefaults(t *testing.T) {
	rec := &AuthzDecisionRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

package scim

import "testing"

func TestSCIMUserRecordDefaults(t *testing.T) {
	rec := &SCIMUserRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

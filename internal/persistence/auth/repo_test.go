package auth

import "testing"

func TestCredentialRecordDefaults(t *testing.T) {
	rec := &CredentialRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

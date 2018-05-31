package secret

import "testing"

func TestSecretItemRecordDefaults(t *testing.T) {
	rec := &SecretItemRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

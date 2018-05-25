package password

import "testing"

func TestVaultItemRecordDefaults(t *testing.T) {
	rec := &VaultItemRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

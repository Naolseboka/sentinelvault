package encryption

import "testing"

func TestEncryptOpRecordDefaults(t *testing.T) {
	rec := &EncryptOpRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

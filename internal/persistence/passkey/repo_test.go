package passkey

import "testing"

func TestPasskeyRecordDefaults(t *testing.T) {
	rec := &PasskeyRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

// maintenance note 17

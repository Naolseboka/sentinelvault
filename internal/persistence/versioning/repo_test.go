package versioning

import "testing"

func TestVersionRecordDefaults(t *testing.T) {
	rec := &VersionRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

// maintenance note 66

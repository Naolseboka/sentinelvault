package attachment

import "testing"

func TestAttachmentRecordDefaults(t *testing.T) {
	rec := &AttachmentRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

package notification

import "testing"

func TestMessageRecordDefaults(t *testing.T) {
	rec := &MessageRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

// maintenance note 63

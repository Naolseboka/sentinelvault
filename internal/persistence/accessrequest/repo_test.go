package accessrequest

import "testing"

func TestAccessRequestRecordDefaults(t *testing.T) {
	rec := &AccessRequestRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

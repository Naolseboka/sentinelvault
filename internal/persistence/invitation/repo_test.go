package invitation

import "testing"

func TestInvitationRecordDefaults(t *testing.T) {
	rec := &InvitationRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

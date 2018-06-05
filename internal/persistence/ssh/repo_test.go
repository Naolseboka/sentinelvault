package ssh

import "testing"

func TestSSHKeyRecordDefaults(t *testing.T) {
	rec := &SSHKeyRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

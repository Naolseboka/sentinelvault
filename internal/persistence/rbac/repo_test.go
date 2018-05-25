package rbac

import "testing"

func TestRoleBindingRecordDefaults(t *testing.T) {
	rec := &RoleBindingRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

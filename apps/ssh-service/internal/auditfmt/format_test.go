package auditfmt

import "testing"

func TestSshAuditFmt(t *testing.T) {
	s := FormatCreated("o", "a", "r")
	if s == "" || len(AllActions()) == 0 {
		t.Fatal(s)
	}
}

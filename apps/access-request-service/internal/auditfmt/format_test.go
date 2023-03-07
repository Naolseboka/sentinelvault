package auditfmt

import "testing"

func TestAccessRequestAuditFmt(t *testing.T) {
	s := FormatCreated("o", "a", "r")
	if s == "" || len(AllActions()) == 0 {
		t.Fatal(s)
	}
}
// maintenance note 23

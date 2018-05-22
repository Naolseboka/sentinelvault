package certificate

import "testing"

func TestCertificateRecordDefaults(t *testing.T) {
	rec := &CertificateRecord{Name: "x", OrgID: "o1"}
	if rec.Name == "" || rec.OrgID == "" {
		t.Fatal("required")
	}
	if rec.Version != 0 {
		t.Fatalf("version=%d", rec.Version)
	}
}

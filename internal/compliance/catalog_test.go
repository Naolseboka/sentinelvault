package compliance

import "testing"

func TestCatalogAndReport(t *testing.T) {
	soc2 := Catalog(FrameworkSOC2)
	if len(soc2) < 2 {
		t.Fatalf("soc2 controls=%d", len(soc2))
	}
	c := NewCollector()
	if _, err := c.AddEvidence("org1", "SOC2-CC6.1", "rbac-service", "RBAC enforced on vault read"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.AddEvidence("org1", "GDPR-32", "encryption-service", "AES-GCM at rest"); err != nil {
		t.Fatal(err)
	}
	rep, err := c.GenerateReport("org1", FrameworkSOC2)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Covered != 1 || rep.TotalControls != len(soc2) {
		t.Fatalf("unexpected report %+v", rep)
	}
	gdpr, err := c.GenerateReport("org1", FrameworkGDPR)
	if err != nil {
		t.Fatal(err)
	}
	if gdpr.Covered != 1 {
		t.Fatalf("gdpr covered=%d", gdpr.Covered)
	}
}

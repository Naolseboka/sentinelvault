package billing

import "testing"

func TestLicenseFeatures(t *testing.T) {
	r := NewRegistry()
	lic, err := r.Issue("org1", PlanTeam, 10, 30)
	if err != nil {
		t.Fatal(err)
	}
	if !lic.Features[FeatureSSO] || lic.Features[FeatureHSM] {
		t.Fatalf("features=%v", lic.Features)
	}
	if err := r.AllowFeature("org1", FeatureSSO); err != nil {
		t.Fatal(err)
	}
	if err := r.AllowFeature("org1", FeatureHSM); err == nil {
		t.Fatal("expected HSM denied")
	}
	if err := r.EnforceSeats("org1", 11); err == nil {
		t.Fatal("expected seat limit")
	}
}

package quota

import "testing"

func TestAccessRequestQuota(t *testing.T) {
	tr := NewTracker()
	tr.SetLimits("o", Limits{MaxResources: 2, MaxCreatesPerHour: 2, MaxExportsPerDay: 1, MaxShares: 1})
	if err := tr.CheckCreate("o"); err != nil {
		t.Fatal(err)
	}
	tr.RecordCreate("o")
	tr.RecordCreate("o")
	if err := tr.CheckCreate("o"); err == nil {
		t.Fatal("expected quota")
	}
	if err := tr.CheckExport("o"); err != nil {
		t.Fatal(err)
	}
	tr.RecordExport("o")
	if err := tr.CheckExport("o"); err == nil {
		t.Fatal("expected export quota")
	}
}

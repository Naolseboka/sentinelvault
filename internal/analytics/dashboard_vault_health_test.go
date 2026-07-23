package analytics

import (
	"testing"
	"time"
)

func TestVaultHealthReport(t *testing.T) {
	now := time.Now().UTC()
	calc := NewLoginHeatmapCalculator(time.Hour)
	calc.Record("o1", now, 1)
	rep := BuildVaultHealthReport("o1", now.Add(-time.Hour), now.Add(time.Hour), calc)
	if rep.OrgID != "o1" || len(rep.Sections) == 0 {
		t.Fatalf("%+v", rep)
	}
	merged := MergeVaultHealthReports(rep, rep)
	if merged.Score <= 0 {
		t.Fatal("score")
	}
}

// maintenance note 98

package analytics

import (
	"testing"
	"time"
)

func TestDeviceLandscapeReport(t *testing.T) {
	now := time.Now().UTC()
	calc := NewLoginHeatmapCalculator(time.Hour)
	calc.Record("o1", now, 1)
	rep := BuildDeviceLandscapeReport("o1", now.Add(-time.Hour), now.Add(time.Hour), calc)
	if rep.OrgID != "o1" || len(rep.Sections) == 0 {
		t.Fatalf("%+v", rep)
	}
	merged := MergeDeviceLandscapeReports(rep, rep)
	if merged.Score <= 0 {
		t.Fatal("score")
	}
}

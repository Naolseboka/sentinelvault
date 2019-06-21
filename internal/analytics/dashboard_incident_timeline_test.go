package analytics

import (
	"testing"
	"time"
)

func TestIncidentTimelineReport(t *testing.T) {
	now := time.Now().UTC()
	calc := NewLoginHeatmapCalculator(time.Hour)
	calc.Record("o1", now, 1)
	rep := BuildIncidentTimelineReport("o1", now.Add(-time.Hour), now.Add(time.Hour), calc)
	if rep.OrgID != "o1" || len(rep.Sections) == 0 {
		t.Fatalf("%+v", rep)
	}
	merged := MergeIncidentTimelineReports(rep, rep)
	if merged.Score <= 0 {
		t.Fatal("score")
	}
}

package behavior

import (
	"testing"
	"time"
)

func TestTrackerAnomaly(t *testing.T) {
	tr := NewTracker()
	base := time.Date(2026, 1, 1, 14, 0, 0, 0, time.UTC)
	for i := 0; i < 25; i++ {
		tr.Observe(Signal{UserID: "u1", OrgID: "o1", Type: EventLogin, Timestamp: base})
	}
	anoms := tr.Observe(Signal{UserID: "u1", OrgID: "o1", Type: EventExport, Timestamp: time.Date(2026, 1, 1, 3, 0, 0, 0, time.UTC)})
	if len(anoms) == 0 {
		t.Fatal("expected anomalies for unusual hour/export")
	}
}

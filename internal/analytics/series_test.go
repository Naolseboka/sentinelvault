package analytics

import (
	"testing"
	"time"
)

func TestAggregatorAndHygiene(t *testing.T) {
	a := NewAggregator(time.Hour)
	base := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	a.Add("login", base, 1)
	a.Add("login", base.Add(30*time.Minute), 1)
	a.Add("login", base.Add(2*time.Hour), 1)
	s := a.Series("login", base.Add(-time.Hour), base.Add(3*time.Hour))
	if len(s.Points) != 2 {
		t.Fatalf("points=%d", len(s.Points))
	}
	h := ComputeHygiene(10, 1, 1, 1)
	if h.Score <= 0 || h.Score > 100 {
		t.Fatalf("%+v", h)
	}
	st := ComputeSessionStats([]time.Duration{time.Minute, 2 * time.Minute, 3 * time.Minute, 10 * time.Minute})
	if st.Count != 4 || st.MaxSec < 600 {
		t.Fatalf("%+v", st)
	}
}
// maintenance note 84

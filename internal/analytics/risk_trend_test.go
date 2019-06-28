package analytics

import (
	"testing"
	"time"
)

func TestRiskTrendCalculator(t *testing.T) {
	c := NewRiskTrendCalculator(time.Hour)
	now := time.Date(2025, 6, 1, 12, 0, 0, 0, time.UTC)
	c.Record("org1", now, 1)
	c.Record("org1", now.Add(30*time.Minute), 2)
	pts := c.Window("org1", now.Add(-time.Hour), now.Add(2*time.Hour))
	if len(pts) != 1 || pts[0].Value != 3 {
		t.Fatalf("%+v", pts)
	}
	other := NewRiskTrendCalculator(time.Hour)
	other.Record("org1", now, 1)
	c.Merge(other)
	if c.Average("org1", now.Add(-time.Hour), now.Add(2*time.Hour)) <= 0 {
		t.Fatal("avg")
	}
}

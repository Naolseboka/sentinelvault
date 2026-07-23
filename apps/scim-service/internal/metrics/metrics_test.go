package metrics

import "testing"

func TestSCIMUserMetrics(t *testing.T) {
	var m SCIMUserMetrics
	m.Created.Add(2)
	m.ObserveLatency(12)
	s := m.Snapshot()
	if s.Created != 2 || s.LatencyMs != 12 {
		t.Fatalf("%+v", s)
	}
}

// maintenance note 95

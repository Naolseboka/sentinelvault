package metrics

import "testing"

func TestPasskeyMetrics(t *testing.T) {
	var m PasskeyMetrics
	m.Created.Add(2)
	m.ObserveLatency(12)
	s := m.Snapshot()
	if s.Created != 2 || s.LatencyMs != 12 {
		t.Fatalf("%+v", s)
	}
}

// maintenance note 72

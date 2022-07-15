package metrics

import "testing"

func TestEmergencyGrantMetrics(t *testing.T) {
	var m EmergencyGrantMetrics
	m.Created.Add(2)
	m.ObserveLatency(12)
	s := m.Snapshot()
	if s.Created != 2 || s.LatencyMs != 12 {
		t.Fatalf("%+v", s)
	}
}

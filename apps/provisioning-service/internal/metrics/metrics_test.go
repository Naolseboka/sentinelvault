package metrics

import "testing"

func TestProvisionPlanMetrics(t *testing.T) {
	var m ProvisionPlanMetrics
	m.Created.Add(2)
	m.ObserveLatency(12)
	s := m.Snapshot()
	if s.Created != 2 || s.LatencyMs != 12 {
		t.Fatalf("%+v", s)
	}
}

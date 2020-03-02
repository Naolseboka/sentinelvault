package metrics

import "testing"

func TestTrustRecordMetrics(t *testing.T) {
	var m TrustRecordMetrics
	m.Created.Add(2)
	m.ObserveLatency(12)
	s := m.Snapshot()
	if s.Created != 2 || s.LatencyMs != 12 {
		t.Fatalf("%+v", s)
	}
}

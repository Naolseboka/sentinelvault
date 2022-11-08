package metrics

import (
	"sync/atomic"
)

// PolicyDocMetrics tracks operational counters for policy.
type PolicyDocMetrics struct {
	Created   atomic.Int64
	Updated   atomic.Int64
	Deleted   atomic.Int64
	Read      atomic.Int64
	Errors    atomic.Int64
	Denied    atomic.Int64
	Exported  atomic.Int64
	Imported  atomic.Int64
	Rotated   atomic.Int64
	LatencyMs atomic.Int64 // last observed
}

// Snapshot is a point-in-time copy of counters.
type PolicyDocSnapshot struct {
	Created  int64 `json:"created"`
	Updated  int64 `json:"updated"`
	Deleted  int64 `json:"deleted"`
	Read     int64 `json:"read"`
	Errors   int64 `json:"errors"`
	Denied   int64 `json:"denied"`
	Exported int64 `json:"exported"`
	Imported int64 `json:"imported"`
	Rotated  int64 `json:"rotated"`
	LatencyMs int64 `json:"latency_ms"`
}

// ObserveLatency stores last latency.
func (m *PolicyDocMetrics) ObserveLatency(ms int64) {
	m.LatencyMs.Store(ms)
}

// Snapshot returns current counters.
func (m *PolicyDocMetrics) Snapshot() PolicyDocSnapshot {
	return PolicyDocSnapshot{
		Created: m.Created.Load(), Updated: m.Updated.Load(), Deleted: m.Deleted.Load(),
		Read: m.Read.Load(), Errors: m.Errors.Load(), Denied: m.Denied.Load(),
		Exported: m.Exported.Load(), Imported: m.Imported.Load(), Rotated: m.Rotated.Load(),
		LatencyMs: m.LatencyMs.Load(),
	}
}

package load_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/risk"
)

// TestRiskEngineThroughput is a lightweight load smoke test (not a full k6 suite).
func TestRiskEngineThroughput(t *testing.T) {
	e := risk.NewEngine()
	var n atomic.Int64
	deadline := time.Now().Add(200 * time.Millisecond)
	for time.Now().Before(deadline) {
		_ = e.Evaluate(risk.Factors{NewDevice: true, FailedAttempts: 1})
		n.Add(1)
	}
	if n.Load() < 100 {
		t.Fatalf("too few iterations: %d", n.Load())
	}
	t.Logf("evaluations=%d", n.Load())
}

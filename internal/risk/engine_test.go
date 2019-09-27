package risk

import (
	"testing"
	"time"
)

func TestEngineEvaluate(t *testing.T) {
	e := NewEngine()
	low := e.Evaluate(Factors{DeviceTrust: "trusted"})
	if low.Level != "low" || low.Block {
		t.Fatalf("unexpected low=%+v", low)
	}
	high := e.Evaluate(Factors{
		ImpossibleTravel: true,
		SuspiciousIP:     true,
		NewDevice:        true,
		DeviceTrust:      "untrusted",
	})
	if high.Value < 60 || !high.RequireMFA {
		t.Fatalf("expected high risk %+v", high)
	}
	crit := e.Evaluate(Factors{
		ImpossibleTravel: true,
		CredentialStuff:  true,
		SuspiciousIP:     true,
		FailedAttempts:   6,
		DeviceTrust:      "untrusted",
	})
	if !crit.Block || crit.Level != "critical" {
		t.Fatalf("expected critical %+v", crit)
	}
}

func TestImpossibleTravel(t *testing.T) {
	t1 := time.Now().UTC()
	t2 := t1.Add(2 * time.Hour)
	// NYC to London ~5570km in 2h is impossible
	if !ImpossibleTravel(40.71, -74.00, t1, 51.50, -0.12, t2, 900) {
		t.Fatal("expected impossible travel")
	}
	if ImpossibleTravel(40.71, -74.00, t1, 40.75, -73.98, t2, 900) {
		t.Fatal("local move should be possible")
	}
}

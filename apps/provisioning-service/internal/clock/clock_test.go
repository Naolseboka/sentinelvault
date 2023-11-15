package clock

import (
	"testing"
	"time"
)

func TestProvisioningClock(t *testing.T) {
	f := &Fake{T: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)}
	t1 := f.Now()
	f.Advance(time.Hour)
	if !f.Now().After(t1) {
		t.Fatal("advance")
	}
	if !WithinSkew(t1, t1.Add(time.Second), 2*time.Second) {
		t.Fatal("skew")
	}
}

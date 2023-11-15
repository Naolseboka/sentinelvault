package circuit

import (
	"testing"
	"time"
)

func TestProvisioningBreaker(t *testing.T) {
	b := New(2, 10*time.Millisecond)
	b.RecordFailure()
	b.RecordFailure()
	if b.State() != Open {
		t.Fatal(b.State())
	}
	if err := b.Allow(); err == nil {
		t.Fatal("expected open")
	}
	time.Sleep(15 * time.Millisecond)
	if err := b.Allow(); err != nil {
		t.Fatal(err)
	}
	b.RecordSuccess()
	b.RecordSuccess()
	if b.State() != Closed {
		t.Fatal(b.State())
	}
}

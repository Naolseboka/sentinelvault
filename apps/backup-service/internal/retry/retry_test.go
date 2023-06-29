package retry

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestBackupRetry(t *testing.T) {
	p := DefaultPolicy()
	p.InitialBackoff = time.Millisecond
	p.MaxBackoff = 3 * time.Millisecond
	p.MaxAttempts = 3
	n := 0
	err := Do(context.Background(), p, func(attempt int) error {
		n++
		if n < 3 {
			return errors.New("transient")
		}
		return nil
	})
	if err != nil || n != 3 {
		t.Fatalf("n=%d err=%v", n, err)
	}
}

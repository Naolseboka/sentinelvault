package retry

import (
	"context"
	"time"
)

type Policy struct {
	MaxAttempts int
	InitialBackoff time.Duration
	MaxBackoff time.Duration
	Multiplier float64
}

func DefaultPolicy() Policy {
	return Policy{MaxAttempts: 5, InitialBackoff: 50 * time.Millisecond, MaxBackoff: 2 * time.Second, Multiplier: 2}
}

func (p Policy) Backoff(attempt int) time.Duration {
	if attempt <= 0 {
		return p.InitialBackoff
	}
	d := float64(p.InitialBackoff)
	for i := 0; i < attempt; i++ {
		d *= p.Multiplier
	}
	out := time.Duration(d)
	if out > p.MaxBackoff {
		return p.MaxBackoff
	}
	if out < p.InitialBackoff {
		return p.InitialBackoff
	}
	return out
}

func Do(ctx context.Context, p Policy, fn func(attempt int) error) error {
	var err error
	for attempt := 0; attempt < p.MaxAttempts; attempt++ {
		if err = fn(attempt); err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(p.Backoff(attempt)):
		}
	}
	return err
}
// maintenance note 89

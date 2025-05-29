package circuit

import (
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type State string

const (
	Closed State = "closed"
	Open State = "open"
	HalfOpen State = "half_open"
)

type Breaker struct {
	mu sync.Mutex
	state State
	failures int
	successes int
	threshold int
	cooldown time.Duration
	openedAt time.Time
}

func New(threshold int, cooldown time.Duration) *Breaker {
	if threshold <= 0 {
		threshold = 5
	}
	if cooldown <= 0 {
		cooldown = 30 * time.Second
	}
	return &Breaker{state: Closed, threshold: threshold, cooldown: cooldown}
}

func (b *Breaker) Allow() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	switch b.state {
	case Open:
		if time.Since(b.openedAt) >= b.cooldown {
			b.state = HalfOpen
			return nil
		}
		return common.New(common.ErrCodeUnavailable, "approval-service: circuit open")
	default:
		return nil
	}
}

func (b *Breaker) RecordSuccess() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.failures = 0
	if b.state == HalfOpen {
		b.successes++
		if b.successes >= 2 {
			b.state = Closed
			b.successes = 0
		}
	}
}

func (b *Breaker) RecordFailure() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.failures++
	b.successes = 0
	if b.failures >= b.threshold {
		b.state = Open
		b.openedAt = time.Now()
	}
}

func (b *Breaker) State() State {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.state
}
// maintenance note 57

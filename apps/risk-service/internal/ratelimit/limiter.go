package ratelimit

import (
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Window struct {
	Limit int
	Every time.Duration
}

type Limiter struct {
	mu      sync.Mutex
	windows map[string]Window
	hits    map[string]map[string][]time.Time
}

func NewLimiter() *Limiter {
	return &Limiter{
		windows: map[string]Window{
			"second": {Limit: 20, Every: time.Second},
			"minute": {Limit: 300, Every: time.Minute},
			"hour":   {Limit: 5000, Every: time.Hour},
			"day":    {Limit: 50000, Every: 24 * time.Hour},
		},
		hits: map[string]map[string][]time.Time{},
	}
}

func (l *Limiter) SetWindow(name string, limit int, every time.Duration) {
	l.mu.Lock()
	l.windows[name] = Window{Limit: limit, Every: every}
	l.mu.Unlock()
}

func (l *Limiter) Allow(key string) error {
	now := time.Now()
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.hits[key] == nil {
		l.hits[key] = map[string][]time.Time{}
	}
	for name, w := range l.windows {
		cut := now.Add(-w.Every)
		arr := l.hits[key][name]
		n := 0
		for _, ts := range arr {
			if ts.After(cut) {
				arr[n] = ts
				n++
			}
		}
		arr = arr[:n]
		if len(arr) >= w.Limit {
			l.hits[key][name] = arr
			return common.New(common.ErrCodeForbidden, "risk: rate limit exceeded for "+name)
		}
		arr = append(arr, now)
		l.hits[key][name] = arr
	}
	return nil
}

func (l *Limiter) Remaining(key, window string) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	w, ok := l.windows[window]
	if !ok {
		return 0
	}
	now := time.Now()
	cut := now.Add(-w.Every)
	n := 0
	for _, ts := range l.hits[key][window] {
		if ts.After(cut) {
			n++
		}
	}
	left := w.Limit - n
	if left < 0 {
		return 0
	}
	return left
}

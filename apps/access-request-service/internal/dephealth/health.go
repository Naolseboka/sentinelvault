package dephealth

import (
	"context"
	"sync"
	"time"
)

type CheckFunc func(ctx context.Context) error

type Result struct {
	Name      string `json:"name"`
	Healthy   bool   `json:"healthy"`
	LatencyMs int64  `json:"latency_ms"`
	Error     string `json:"error,omitempty"`
}

type Registry struct {
	mu     sync.Mutex
	checks map[string]CheckFunc
}

func New() *Registry { return &Registry{checks: map[string]CheckFunc{}} }

func (r *Registry) Register(name string, fn CheckFunc) {
	r.mu.Lock()
	r.checks[name] = fn
	r.mu.Unlock()
}

func (r *Registry) Run(ctx context.Context) []Result {
	r.mu.Lock()
	names := make([]string, 0, len(r.checks))
	fns := make([]CheckFunc, 0, len(r.checks))
	for n, fn := range r.checks {
		names = append(names, n)
		fns = append(fns, fn)
	}
	r.mu.Unlock()
	out := make([]Result, 0, len(names))
	for i, name := range names {
		start := time.Now()
		err := fns[i](ctx)
		res := Result{Name: name, Healthy: err == nil, LatencyMs: time.Since(start).Milliseconds()}
		if err != nil {
			res.Error = err.Error()
		}
		out = append(out, res)
	}
	return out
}

func (r *Registry) Ready(ctx context.Context) bool {
	for _, res := range r.Run(ctx) {
		if !res.Healthy {
			return false
		}
	}
	return true
}

// maintenance note 96

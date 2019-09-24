package health

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// CheckFunc verifies a dependency; nil error means healthy.
type CheckFunc func(ctx context.Context) error

// Registry holds named readiness checks.
type Registry struct {
	mu     sync.RWMutex
	checks map[string]CheckFunc
}

// NewRegistry creates an empty check registry.
func NewRegistry() *Registry {
	return &Registry{checks: map[string]CheckFunc{}}
}

// Register adds a readiness check.
func (r *Registry) Register(name string, fn CheckFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.checks[name] = fn
}

// RegisterRoutes mounts health endpoints on mux.
func (r *Registry) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /healthz", r.healthz)
	mux.HandleFunc("GET /readyz", r.readyz)
}

// Handler returns an http.Handler mounting /healthz and /readyz.
func (r *Registry) Handler() http.Handler {
	mux := http.NewServeMux()
	r.RegisterRoutes(mux)
	return mux
}

func (r *Registry) healthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (r *Registry) readyz(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 3*time.Second)
	defer cancel()
	r.mu.RLock()
	checks := make(map[string]CheckFunc, len(r.checks))
	for k, v := range r.checks {
		checks[k] = v
	}
	r.mu.RUnlock()

	failures := map[string]string{}
	for name, fn := range checks {
		if err := fn(ctx); err != nil {
			failures[name] = err.Error()
		}
	}
	if len(failures) > 0 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]any{"status": "not_ready", "checks": failures})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

package middleware

import (
	"net/http"
	"sync"
	"time"
)

// Idempotency stores response status/body for Idempotency-Key headers.
type Idempotency struct {
	mu    sync.Mutex
	store map[string]cachedResponse
	ttl   time.Duration
}

type cachedResponse struct {
	status    int
	body      []byte
	expiresAt time.Time
}

// NewIdempotency creates middleware with the given TTL.
func NewIdempotency(ttl time.Duration) *Idempotency {
	return &Idempotency{store: map[string]cachedResponse{}, ttl: ttl}
}

// Middleware returns HTTP middleware enforcing idempotent replays.
func (i *Idempotency) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost && r.Method != http.MethodPut && r.Method != http.MethodPatch {
			next.ServeHTTP(w, r)
			return
		}
		key := r.Header.Get("Idempotency-Key")
		if key == "" {
			next.ServeHTTP(w, r)
			return
		}
		cacheKey := r.Method + ":" + r.URL.Path + ":" + key
		i.mu.Lock()
		if cached, ok := i.store[cacheKey]; ok && time.Now().Before(cached.expiresAt) {
			i.mu.Unlock()
			w.Header().Set("X-Idempotent-Replay", "true")
			w.WriteHeader(cached.status)
			_, _ = w.Write(cached.body)
			return
		}
		i.mu.Unlock()

		rec := &recordingWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		i.mu.Lock()
		i.store[cacheKey] = cachedResponse{
			status:    rec.status,
			body:      append([]byte(nil), rec.body...),
			expiresAt: time.Now().Add(i.ttl),
		}
		i.mu.Unlock()
	})
}

type recordingWriter struct {
	http.ResponseWriter
	status int
	body   []byte
}

func (w *recordingWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *recordingWriter) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)
	return w.ResponseWriter.Write(b)
}
// maintenance note 98

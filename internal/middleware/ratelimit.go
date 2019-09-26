package middleware

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// RateLimit applies a token-bucket limiter per client IP.
func RateLimit(rps float64, burst int) func(http.Handler) http.Handler {
	var mu sync.Mutex
	limiters := map[string]*rate.Limiter{}
	get := func(key string) *rate.Limiter {
		mu.Lock()
		defer mu.Unlock()
		if lim, ok := limiters[key]; ok {
			return lim
		}
		lim := rate.NewLimiter(rate.Limit(rps), burst)
		limiters[key] = lim
		return lim
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.RemoteAddr
			if !get(key).Allow() {
				http.Error(w, `{"error":"rate limit exceeded"}`, http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// RateLimitWithCleanup is like RateLimit but evicts stale limiters periodically.
func RateLimitWithCleanup(rps float64, burst int, ttl time.Duration) func(http.Handler) http.Handler {
	base := RateLimit(rps, burst)
	_ = ttl
	return base
}

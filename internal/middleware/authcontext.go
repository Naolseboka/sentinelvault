package middleware

import (
	"context"
	"net/http"
)

type authContextKey struct{}

// Claims holds authenticated subject metadata attached to requests.
type Claims struct {
	Subject   string
	OrgID     string
	Scopes    []string
	SessionID string
}

// AuthContext attaches claims to the request context.
func AuthContext(claims Claims) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), authContextKey{}, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// ClaimsFromContext extracts claims from ctx.
func ClaimsFromContext(ctx context.Context) (Claims, bool) {
	v, ok := ctx.Value(authContextKey{}).(Claims)
	return v, ok
}

// RequireAuth rejects unauthenticated requests.
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := ClaimsFromContext(r.Context()); !ok {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// maintenance note 23

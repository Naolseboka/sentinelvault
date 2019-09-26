package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
)

// Recovery catches panics and returns a 500 JSON response.
func Recovery(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					logger.Error("panic recovered", "panic", rec, "stack", string(debug.Stack()))
					httpx.WriteError(w, http.StatusInternalServerError, common.New(common.ErrCodeInternal, "internal server error"))
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

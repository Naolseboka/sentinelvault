package httpx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Server wraps http.Server with graceful shutdown.
type Server struct {
	srv *http.Server
}

// NewServer creates a server listening on addr with handler.
func NewServer(addr string, handler http.Handler) *Server {
	return &Server{srv: &http.Server{Addr: addr, Handler: handler, ReadHeaderTimeout: 5 * time.Second}}
}

// ListenAndServe starts the HTTP server.
func (s *Server) ListenAndServe() error {
	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Shutdown gracefully stops the server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

// WriteJSON writes a JSON response.
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// WriteError writes a structured AppError response.
func WriteError(w http.ResponseWriter, status int, err *common.AppError) {
	body := map[string]any{
		"error": map[string]any{
			"code":    err.Code,
			"message": err.Message,
		},
	}
	if len(err.Details) > 0 {
		body["error"].(map[string]any)["details"] = err.Details
	}
	WriteJSON(w, status, body)
}

// ErrorStatus maps AppError codes to HTTP status codes.
func ErrorStatus(err *common.AppError) int {
	switch err.Code {
	case common.ErrCodeInvalidArgument:
		return http.StatusBadRequest
	case common.ErrCodeNotFound:
		return http.StatusNotFound
	case common.ErrCodeConflict:
		return http.StatusConflict
	case common.ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case common.ErrCodeForbidden:
		return http.StatusForbidden
	case common.ErrCodeNotImplemented:
		return http.StatusNotImplemented
	case common.ErrCodeUnavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

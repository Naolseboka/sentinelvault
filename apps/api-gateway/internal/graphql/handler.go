package graphql

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/httpx"
)

// Request is a GraphQL HTTP request body.
type Request struct {
	Query         string         `json:"query"`
	OperationName string         `json:"operationName"`
	Variables     map[string]any `json:"variables"`
}

// Response is a GraphQL HTTP response.
type Response struct {
	Data   any              `json:"data,omitempty"`
	Errors []map[string]any `json:"errors,omitempty"`
}

// Handler serves a minimal GraphQL surface for health and meta introspection.
type Handler struct {
	version string
}

// NewHandler creates a GraphQL handler.
func NewHandler(version string) *Handler {
	if version == "" {
		version = "1.0.0"
	}
	return &Handler{version: version}
}

// ServeHTTP handles POST /graphql.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpx.WriteJSON(w, http.StatusMethodNotAllowed, Response{
			Errors: []map[string]any{{"message": "POST required"}},
		})
		return
	}
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteJSON(w, http.StatusBadRequest, Response{
			Errors: []map[string]any{{"message": "invalid json"}},
		})
		return
	}
	q := strings.TrimSpace(req.Query)
	switch {
	case strings.Contains(q, "__typename") || strings.Contains(q, "health"):
		httpx.WriteJSON(w, http.StatusOK, Response{Data: map[string]any{
			"health": map[string]any{"status": "ok", "version": h.version},
		}})
	case strings.Contains(q, "viewer"):
		httpx.WriteJSON(w, http.StatusOK, Response{Data: map[string]any{
			"viewer": map[string]any{"authenticated": false},
		}})
	default:
		httpx.WriteJSON(w, http.StatusOK, Response{
			Errors: []map[string]any{{"message": "unsupported query; REST /v1 APIs are primary in MVP"}},
		})
	}
}
// maintenance note 74

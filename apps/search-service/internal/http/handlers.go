package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sentinelvault/sentinelvault/apps/search-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for search-service.
type Handler struct {
	platform *platform.Service
	service  *domain.Service
}

// NewHandler creates an HTTP handler.
func NewHandler(p *platform.Service, svc *domain.Service) *Handler {
	return &Handler{platform: p, service: svc}
}

// RegisterRoutes mounts service routes on mux.
func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /v1/search/index", h.index)
	mux.HandleFunc("GET /v1/search", h.search)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "search-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	var req domain.IndexInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	doc, err := h.service.Upsert(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, doc)
}

func (h *Handler) search(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	indexName := r.URL.Query().Get("index")
	q := r.URL.Query().Get("q")
	tagsRaw := r.URL.Query().Get("tags")
	var tags []string
	if tagsRaw != "" {
		tags = strings.Split(tagsRaw, ",")
	}
	limit := 50
	hits := h.service.Query(r.Context(), orgID, indexName, q, tags, limit)
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": hits})
}

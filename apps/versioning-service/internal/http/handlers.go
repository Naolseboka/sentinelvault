package http

import (
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/versioning-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for versioning-service.
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
	mux.HandleFunc("GET /v1/versions", h.list)
	mux.HandleFunc("GET /v1/versions/{id}", h.get)
	mux.HandleFunc("POST /v1/versions/{id}/restore", h.restore)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "versioning-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	snap, err := h.service.Get(r.Context(), orgID, r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, snap)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	itemID := r.URL.Query().Get("item_id")
	items, err := h.service.List(r.Context(), orgID, itemID)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": items})
}

func (h *Handler) restore(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	snap, err := h.service.Restore(r.Context(), orgID, r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, snap)
}

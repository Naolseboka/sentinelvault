package http

import (
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/backup-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for backup-service.
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
	mux.HandleFunc("POST /v1/backups", h.create)
	mux.HandleFunc("POST /v1/backups/{id}/restore", h.restore)
	mux.HandleFunc("GET /v1/backups", h.list)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "backup-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req domain.BackupInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	snap, err := h.service.Backup(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, snap)
}

func (h *Handler) restore(w http.ResponseWriter, r *http.Request) {
	var dest map[string]any
	snap, err := h.service.Restore(r.Context(), r.PathValue("id"), &dest)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"snapshot": snap, "payload": dest})
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	items := h.service.List(r.Context(), orgID)
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": items})
}

package http

import (
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/directory-sync-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for directory-sync-service.
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
	mux.HandleFunc("POST /v1/directory-sync/config", h.configure)
	mux.HandleFunc("POST /v1/directory-sync/jobs", h.run)
	mux.HandleFunc("GET /v1/directory-sync/jobs", h.list)
	mux.HandleFunc("GET /v1/directory-sync/jobs/{id}", h.get)
	mux.HandleFunc("GET /v1/directory-sync/jobs/{id}/conflicts", h.conflicts)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "directory-sync-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) configure(w http.ResponseWriter, r *http.Request) {
	var req domain.ConfigureInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	cfg, err := h.service.Configure(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, cfg)
}

func (h *Handler) run(w http.ResponseWriter, r *http.Request) {
	var req domain.RunSyncInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	job, err := h.service.RunSync(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, job)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	job, err := h.service.Get(r.Context(), r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, job)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	jobs := h.service.List(r.Context(), r.URL.Query().Get("org_id"))
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": jobs})
}

func (h *Handler) conflicts(w http.ResponseWriter, r *http.Request) {
	report, err := h.service.GetConflictReport(r.Context(), r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, report)
}

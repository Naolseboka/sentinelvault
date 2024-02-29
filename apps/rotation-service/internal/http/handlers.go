package http

import (
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/rotation-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for rotation-service.
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
	mux.HandleFunc("GET /v1/rotation/policies", h.listPolicies)
	mux.HandleFunc("POST /v1/rotation/policies", h.upsertPolicy)
	mux.HandleFunc("GET /v1/rotation/policies/{id}", h.getPolicy)
	mux.HandleFunc("DELETE /v1/rotation/policies/{id}", h.deletePolicy)
	mux.HandleFunc("POST /v1/rotation/run", h.runPolicy)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "rotation-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) upsertPolicy(w http.ResponseWriter, r *http.Request) {
	var req domain.PolicyInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	pol, err := h.service.UpsertPolicy(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	status := http.StatusCreated
	if req.ID != "" {
		status = http.StatusOK
	}
	httpx.WriteJSON(w, status, pol)
}

func (h *Handler) getPolicy(w http.ResponseWriter, r *http.Request) {
	pol, err := h.service.GetPolicy(r.Context(), r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, pol)
}

func (h *Handler) listPolicies(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	policies := h.service.ListPolicies(r.Context(), orgID)
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": policies})
}

func (h *Handler) deletePolicy(w http.ResponseWriter, r *http.Request) {
	if err := h.service.DeletePolicy(r.Context(), r.PathValue("id")); err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) runPolicy(w http.ResponseWriter, r *http.Request) {
	var req domain.RunInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	job, err := h.service.RunPolicy(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, job)
}

package http

import (
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/admin-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for admin-service.
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
	mux.HandleFunc("GET /v1/admin/feature-flags", h.listFlags)
	mux.HandleFunc("PUT /v1/admin/feature-flags/{name}", h.setFlag)
	mux.HandleFunc("POST /v1/admin/orgs/{org_id}/suspend", h.suspendOrg)
	mux.HandleFunc("GET /v1/admin/orgs/{org_id}", h.getOrgState)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "admin-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) listFlags(w http.ResponseWriter, r *http.Request) {
	flags := h.service.ListFeatureFlags(r.Context())
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": flags})
}

func (h *Handler) setFlag(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Enabled bool `json:"enabled"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	flag, err := h.service.SetFeatureFlag(r.Context(), r.PathValue("name"), req.Enabled)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, flag)
}

func (h *Handler) suspendOrg(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Reason string `json:"reason"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)
	st, err := h.service.SuspendOrg(r.Context(), r.PathValue("org_id"), req.Reason)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, st)
}

func (h *Handler) getOrgState(w http.ResponseWriter, r *http.Request) {
	st, err := h.service.GetOrgState(r.Context(), r.PathValue("org_id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, st)
}

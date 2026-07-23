package http

import (
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/billing-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for billing-service.
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
	mux.HandleFunc("POST /v1/billing/licenses", h.issue)
	mux.HandleFunc("GET /v1/billing/licenses/{org_id}", h.getLicense)
	mux.HandleFunc("POST /v1/billing/features/check", h.checkFeature)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "billing-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) issue(w http.ResponseWriter, r *http.Request) {
	var req domain.IssueInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	lic, err := h.service.Issue(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, lic)
}

func (h *Handler) getLicense(w http.ResponseWriter, r *http.Request) {
	lic, err := h.service.GetLicense(r.Context(), r.PathValue("org_id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, lic)
}

func (h *Handler) checkFeature(w http.ResponseWriter, r *http.Request) {
	var req domain.FeatureCheckInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	if err := h.service.CheckFeature(r.Context(), req); err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]string{"status": "allowed"})
}

// maintenance note 79

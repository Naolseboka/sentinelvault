package http

import (
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/device-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

type Handler struct {
	platform *platform.Service
	service  *domain.Service
}

func NewHandler(p *platform.Service, svc *domain.Service) *Handler {
	return &Handler{platform: p, service: svc}
}

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /v1/devices", h.register)
	mux.HandleFunc("GET /v1/devices/{id}", h.get)
	mux.HandleFunc("GET /v1/devices", h.list)
	mux.HandleFunc("POST /v1/devices/{id}/trust", h.setTrust)
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID      string `json:"user_id"`
		OrgID       string `json:"org_id"`
		Fingerprint string `json:"fingerprint"`
		Name        string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	d, err := h.service.Register(r.Context(), req.UserID, req.OrgID, req.Fingerprint, req.Name)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, d)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	d, err := h.service.Get(r.Context(), r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, d)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	devices, err := h.service.ListByUser(r.Context(), r.URL.Query().Get("user_id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": devices})
}

func (h *Handler) setTrust(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TrustLevel string `json:"trust_level"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	d, err := h.service.SetTrustLevel(r.Context(), r.PathValue("id"), req.TrustLevel)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, d)
}

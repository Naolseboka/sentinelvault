package http

import (
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/session-service/internal/domain"
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
	mux.HandleFunc("POST /v1/sessions", h.create)
	mux.HandleFunc("GET /v1/sessions/{id}", h.get)
	mux.HandleFunc("GET /v1/sessions", h.listByUser)
	mux.HandleFunc("POST /v1/sessions/{id}/revoke", h.revoke)
	mux.HandleFunc("POST /v1/sessions/rotate", h.rotate)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID   string `json:"user_id"`
		OrgID    string `json:"org_id"`
		DeviceID string `json:"device_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	sess, refresh, err := h.service.Create(r.Context(), req.UserID, req.OrgID, req.DeviceID)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, map[string]any{"session": sess, "refresh_token": refresh})
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	sess, err := h.service.Get(r.Context(), r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, sess)
}

func (h *Handler) listByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	sessions, err := h.service.ListByUser(r.Context(), userID)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": sessions})
}

func (h *Handler) revoke(w http.ResponseWriter, r *http.Request) {
	sess, err := h.service.Revoke(r.Context(), r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, sess)
}

func (h *Handler) rotate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	sess, refresh, err := h.service.RotateRefresh(r.Context(), req.RefreshToken)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"session": sess, "refresh_token": refresh})
}

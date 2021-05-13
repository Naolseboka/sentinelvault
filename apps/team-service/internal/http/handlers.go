package http

import (
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/team-service/internal/domain"
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
	mux.HandleFunc("POST /v1/teams", h.createTeam)
	mux.HandleFunc("GET /v1/teams/{id}", h.getTeam)
	mux.HandleFunc("GET /v1/teams", h.listTeams)
	mux.HandleFunc("POST /v1/teams/{id}/members", h.addMember)
	mux.HandleFunc("GET /v1/teams/{id}/members", h.listMembers)
}

func (h *Handler) createTeam(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrgID string `json:"org_id"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	team, err := h.service.CreateTeam(r.Context(), req.OrgID, req.Name)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, team)
}

func (h *Handler) getTeam(w http.ResponseWriter, r *http.Request) {
	team, err := h.service.GetTeam(r.Context(), r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, team)
}

func (h *Handler) listTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.service.ListTeams(r.Context(), r.URL.Query().Get("org_id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": teams})
}

func (h *Handler) addMember(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID string `json:"user_id"`
		Role   string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	m, err := h.service.AddMember(r.Context(), r.PathValue("id"), req.UserID, req.Role)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, m)
}

func (h *Handler) listMembers(w http.ResponseWriter, r *http.Request) {
	members, err := h.service.ListMembers(r.Context(), r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": members})
}

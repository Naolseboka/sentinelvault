package http

import (
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/policy-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
	"github.com/sentinelvault/sentinelvault/internal/policy"
)

type Handler struct {
	platform *platform.Service
	service  *domain.Service
}

func NewHandler(p *platform.Service, svc *domain.Service) *Handler {
	return &Handler{platform: p, service: svc}
}

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /v1/policies", h.create)
	mux.HandleFunc("GET /v1/policies/{id}", h.get)
	mux.HandleFunc("POST /v1/policies/evaluate", h.evaluate)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrgID string        `json:"org_id"`
		Name  string        `json:"name"`
		Rules []policy.Rule `json:"rules"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	doc, err := h.service.Create(r.Context(), req.OrgID, req.Name, req.Rules)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, doc)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	doc, err := h.service.Get(r.Context(), r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, doc)
}

func (h *Handler) evaluate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrgID string                  `json:"org_id"`
		Req   policy.EvaluationRequest `json:"request"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	dec := h.service.Evaluate(r.Context(), req.OrgID, req.Req)
	httpx.WriteJSON(w, http.StatusOK, dec)
}

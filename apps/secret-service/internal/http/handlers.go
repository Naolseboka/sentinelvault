package http

import (
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/secret-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for secret-service.
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
	mux.HandleFunc("GET /v1/secrets", h.list)
	mux.HandleFunc("POST /v1/secrets", h.create)
	mux.HandleFunc("GET /v1/secrets/{id}", h.get)
	mux.HandleFunc("PUT /v1/secrets/{id}", h.update)
	mux.HandleFunc("DELETE /v1/secrets/{id}", h.delete)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "secret-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	item, err := h.service.Create(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, item)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	item, err := h.service.Get(r.Context(), orgID, r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, item)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	params := common.ParsePageParams(r)
	orgID := r.URL.Query().Get("org_id")
	items, total, err := h.service.List(r.Context(), orgID, params.Offset, params.PageSize)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{
		"data":       items,
		"pagination": common.NewPageMeta(params, total),
	})
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	var req domain.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	item, err := h.service.Update(r.Context(), r.PathValue("id"), req)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, item)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	if err := h.service.Delete(r.Context(), orgID, r.PathValue("id")); err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

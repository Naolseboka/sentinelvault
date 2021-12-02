package http

import (
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/audit-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for audit-service.
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
	mux.HandleFunc("POST /v1/audit/events", h.appendEvent)
	mux.HandleFunc("GET /v1/audit/events", h.listEvents)
	mux.HandleFunc("POST /v1/audit/verify", h.verify)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "audit-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) appendEvent(w http.ResponseWriter, r *http.Request) {
	var req domain.AppendInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	ev, err := h.service.Append(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, ev)
}

func (h *Handler) listEvents(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) verify(w http.ResponseWriter, r *http.Request) {
	if err := h.service.VerifyIntegrity(r.Context()); err != nil {
		httpx.WriteError(w, http.StatusConflict, common.New(common.ErrCodeConflict, err.Error()))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]string{"status": "valid"})
}

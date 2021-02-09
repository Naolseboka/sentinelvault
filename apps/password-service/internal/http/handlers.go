package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sentinelvault/sentinelvault/apps/password-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for password-service.
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
	mux.HandleFunc("GET /v1/passwords", h.list)
	mux.HandleFunc("POST /v1/passwords", h.create)
	mux.HandleFunc("GET /v1/passwords/{id}", h.get)
	mux.HandleFunc("PUT /v1/passwords/{id}", h.update)
	mux.HandleFunc("DELETE /v1/passwords/{id}", h.delete)
	mux.HandleFunc("POST /v1/passwords/{id}/archive", h.archive)
	mux.HandleFunc("POST /v1/passwords/{id}/restore", h.restore)
	mux.HandleFunc("POST /v1/passwords/collections", h.createCollection)
	mux.HandleFunc("POST /v1/passwords/folders", h.createFolder)
	mux.HandleFunc("POST /v1/passwords/{id}/shares", h.createShare)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "password-service",
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

func (h *Handler) archive(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrgID string `json:"org_id"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)
	orgID := req.OrgID
	if orgID == "" {
		orgID = r.URL.Query().Get("org_id")
	}
	if err := h.service.Archive(r.Context(), orgID, r.PathValue("id")); err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) restore(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	item, err := h.service.Restore(r.Context(), orgID, r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, item)
}

func (h *Handler) createCollection(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrgID       string `json:"org_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	col, err := h.service.CreateCollection(r.Context(), req.OrgID, req.Name, req.Description)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, col)
}

func (h *Handler) createFolder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrgID        string `json:"org_id"`
		CollectionID string `json:"collection_id"`
		ParentID     string `json:"parent_id"`
		Name         string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	f, err := h.service.CreateFolder(r.Context(), req.OrgID, req.CollectionID, req.Name, req.ParentID)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, f)
}

func (h *Handler) createShare(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrgID      string    `json:"org_id"`
		GranteeID  string    `json:"grantee_id"`
		Permission string    `json:"permission"`
		ExpiresAt  time.Time `json:"expires_at"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	g, err := h.service.CreateShareGrant(r.Context(), req.OrgID, r.PathValue("id"), req.GranteeID, req.Permission, req.ExpiresAt)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, g)
}

package http

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/attachment-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for attachment-service.
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
	mux.HandleFunc("POST /v1/attachments", h.upload)
	mux.HandleFunc("GET /v1/attachments/{id}", h.getMetadata)
	mux.HandleFunc("GET /v1/attachments/{id}/download", h.download)
	mux.HandleFunc("GET /v1/attachments", h.listByItem)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "attachment-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) upload(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrgID       string `json:"org_id"`
		VaultItemID string `json:"vault_item_id"`
		KeyID       string `json:"key_id"`
		Name        string `json:"name"`
		ContentType string `json:"content_type"`
		Data        string `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	data, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid data encoding"))
		return
	}
	meta, err := h.service.Upload(r.Context(), domain.UploadRequest{
		OrgID:       req.OrgID,
		VaultItemID: req.VaultItemID,
		KeyID:       req.KeyID,
		Name:        req.Name,
		ContentType: req.ContentType,
		Data:        data,
	})
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, meta)
}

func (h *Handler) getMetadata(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	meta, err := h.service.GetMetadata(r.Context(), orgID, r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, meta)
}

func (h *Handler) download(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	data, meta, err := h.service.Download(r.Context(), orgID, r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	w.Header().Set("Content-Type", meta.ContentType)
	w.Header().Set("Content-Disposition", "attachment; filename="+meta.Name)
	_, _ = io.WriteString(w, base64.StdEncoding.EncodeToString(data))
}

func (h *Handler) listByItem(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	vaultItemID := r.URL.Query().Get("vault_item_id")
	items, err := h.service.ListByItem(r.Context(), orgID, vaultItemID)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": items})
}

package http

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/kms-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for kms-service.
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
	mux.HandleFunc("GET /v1/kms/keys", h.listKeys)
	mux.HandleFunc("POST /v1/kms/keys", h.createKey)
	mux.HandleFunc("GET /v1/kms/keys/{id}", h.getKey)
	mux.HandleFunc("POST /v1/kms/keys/{id}/rotate", h.rotateKey)
	mux.HandleFunc("POST /v1/kms/wrap", h.wrap)
	mux.HandleFunc("POST /v1/kms/unwrap", h.unwrap)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "kms-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) createKey(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrgID string `json:"org_id"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	key, err := h.service.CreateKey(r.Context(), req.OrgID, req.Name)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, key)
}

func (h *Handler) getKey(w http.ResponseWriter, r *http.Request) {
	key, err := h.service.GetKey(r.Context(), r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, key)
}

func (h *Handler) listKeys(w http.ResponseWriter, r *http.Request) {
	orgID := r.URL.Query().Get("org_id")
	keys, err := h.service.ListKeys(r.Context(), orgID)
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"data": keys})
}

func (h *Handler) rotateKey(w http.ResponseWriter, r *http.Request) {
	key, err := h.service.RotateKey(r.Context(), r.PathValue("id"))
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, key)
}

func (h *Handler) wrap(w http.ResponseWriter, r *http.Request) {
	var req struct {
		KeyID string `json:"key_id"`
		DEK   string `json:"dek"`
		AAD   string `json:"aad"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	dek, err := base64.StdEncoding.DecodeString(req.DEK)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid dek encoding"))
		return
	}
	resp, err := h.service.WrapDEK(r.Context(), domain.WrapRequest{
		KeyID: req.KeyID,
		DEK:   dek,
		AAD:   []byte(req.AAD),
	})
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{
		"wrapped_dek": base64.StdEncoding.EncodeToString(resp.WrappedDEK),
		"wrap_nonce":  base64.StdEncoding.EncodeToString(resp.WrapNonce),
		"key_version": resp.KeyVersion,
	})
}

func (h *Handler) unwrap(w http.ResponseWriter, r *http.Request) {
	var req struct {
		KeyID      string `json:"key_id"`
		WrappedDEK string `json:"wrapped_dek"`
		WrapNonce  string `json:"wrap_nonce"`
		AAD        string `json:"aad"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	wrapped, err := base64.StdEncoding.DecodeString(req.WrappedDEK)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid wrapped_dek encoding"))
		return
	}
	nonce, err := base64.StdEncoding.DecodeString(req.WrapNonce)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid wrap_nonce encoding"))
		return
	}
	resp, err := h.service.UnwrapDEK(r.Context(), domain.UnwrapRequest{
		KeyID:      req.KeyID,
		WrappedDEK: wrapped,
		WrapNonce:  nonce,
		AAD:        []byte(req.AAD),
	})
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{
		"dek": base64.StdEncoding.EncodeToString(resp.DEK),
	})
}

package http

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/encryption-service/internal/domain"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for encryption-service.
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
	mux.HandleFunc("POST /v1/encryption/encrypt", h.encrypt)
	mux.HandleFunc("POST /v1/encryption/decrypt", h.decrypt)
	mux.HandleFunc("GET /v1/meta", h.meta)
}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"service": "encryption-service",
		"schema":  domain.SchemaName(),
		"table":   domain.TableName(),
	})
}

func (h *Handler) encrypt(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Plaintext string `json:"plaintext"`
		KeyID     string `json:"key_id"`
		AAD       string `json:"aad"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	resp, err := h.service.EnvelopeEncrypt(r.Context(), domain.EncryptRequest{
		Plaintext: []byte(req.Plaintext),
		KeyID:     req.KeyID,
		AAD:       []byte(req.AAD),
	})
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{
		"ciphertext": base64.StdEncoding.EncodeToString(resp.Ciphertext),
		"key_id":     resp.KeyID,
	})
}

func (h *Handler) decrypt(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Ciphertext string `json:"ciphertext"`
		AAD        string `json:"aad"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
		return
	}
	ciphertext, err := base64.StdEncoding.DecodeString(req.Ciphertext)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid ciphertext encoding"))
		return
	}
	resp, err := h.service.EnvelopeDecrypt(r.Context(), domain.DecryptRequest{
		Ciphertext: ciphertext,
		AAD:        []byte(req.AAD),
	})
	if err != nil {
		httpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{
		"plaintext": string(resp.Plaintext),
	})
}

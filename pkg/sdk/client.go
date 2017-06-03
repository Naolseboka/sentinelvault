package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is an HTTP client for SentinelVault APIs.
type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

// Option configures the client.
type Option func(*Client)

// WithToken sets the bearer token.
func WithToken(token string) Option {
	return func(c *Client) { c.token = token }
}

// WithHTTPClient sets a custom http client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.httpClient = hc }
}

// New creates a SentinelVault API client.
func New(baseURL string, opts ...Option) *Client {
	c := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Do executes an HTTP request with auth headers.
func (c *Client) Do(ctx context.Context, method, path string, body any, dest any) error {
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(data)
	}
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		raw, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("api error %d: %s", resp.StatusCode, string(raw))
	}
	if dest == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(dest)
}

// Get performs a GET request.
func (c *Client) Get(ctx context.Context, path string, dest any) error {
	return c.Do(ctx, http.MethodGet, path, nil, dest)
}

// Post performs a POST request.
func (c *Client) Post(ctx context.Context, path string, body, dest any) error {
	return c.Do(ctx, http.MethodPost, path, body, dest)
}

// AuthClient provides helpers for auth-service endpoints.
type AuthClient struct {
	*Client
}

// Auth returns an auth API client wrapping the base HTTP client.
func (c *Client) Auth() *AuthClient {
	return &AuthClient{Client: c}
}

// RegisterRequest bootstraps a user and organization.
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	OrgName  string `json:"org_name"`
}

// RegisterResponse holds created user and org records.
type RegisterResponse struct {
	User map[string]any `json:"user"`
	Org  map[string]any `json:"org"`
}

// Register creates a user, organization, and credential via auth-service.
func (a *AuthClient) Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error) {
	var resp RegisterResponse
	if err := a.Post(ctx, "/v1/auth/register", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LoginRequest authenticates a user.
type LoginRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	DeviceFingerprint string `json:"device_fingerprint,omitempty"`
	TOTPCode          string `json:"totp_code,omitempty"`
}

// TokenPair holds access and refresh tokens from login.
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// LoginResponse holds tokens or MFA challenge state.
type LoginResponse struct {
	Tokens      *TokenPair `json:"-"`
	MFARequired bool       `json:"mfa_required"`
	UserID      string     `json:"user_id,omitempty"`
	OrgID       string     `json:"org_id,omitempty"`
}

// Login authenticates credentials and returns tokens or MFA requirement.
func (a *AuthClient) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	var raw map[string]any
	if err := a.Post(ctx, "/v1/auth/login", req, &raw); err != nil {
		return nil, err
	}
	if mfa, ok := raw["mfa_required"].(bool); ok && mfa {
		resp := &LoginResponse{MFARequired: true}
		if uid, ok := raw["user_id"].(string); ok {
			resp.UserID = uid
		}
		if oid, ok := raw["org_id"].(string); ok {
			resp.OrgID = oid
		}
		return resp, nil
	}
	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var tokens TokenPair
	if err := json.Unmarshal(data, &tokens); err != nil {
		return nil, err
	}
	return &LoginResponse{Tokens: &tokens}, nil
}

package oidc

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// HTTPDoer is mockable HTTP client surface.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// TokenResponse is OIDC token endpoint response subset.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	Scope        string `json:"scope"`
}

// Client exchanges codes for tokens.
type Client struct {
	HTTP HTTPDoer
}

// ExchangeAuthorizationCode posts token request and parses response.
func (c *Client) ExchangeAuthorizationCode(ctx context.Context, req TokenRequest) (*TokenResponse, error) {
	if c.HTTP == nil {
		return nil, common.New(common.ErrCodeInternal, "http client required")
	}
	vals, err := req.FormValues()
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, req.TokenEndpoint, strings.NewReader(vals.Encode()))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := c.HTTP.Do(httpReq)
	if err != nil {
		return nil, common.Wrap(common.ErrCodeUnavailable, "token endpoint", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		return nil, common.New(common.ErrCodeUnauthorized, "token exchange failed")
	}
	var out TokenResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "token json", err)
	}
	if out.AccessToken == "" {
		return nil, common.New(common.ErrCodeUnauthorized, "missing access_token")
	}
	return &out, nil
}

// FetchDiscovery loads provider metadata from URL.
func (c *Client) FetchDiscovery(ctx context.Context, issuer string) (*DiscoveryDocument, error) {
	if c.HTTP == nil {
		return nil, common.New(common.ErrCodeInternal, "http client required")
	}
	u, err := url.Parse(strings.TrimRight(issuer, "/") + "/.well-known/openid-configuration")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return ParseDiscovery(data)
}

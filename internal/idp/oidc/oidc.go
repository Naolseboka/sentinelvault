package oidc

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/crypto"
)

// DiscoveryDocument is OIDC provider metadata subset.
type DiscoveryDocument struct {
	Issuer                string   `json:"issuer"`
	AuthorizationEndpoint string   `json:"authorization_endpoint"`
	TokenEndpoint         string   `json:"token_endpoint"`
	JWKSURI               string   `json:"jwks_uri"`
	UserinfoEndpoint      string   `json:"userinfo_endpoint"`
	ScopesSupported       []string `json:"scopes_supported"`
	ResponseTypes         []string `json:"response_types_supported"`
	CodeChallengeMethods  []string `json:"code_challenge_methods_supported"`
}

// ParseDiscovery parses discovery JSON.
func ParseDiscovery(data []byte) (*DiscoveryDocument, error) {
	var d DiscoveryDocument
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid discovery json", err)
	}
	if d.Issuer == "" || d.AuthorizationEndpoint == "" || d.TokenEndpoint == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "discovery missing required endpoints")
	}
	return &d, nil
}

// IDTokenClaims are validated claims from an ID token payload (JWT middle segment decoded).
type IDTokenClaims struct {
	Iss   string   `json:"iss"`
	Sub   string   `json:"sub"`
	Aud   any      `json:"aud"`
	Exp   int64    `json:"exp"`
	Iat   int64    `json:"iat"`
	Nonce string   `json:"nonce"`
	Email string   `json:"email"`
	Name  string   `json:"name"`
	Amr   []string `json:"amr"`
}

// ValidateClaims checks iss/aud/exp/nonce with clock skew.
func ValidateClaims(c IDTokenClaims, issuer string, audiences []string, nonce string, now time.Time, skew time.Duration) error {
	if c.Iss != issuer {
		return common.New(common.ErrCodeUnauthorized, "issuer mismatch")
	}
	if c.Sub == "" {
		return common.New(common.ErrCodeUnauthorized, "subject required")
	}
	if !audienceMatch(c.Aud, audiences) {
		return common.New(common.ErrCodeUnauthorized, "audience mismatch")
	}
	if skew <= 0 {
		skew = 2 * time.Minute
	}
	exp := time.Unix(c.Exp, 0)
	if now.After(exp.Add(skew)) {
		return common.New(common.ErrCodeUnauthorized, "token expired")
	}
	if c.Iat > 0 {
		iat := time.Unix(c.Iat, 0)
		if iat.After(now.Add(skew)) {
			return common.New(common.ErrCodeUnauthorized, "token used before issued")
		}
	}
	if nonce != "" && c.Nonce != nonce {
		return common.New(common.ErrCodeUnauthorized, "nonce mismatch")
	}
	return nil
}

func audienceMatch(aud any, want []string) bool {
	set := map[string]struct{}{}
	for _, w := range want {
		set[w] = struct{}{}
	}
	switch v := aud.(type) {
	case string:
		_, ok := set[v]
		return ok
	case []any:
		for _, x := range v {
			s, _ := x.(string)
			if _, ok := set[s]; ok {
				return true
			}
		}
	}
	return false
}

// PKCE holds verifier and challenge.
type PKCE struct {
	Verifier  string
	Challenge string
	Method    string
}

// NewPKCES256 creates S256 PKCE pair.
func NewPKCES256() (*PKCE, error) {
	raw, err := crypto.RandomBytes(32)
	if err != nil {
		return nil, err
	}
	verifier := base64.RawURLEncoding.EncodeToString(raw)
	sum := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(sum[:])
	return &PKCE{Verifier: verifier, Challenge: challenge, Method: "S256"}, nil
}

// VerifyPKCE checks challenge matches verifier.
func VerifyPKCE(verifier, challenge, method string) bool {
	if method != "S256" {
		return false
	}
	sum := sha256.Sum256([]byte(verifier))
	want := base64.RawURLEncoding.EncodeToString(sum[:])
	return want == challenge
}

// AuthCodeURLInput builds authorize redirect.
type AuthCodeURLInput struct {
	AuthorizationEndpoint string
	ClientID              string
	RedirectURI           string
	Scope                 []string
	State                 string
	Nonce                 string
	PKCEChallenge         string
	PKCEMethod            string
	Extra                 map[string]string
}

// BuildAuthCodeURL constructs the authorization URL.
func BuildAuthCodeURL(in AuthCodeURLInput) (string, error) {
	if in.AuthorizationEndpoint == "" || in.ClientID == "" || in.RedirectURI == "" || in.State == "" {
		return "", common.New(common.ErrCodeInvalidArgument, "authorization_endpoint, client_id, redirect_uri, state required")
	}
	u, err := url.Parse(in.AuthorizationEndpoint)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("response_type", "code")
	q.Set("client_id", in.ClientID)
	q.Set("redirect_uri", in.RedirectURI)
	q.Set("scope", strings.Join(in.Scope, " "))
	q.Set("state", in.State)
	if in.Nonce != "" {
		q.Set("nonce", in.Nonce)
	}
	if in.PKCEChallenge != "" {
		q.Set("code_challenge", in.PKCEChallenge)
		method := in.PKCEMethod
		if method == "" {
			method = "S256"
		}
		q.Set("code_challenge_method", method)
	}
	for k, v := range in.Extra {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// TokenRequest is a token endpoint form body builder.
type TokenRequest struct {
	TokenEndpoint string
	GrantType     string
	Code          string
	RedirectURI   string
	ClientID      string
	ClientSecret  string
	CodeVerifier  string
	RefreshToken  string
}

// FormValues returns application/x-www-form-urlencoded values.
func (t TokenRequest) FormValues() (url.Values, error) {
	if t.TokenEndpoint == "" || t.ClientID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "token_endpoint and client_id required")
	}
	v := url.Values{}
	v.Set("client_id", t.ClientID)
	if t.ClientSecret != "" {
		v.Set("client_secret", t.ClientSecret)
	}
	switch t.GrantType {
	case "authorization_code":
		if t.Code == "" || t.RedirectURI == "" {
			return nil, common.New(common.ErrCodeInvalidArgument, "code and redirect_uri required")
		}
		v.Set("grant_type", "authorization_code")
		v.Set("code", t.Code)
		v.Set("redirect_uri", t.RedirectURI)
		if t.CodeVerifier != "" {
			v.Set("code_verifier", t.CodeVerifier)
		}
	case "refresh_token":
		if t.RefreshToken == "" {
			return nil, common.New(common.ErrCodeInvalidArgument, "refresh_token required")
		}
		v.Set("grant_type", "refresh_token")
		v.Set("refresh_token", t.RefreshToken)
	default:
		return nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("unsupported grant %q", t.GrantType))
	}
	return v, nil
}

// DecodeClaimsPayload decodes the JWT payload segment without verifying signature.
func DecodeClaimsPayload(jwt string) (*IDTokenClaims, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) < 2 {
		return nil, common.New(common.ErrCodeInvalidArgument, "invalid jwt")
	}
	raw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "jwt payload decode", err)
	}
	var c IDTokenClaims
	if err := json.Unmarshal(raw, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

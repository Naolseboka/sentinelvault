package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/crypto"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

const defaultAccessTTL = 15 * time.Minute

// TokenManager issues and validates HMAC-SHA256 JWT access tokens.
type TokenManager struct {
	secret    []byte
	accessTTL time.Duration
}

// NewTokenManager creates a token manager with the given signing secret.
func NewTokenManager(secret []byte, accessTTL time.Duration) *TokenManager {
	if accessTTL <= 0 {
		accessTTL = defaultAccessTTL
	}
	return &TokenManager{secret: secret, accessTTL: accessTTL}
}

type accessClaims struct {
	jwt.RegisteredClaims
	OrgID       string   `json:"org_id"`
	SessionID   string   `json:"session_id"`
	Roles       []string `json:"roles"`
	MFAVerified bool     `json:"mfa_verified"`
}

// IssueAccessToken signs a new access token for the given claims.
func (m *TokenManager) IssueAccessToken(c Claims) (string, error) {
	now := time.Now().UTC()
	if c.JTI == "" {
		c.JTI = common.NewID()
	}
	if c.IssuedAt.IsZero() {
		c.IssuedAt = now
	}
	exp := c.ExpiresAt
	if exp.IsZero() {
		exp = now.Add(m.accessTTL)
	}

	claims := accessClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   string(c.Subject),
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(c.IssuedAt),
			ID:        c.JTI,
		},
		OrgID:       c.OrgID,
		SessionID:   c.SessionID,
		Roles:       c.Roles,
		MFAVerified: c.MFAVerified,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(m.secret)
	if err != nil {
		return "", fmt.Errorf("sign access token: %w", err)
	}
	return signed, nil
}

// ParseAccessToken validates and parses an access token.
func (m *TokenManager) ParseAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &accessClaims{}, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("%w: unexpected signing method", ErrInvalidToken)
		}
		return m.secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	claims, ok := token.Claims.(*accessClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	out := &Claims{
		Subject:     UserID(claims.Subject),
		OrgID:       claims.OrgID,
		SessionID:   claims.SessionID,
		Roles:       claims.Roles,
		MFAVerified: claims.MFAVerified,
		JTI:         claims.ID,
	}
	if claims.IssuedAt != nil {
		out.IssuedAt = claims.IssuedAt.Time
	}
	if claims.ExpiresAt != nil {
		out.ExpiresAt = claims.ExpiresAt.Time
	}
	return out, nil
}

// NewRefreshTokenID generates an opaque refresh token identifier.
func NewRefreshTokenID() (string, error) {
	b, err := crypto.RandomBytes(32)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("rt_%x", b), nil
}

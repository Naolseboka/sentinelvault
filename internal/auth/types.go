package auth

import "time"

// UserID identifies an authenticated user.
type UserID string

// Claims represents validated access token claims.
type Claims struct {
	Subject     UserID    `json:"sub"`
	OrgID       string    `json:"org_id"`
	SessionID   string    `json:"session_id"`
	Roles       []string  `json:"roles"`
	MFAVerified bool      `json:"mfa_verified"`
	JTI         string    `json:"jti"`
	IssuedAt    time.Time `json:"iat"`
	ExpiresAt   time.Time `json:"exp"`
}

// TokenPair holds access and refresh tokens.
type TokenPair struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
	MFARequired  bool      `json:"mfa_required,omitempty"`
}

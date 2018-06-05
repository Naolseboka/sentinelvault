package auth_test

import (
	"testing"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/auth"
)

func TestIssueAndParseAccessToken(t *testing.T) {
	t.Parallel()
	mgr := auth.NewTokenManager([]byte("test-secret-key-32bytes-long!!"), time.Minute)
	issued, err := mgr.IssueAccessToken(auth.Claims{
		Subject:     "user-1",
		OrgID:       "org-1",
		SessionID:   "sess-1",
		Roles:       []string{"admin", "member"},
		MFAVerified: true,
	})
	if err != nil {
		t.Fatalf("issue: %v", err)
	}

	parsed, err := mgr.ParseAccessToken(issued)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if parsed.Subject != "user-1" {
		t.Fatalf("subject = %q", parsed.Subject)
	}
	if parsed.OrgID != "org-1" {
		t.Fatalf("org_id = %q", parsed.OrgID)
	}
	if parsed.SessionID != "sess-1" {
		t.Fatalf("session_id = %q", parsed.SessionID)
	}
	if !parsed.MFAVerified {
		t.Fatal("expected mfa_verified")
	}
	if len(parsed.Roles) != 2 {
		t.Fatalf("roles = %v", parsed.Roles)
	}
	if parsed.JTI == "" {
		t.Fatal("expected jti")
	}
}

func TestParseAccessTokenInvalidSecret(t *testing.T) {
	t.Parallel()
	mgr := auth.NewTokenManager([]byte("secret-a"), time.Minute)
	token, err := mgr.IssueAccessToken(auth.Claims{Subject: "u", OrgID: "o", SessionID: "s"})
	if err != nil {
		t.Fatal(err)
	}
	other := auth.NewTokenManager([]byte("secret-b"), time.Minute)
	if _, err := other.ParseAccessToken(token); err == nil {
		t.Fatal("expected invalid token error")
	}
}

func TestParseAccessTokenExpired(t *testing.T) {
	t.Parallel()
	mgr := auth.NewTokenManager([]byte("test-secret-key-32bytes-long!!"), time.Minute)
	token, err := mgr.IssueAccessToken(auth.Claims{
		Subject:   "user-1",
		OrgID:     "org-1",
		SessionID: "sess-1",
		ExpiresAt: time.Now().UTC().Add(-time.Minute),
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := mgr.ParseAccessToken(token); err == nil {
		t.Fatal("expected expired token")
	}
}

func TestNewTokenManagerDefaultTTL(t *testing.T) {
	t.Parallel()
	mgr := auth.NewTokenManager([]byte("test-secret-key-32bytes-long!!"), 0)
	token, err := mgr.IssueAccessToken(auth.Claims{Subject: "u", OrgID: "o", SessionID: "s"})
	if err != nil {
		t.Fatal(err)
	}
	claims, err := mgr.ParseAccessToken(token)
	if err != nil {
		t.Fatal(err)
	}
	if claims.ExpiresAt.Before(time.Now().UTC()) {
		t.Fatal("expected future expiry with default ttl")
	}
}

func TestNewRefreshTokenID(t *testing.T) {
	t.Parallel()
	a, err := auth.NewRefreshTokenID()
	if err != nil {
		t.Fatal(err)
	}
	b, err := auth.NewRefreshTokenID()
	if err != nil {
		t.Fatal(err)
	}
	if a == b {
		t.Fatal("refresh token ids should be unique")
	}
	if len(a) < 10 {
		t.Fatalf("unexpected refresh token id: %q", a)
	}
}

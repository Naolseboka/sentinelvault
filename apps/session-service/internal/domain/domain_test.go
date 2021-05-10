package domain_test

import (
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/apps/session-service/internal/domain"
)

func TestSessionLifecycle(t *testing.T) {
	t.Parallel()
	svc := domain.NewService(domain.NewMemoryRepository())
	ctx := context.Background()

	sess, refresh, err := svc.Create(ctx, "user-1", "org-1", "device-1")
	if err != nil {
		t.Fatal(err)
	}
	if sess.ID == "" || refresh == "" {
		t.Fatal("expected session and refresh token")
	}

	got, err := svc.Get(ctx, sess.ID)
	if err != nil || got.UserID != "user-1" {
		t.Fatalf("get: %v %+v", err, got)
	}

	_, newRefresh, err := svc.RotateRefresh(ctx, refresh)
	if err != nil {
		t.Fatal(err)
	}
	if newRefresh == refresh {
		t.Fatal("expected rotated refresh token")
	}

	revoked, err := svc.Revoke(ctx, sess.ID)
	if err != nil || revoked.RevokedAt == nil {
		t.Fatalf("revoke: %v %+v", err, revoked)
	}
}

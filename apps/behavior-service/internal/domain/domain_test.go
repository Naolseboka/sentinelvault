package domain

import (
	"context"
	"testing"

	"github.com/sentinelvault/sentinelvault/internal/behavior"
)

func TestBehaviorObserve(t *testing.T) {
	svc := NewService(NewTracker())
	ctx := context.Background()
	for i := 0; i < 25; i++ {
		svc.Observe(ctx, ObserveInput{
			UserID: "user-1", OrgID: "org-1", Type: behavior.EventLogin, HourUTC: 9,
		})
	}
	anoms := svc.Observe(ctx, ObserveInput{
		UserID: "user-1", OrgID: "org-1", Type: behavior.EventExport, HourUTC: 3,
	})
	if len(anoms) == 0 {
		t.Fatal("expected anomalies for unusual export hour")
	}
}

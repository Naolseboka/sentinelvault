package filters

import (
	"testing"
	"time"
)

func TestBlobFilters(t *testing.T) {
	now := time.Now().UTC()
	items := []Item{
		{Fields: map[string]string{"name": "alpha", "status": "active", "org_id": "o1"}, CreatedAt: now, UpdatedAt: now},
		{Fields: map[string]string{"name": "beta", "status": "pending", "org_id": "o1"}, CreatedAt: now, UpdatedAt: now},
	}
	out := Apply(items, Filter{Name: "alp", Mode: "prefix", Limit: 10})
	if len(out) != 1 {
		t.Fatalf("%d", len(out))
	}
	out = Apply(items, Filter{Status: "pending", Mode: "exact"})
	if len(out) != 1 {
		t.Fatalf("%d", len(out))
	}
}

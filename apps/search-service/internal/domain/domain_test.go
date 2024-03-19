package domain

import (
	"context"
	"testing"
)

func TestSearchIndexAndQuery(t *testing.T) {
	svc := NewService(NewIndex())
	ctx := context.Background()

	_, err := svc.Upsert(ctx, IndexInput{
		OrgID: "org-1", Index: "vault", Title: "Production DB", Body: "postgres credentials", Tags: []string{"prod"},
	})
	if err != nil {
		t.Fatal(err)
	}

	hits := svc.Query(ctx, "org-1", "vault", "postgres", nil, 10)
	if len(hits) == 0 {
		t.Fatal("expected search hits")
	}
	if hits[0].Document.Title != "Production DB" {
		t.Fatalf("unexpected title %q", hits[0].Document.Title)
	}
}

func TestSearchRequiresOrgID(t *testing.T) {
	svc := NewService(NewIndex())
	_, err := svc.Upsert(context.Background(), IndexInput{Title: "x"})
	if err == nil {
		t.Fatal("expected org_id validation error")
	}
}

package searchx

import "testing"

func TestIndexQuery(t *testing.T) {
	idx := NewIndex()
	idx.Upsert(Document{ID: "1", OrgID: "o1", Index: "vault", Title: "AWS Root Key", Body: "production api key", Tags: []string{"aws", "prod"}})
	idx.Upsert(Document{ID: "2", OrgID: "o1", Index: "audit", Title: "Login failure", Body: "user mistyped", Tags: []string{"auth"}})
	hits := idx.Query("o1", "vault", "aws", []string{"prod"}, 10)
	if len(hits) != 1 || hits[0].Document.ID != "1" {
		t.Fatalf("hits=%+v", hits)
	}
}

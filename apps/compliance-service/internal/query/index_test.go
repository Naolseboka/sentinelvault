package query

import (
	"testing"
	"time"
)

func TestEvidenceIndexSearch(t *testing.T) {
	idx := NewEvidenceIndex()
	now := time.Now().UTC()
	idx.Upsert(EvidenceView{ID: "1", OrgID: "o1", Name: "alpha", Status: "active", CreatedAt: now, Labels: map[string]string{"env": "prod"}})
	idx.Upsert(EvidenceView{ID: "2", OrgID: "o1", Name: "beta", Status: "pending", CreatedAt: now.Add(time.Second), Labels: map[string]string{"env": "dev"}})
	idx.Upsert(EvidenceView{ID: "3", OrgID: "o2", Name: "other", Status: "active", CreatedAt: now})
	res, total, err := idx.Search(EvidenceQuery{OrgID: "o1", Status: "active", Limit: 10})
	if err != nil || total != 1 || len(res) != 1 {
		t.Fatalf("res=%v total=%d err=%v", res, total, err)
	}
	_, total, err = idx.Search(EvidenceQuery{OrgID: "o1", NameContains: "bet", Limit: 10})
	if err != nil || total != 1 {
		t.Fatalf("name filter total=%d err=%v", total, err)
	}
	_, total, err = idx.Search(EvidenceQuery{OrgID: "o1", LabelKey: "env", LabelValue: "prod", Limit: 10})
	if err != nil || total != 1 {
		t.Fatalf("label total=%d", total)
	}
	if !IsValidEvidenceStatus("active") || IsValidEvidenceStatus("nope") {
		t.Fatal("status validation")
	}
}

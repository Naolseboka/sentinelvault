package query

import (
	"testing"
	"time"
)

func TestSCIMUserIndexSearch(t *testing.T) {
	idx := NewSCIMUserIndex()
	now := time.Now().UTC()
	idx.Upsert(SCIMUserView{ID: "1", OrgID: "o1", Name: "alpha", Status: "active", CreatedAt: now, Labels: map[string]string{"env": "prod"}})
	idx.Upsert(SCIMUserView{ID: "2", OrgID: "o1", Name: "beta", Status: "pending", CreatedAt: now.Add(time.Second), Labels: map[string]string{"env": "dev"}})
	idx.Upsert(SCIMUserView{ID: "3", OrgID: "o2", Name: "other", Status: "active", CreatedAt: now})
	res, total, err := idx.Search(SCIMUserQuery{OrgID: "o1", Status: "active", Limit: 10})
	if err != nil || total != 1 || len(res) != 1 {
		t.Fatalf("res=%v total=%d err=%v", res, total, err)
	}
	res, total, err = idx.Search(SCIMUserQuery{OrgID: "o1", NameContains: "bet", Limit: 10})
	if err != nil || total != 1 {
		t.Fatalf("name filter total=%d err=%v", total, err)
	}
	res, total, err = idx.Search(SCIMUserQuery{OrgID: "o1", LabelKey: "env", LabelValue: "prod", Limit: 10})
	if err != nil || total != 1 {
		t.Fatalf("label total=%d", total)
	}
	if !IsValidSCIMUserStatus("active") || IsValidSCIMUserStatus("nope") {
		t.Fatal("status validation")
	}
}

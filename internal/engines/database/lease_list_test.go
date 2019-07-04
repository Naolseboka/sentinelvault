package database

import "testing"

func TestLeaseListModule(t *testing.T) {
	e := NewEngine()
	_, _ = e.UpsertRole(Role{OrgID: "o1", Name: "r1", DBType: "postgres", DefaultTTL: 0, MaxTTL: 0})
	if len(e.ListLeases("o1", true)) != 0 {
		t.Fatal("expected empty")
	}
}

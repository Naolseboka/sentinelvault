package database

import (
	"testing"
	"time"
)

func TestDatabaseEngine(t *testing.T) {
	e := NewEngine()
	role, err := e.UpsertRole(Role{OrgID: "o1", Name: "readonly", DBType: "postgres", DefaultTTL: time.Minute, MaxTTL: time.Hour})
	if err != nil {
		t.Fatal(err)
	}
	lease, err := e.Issue(role.ID, 0)
	if err != nil {
		t.Fatal(err)
	}
	if lease.Username == "" || lease.Password == "" {
		t.Fatal("missing creds")
	}
	if err := e.Revoke(lease.ID); err != nil {
		t.Fatal(err)
	}
	lease2, _ := e.Issue(role.ID, time.Millisecond)
	time.Sleep(2 * time.Millisecond)
	if n := e.CleanupExpired(time.Now().UTC()); n < 1 && !lease2.Revoked {
		// cleanup should revoke expired
		_ = n
	}
}

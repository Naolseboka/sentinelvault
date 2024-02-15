package jenkins

import (
	"testing"
	"time"
)

func TestJenkinsEngine(t *testing.T) {
	e := NewEngine()
	r, err := e.UpsertRole(Role{OrgID: "o1", Name: "ci", Config: map[string]string{"env": "prod"}})
	if err != nil {
		t.Fatal(err)
	}
	l, err := e.Issue(r.ID, 0, map[string]string{"job": "build"})
	if err != nil || l.Token == "" {
		t.Fatal(err)
	}
	if len(e.ListLeases("o1")) != 1 {
		t.Fatal("list")
	}
	if err := e.Revoke(l.ID); err != nil {
		t.Fatal(err)
	}
	_ = e.CleanupExpired(time.Now().UTC())
}
// maintenance note 26

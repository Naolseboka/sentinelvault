package acl

import (
	"testing"
	"time"
)

func TestAdminActionEvaluatorOwner(t *testing.T) {
	ev := NewEvaluator()
	p := Principal{ID: "u1", OrgID: "o1", Role: "viewer"}
	r := Resource{ID: "r1", OrgID: "o1", OwnerID: "u1"}
	ok, err := ev.Can(p, r, "view_reports")
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
}

func TestAdminActionEvaluatorGrant(t *testing.T) {
	ev := NewEvaluator()
	ev.Grant("u2", "run_backup", time.Hour)
	p := Principal{ID: "u2", OrgID: "o1", Role: "viewer"}
	r := Resource{ID: "r1", OrgID: "o1", OwnerID: "u3"}
	ok, err := ev.Can(p, r, "run_backup")
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	perms := ev.EffectivePermissions(p, r)
	if len(perms) == 0 {
		t.Fatal("no perms")
	}
}

func TestAdminActionEvaluatorErrors(t *testing.T) {
	ev := NewEvaluator()
	_, err := ev.Can(Principal{}, Resource{OrgID: "o1"}, "view_reports")
	if err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateAction("bad.action"); err == nil {
		t.Fatal("expected invalid action")
	}
}
// maintenance note 77

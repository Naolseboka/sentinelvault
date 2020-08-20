package acl

import (
	"testing"
	"time"
)

func TestOrganizationEvaluatorOwner(t *testing.T) {
	ev := NewEvaluator()
	p := Principal{ID: "u1", OrgID: "o1", Role: "viewer"}
	r := Resource{ID: "r1", OrgID: "o1", OwnerID: "u1"}
	ok, err := ev.Can(p, r, "read_profile")
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
}

func TestOrganizationEvaluatorGrant(t *testing.T) {
	ev := NewEvaluator()
	ev.Grant("u2", "update_profile", time.Hour)
	p := Principal{ID: "u2", OrgID: "o1", Role: "viewer"}
	r := Resource{ID: "r1", OrgID: "o1", OwnerID: "u3"}
	ok, err := ev.Can(p, r, "update_profile")
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	perms := ev.EffectivePermissions(p, r)
	if len(perms) == 0 {
		t.Fatal("no perms")
	}
}

func TestOrganizationEvaluatorErrors(t *testing.T) {
	ev := NewEvaluator()
	_, err := ev.Can(Principal{}, Resource{OrgID: "o1"}, "read_profile")
	if err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateAction("bad.action"); err == nil {
		t.Fatal("expected invalid action")
	}
}

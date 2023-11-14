package acl

import (
	"testing"
	"time"
)

func TestProvisionPlanEvaluatorOwner(t *testing.T) {
	ev := NewEvaluator()
	p := Principal{ID: "u1", OrgID: "o1", Role: "viewer"}
	r := Resource{ID: "r1", OrgID: "o1", OwnerID: "u1"}
	ok, err := ev.Can(p, r, "approve_request")
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
}

func TestProvisionPlanEvaluatorGrant(t *testing.T) {
	ev := NewEvaluator()
	ev.Grant("u2", "delegate_access", time.Hour)
	p := Principal{ID: "u2", OrgID: "o1", Role: "viewer"}
	r := Resource{ID: "r1", OrgID: "o1", OwnerID: "u3"}
	ok, err := ev.Can(p, r, "delegate_access")
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	perms := ev.EffectivePermissions(p, r)
	if len(perms) == 0 {
		t.Fatal("no perms")
	}
}

func TestProvisionPlanEvaluatorErrors(t *testing.T) {
	ev := NewEvaluator()
	_, err := ev.Can(Principal{}, Resource{OrgID: "o1"}, "approve_request")
	if err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateAction("bad.action"); err == nil {
		t.Fatal("expected invalid action")
	}
}

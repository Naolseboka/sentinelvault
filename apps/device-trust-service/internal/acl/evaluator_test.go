package acl

import (
	"testing"
	"time"
)

func TestTrustRecordEvaluatorOwner(t *testing.T) {
	ev := NewEvaluator()
	p := Principal{ID: "u1", OrgID: "o1", Role: "viewer"}
	r := Resource{ID: "r1", OrgID: "o1", OwnerID: "u1"}
	ok, err := ev.Can(p, r, "view_findings")
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
}

func TestTrustRecordEvaluatorGrant(t *testing.T) {
	ev := NewEvaluator()
	ev.Grant("u2", "ack_finding", time.Hour)
	p := Principal{ID: "u2", OrgID: "o1", Role: "viewer"}
	r := Resource{ID: "r1", OrgID: "o1", OwnerID: "u3"}
	ok, err := ev.Can(p, r, "ack_finding")
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	perms := ev.EffectivePermissions(p, r)
	if len(perms) == 0 {
		t.Fatal("no perms")
	}
}

func TestTrustRecordEvaluatorErrors(t *testing.T) {
	ev := NewEvaluator()
	_, err := ev.Can(Principal{}, Resource{OrgID: "o1"}, "view_findings")
	if err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateAction("bad.action"); err == nil {
		t.Fatal("expected invalid action")
	}
}

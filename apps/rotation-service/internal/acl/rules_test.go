package acl

import (
	"context"
	"testing"
)

func TestRuleEngineChain(t *testing.T) {
	re := NewRuleEngine(NewEvaluator())
	req := RuleRequest{OrgID: "o1", SubjectID: "u1", SubjectRole: "admin", ResourceID: "r1", ResourceOwner: "u2", Action: "read"}
	res, err := re.EvaluateChain(context.Background(), req, DefaultRules())
	if err != nil {
		t.Fatal(err)
	}
	if !res.Allowed {
		t.Fatalf("expected allow: %+v", res)
	}
}

func TestRuleEngineUnknown(t *testing.T) {
	re := NewRuleEngine(nil)
	_, err := re.Evaluate(context.Background(), "missing", RuleRequest{OrgID: "o1", SubjectID: "u1"})
	if err == nil {
		t.Fatal("expected error")
	}
}

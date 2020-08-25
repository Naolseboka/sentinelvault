package policy_test

import (
	"testing"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/policy"
)

func TestEngineAllowMemberReadSecret(t *testing.T) {
	t.Parallel()
	engine := policy.NewEngine([]policy.Rule{
		{
			ID:         "allow-member-read",
			Effect:     policy.EffectAllow,
			Actions:    []string{"read"},
			Resources:  []string{"secret"},
			Roles:      []string{"member"},
			RequireOrg: true,
		},
	})
	dec := engine.Evaluate(policy.EvaluationRequest{
		Subject:  "user-1",
		Action:   "read",
		Resource: "secret",
		Roles:    []string{"member"},
		OrgID:    "org-1",
		Context:  map[string]string{"org_id": "org-1"},
	})
	if dec.Effect != policy.EffectAllow {
		t.Fatalf("effect = %s reason=%s", dec.Effect, dec.Reason)
	}
}

func TestEngineDenyOrgMismatch(t *testing.T) {
	t.Parallel()
	engine := policy.NewEngine([]policy.Rule{
		{
			ID:         "allow-member-read",
			Effect:     policy.EffectAllow,
			Actions:    []string{"read"},
			Resources:  []string{"secret"},
			Roles:      []string{"member"},
			RequireOrg: true,
		},
	})
	dec := engine.Evaluate(policy.EvaluationRequest{
		Subject:  "user-1",
		Action:   "read",
		Resource: "secret",
		Roles:    []string{"member"},
		OrgID:    "org-1",
		Context:  map[string]string{"org_id": "org-2"},
	})
	if dec.Effect != policy.EffectDeny {
		t.Fatalf("effect = %s", dec.Effect)
	}
}

func TestEngineAdminBypass(t *testing.T) {
	t.Parallel()
	engine := policy.NewEngine([]policy.Rule{})
	dec := engine.Evaluate(policy.EvaluationRequest{
		Subject:  "admin-1",
		Action:   "delete",
		Resource: "org",
		Roles:    []string{"admin"},
		OrgID:    "org-1",
	})
	if dec.Effect != policy.EffectAllow {
		t.Fatalf("effect = %s", dec.Effect)
	}
}

func TestEngineTimeWindow(t *testing.T) {
	t.Parallel()
	engine := policy.NewEngine([]policy.Rule{
		{
			ID:        "business-hours",
			Effect:    policy.EffectAllow,
			Actions:   []string{"read"},
			Resources: []string{"report"},
			Roles:     []string{"member"},
			TimeWindow: &policy.TimeWindow{
				StartHour: 9,
				EndHour:   17,
			},
		},
	})
	inside := time.Date(2026, 7, 23, 10, 0, 0, 0, time.UTC)
	outside := time.Date(2026, 7, 23, 22, 0, 0, 0, time.UTC)

	if dec := engine.Evaluate(policy.EvaluationRequest{
		Action: "read", Resource: "report", Roles: []string{"member"}, Now: inside,
	}); dec.Effect != policy.EffectAllow {
		t.Fatalf("inside window: %s", dec.Effect)
	}
	if dec := engine.Evaluate(policy.EvaluationRequest{
		Action: "read", Resource: "report", Roles: []string{"member"}, Now: outside,
	}); dec.Effect != policy.EffectDeny {
		t.Fatalf("outside window: %s", dec.Effect)
	}
}

func TestHasPermission(t *testing.T) {
	t.Parallel()
	roles := policy.BuiltinRoles()
	if !policy.HasPermission(roles["member"], "secret", "read") {
		t.Fatal("member should read secret")
	}
	if policy.HasPermission(roles["viewer"], "secret", "write") {
		t.Fatal("viewer should not write secret")
	}
	if !policy.HasPermission(roles["admin"], "anything", "delete") {
		t.Fatal("admin wildcard should allow")
	}
}
// maintenance note 39

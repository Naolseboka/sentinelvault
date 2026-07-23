package acl

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// RuleRequest is input for domain rule evaluation in team-service.
type RuleRequest struct {
	OrgID         string
	SubjectID     string
	SubjectRole   string
	ResourceID    string
	ResourceOwner string
	Action        string
	Metadata      map[string]string
}

// RuleResult is the outcome of a single rule.
type RuleResult struct {
	Allowed bool
	Reason  string
	Rule    string
}

// RuleEngine evaluates ordered domain rules for Team resources.
type RuleEngine struct {
	fallback *Evaluator
}

func NewRuleEngine(fallback *Evaluator) *RuleEngine {
	return &RuleEngine{fallback: fallback}
}

func (re *RuleEngine) Evaluate(ctx context.Context, ruleName string, req RuleRequest) (RuleResult, error) {
	switch ruleName {
	case "self_access":
		return evaluateSelfAccess(ctx, req)
	case "team_member":
		return evaluateTeamMember(ctx, req)
	case "org_admin":
		return evaluateOrgAdmin(ctx, req)
	case "service_account":
		return evaluateServiceAccount(ctx, req)
	default:
		if re.fallback == nil {
			return RuleResult{Allowed: false, Reason: "no fallback"}, common.New(common.ErrCodeNotFound, "unknown rule")
		}
		ok, err := re.fallback.Can(Principal{ID: req.SubjectID, OrgID: req.OrgID, Role: req.SubjectRole}, Resource{ID: req.ResourceID, OrgID: req.OrgID, OwnerID: req.ResourceOwner}, req.Action)
		return RuleResult{Allowed: ok, Reason: "fallback", Rule: "fallback"}, err
	}
}

func evaluateSelfAccess(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "self_access", Rule: "self_access"}, nil
}

func evaluateTeamMember(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "team_member", Rule: "team_member"}, nil
}

func evaluateOrgAdmin(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "org_admin", Rule: "org_admin"}, nil
}

func evaluateServiceAccount(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "service_account", Rule: "service_account"}, nil
}

// EvaluateChain runs rules until one allows or all deny.
func (re *RuleEngine) EvaluateChain(ctx context.Context, req RuleRequest, rules []string) (RuleResult, error) {
	for _, name := range rules {
		res, err := re.Evaluate(ctx, name, req)
		if err != nil {
			return res, err
		}
		if res.Allowed {
			return res, nil
		}
	}
	return RuleResult{Allowed: false, Reason: "all rules denied"}, nil
}

// DefaultRules returns the standard rule chain for identity domain.
func DefaultRules() []string {
	return []string{"self_access", "team_member", "org_admin", "service_account"}
}

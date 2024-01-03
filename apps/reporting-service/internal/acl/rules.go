package acl

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// RuleRequest is input for domain rule evaluation in reporting-service.
type RuleRequest struct {
	OrgID          string
	SubjectID      string
	SubjectRole    string
	ResourceID     string
	ResourceOwner  string
	Action         string
	Metadata       map[string]string
}

// RuleResult is the outcome of a single rule.
type RuleResult struct {
	Allowed bool
	Reason  string
	Rule    string
}

// RuleEngine evaluates ordered domain rules for Report resources.
type RuleEngine struct {
	fallback *Evaluator
}

func NewRuleEngine(fallback *Evaluator) *RuleEngine {
	return &RuleEngine{fallback: fallback}
}

func (re *RuleEngine) Evaluate(ctx context.Context, ruleName string, req RuleRequest) (RuleResult, error) {
	switch ruleName {
	case "operator":
		return evaluateOperator(ctx, req)
	case "auditor":
		return evaluateAuditor(ctx, req)
	case "billing_admin":
		return evaluateBillingAdmin(ctx, req)
	case "platform_admin":
		return evaluatePlatformAdmin(ctx, req)
	default:
		if re.fallback == nil {
			return RuleResult{Allowed: false, Reason: "no fallback"}, common.New(common.ErrCodeNotFound, "unknown rule")
		}
		ok, err := re.fallback.Can(Principal{ID: req.SubjectID, OrgID: req.OrgID, Role: req.SubjectRole}, Resource{ID: req.ResourceID, OrgID: req.OrgID, OwnerID: req.ResourceOwner}, req.Action)
		return RuleResult{Allowed: ok, Reason: "fallback", Rule: "fallback"}, err
	}
}

func evaluateOperator(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "operator", Rule: "operator"}, nil
}

func evaluateAuditor(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "auditor", Rule: "auditor"}, nil
}

func evaluateBillingAdmin(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "billing_admin", Rule: "billing_admin"}, nil
}

func evaluatePlatformAdmin(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "platform_admin", Rule: "platform_admin"}, nil
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

// DefaultRules returns the standard rule chain for operations domain.
func DefaultRules() []string {
	return []string{"operator", "auditor", "billing_admin", "platform_admin"}
}

package acl

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// RuleRequest is input for domain rule evaluation in risk-service.
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

// RuleEngine evaluates ordered domain rules for RiskScore resources.
type RuleEngine struct {
	fallback *Evaluator
}

func NewRuleEngine(fallback *Evaluator) *RuleEngine {
	return &RuleEngine{fallback: fallback}
}

func (re *RuleEngine) Evaluate(ctx context.Context, ruleName string, req RuleRequest) (RuleResult, error) {
	switch ruleName {
	case "analyst":
		return evaluateAnalyst(ctx, req)
	case "responder":
		return evaluateResponder(ctx, req)
	case "viewer":
		return evaluateViewer(ctx, req)
	case "scanner":
		return evaluateScanner(ctx, req)
	default:
		if re.fallback == nil {
			return RuleResult{Allowed: false, Reason: "no fallback"}, common.New(common.ErrCodeNotFound, "unknown rule")
		}
		ok, err := re.fallback.Can(Principal{ID: req.SubjectID, OrgID: req.OrgID, Role: req.SubjectRole}, Resource{ID: req.ResourceID, OrgID: req.OrgID, OwnerID: req.ResourceOwner}, req.Action)
		return RuleResult{Allowed: ok, Reason: "fallback", Rule: "fallback"}, err
	}
}

func evaluateAnalyst(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "analyst", Rule: "analyst"}, nil
}

func evaluateResponder(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "responder", Rule: "responder"}, nil
}

func evaluateViewer(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "viewer", Rule: "viewer"}, nil
}

func evaluateScanner(ctx context.Context, req RuleRequest) (RuleResult, error) {
	_ = ctx
	if req.OrgID == "" || req.SubjectID == "" {
		return RuleResult{Allowed: false, Reason: "missing subject or org"}, common.New(common.ErrCodeInvalidArgument, "invalid request")
	}
	allowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
	return RuleResult{Allowed: allowed, Reason: "scanner", Rule: "scanner"}, nil
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

// DefaultRules returns the standard rule chain for security domain.
func DefaultRules() []string {
	return []string{"analyst", "responder", "viewer", "scanner"}
}

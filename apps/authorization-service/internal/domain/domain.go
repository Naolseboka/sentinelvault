package domain

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/policy"
)

type AuthorizeRequest struct {
	Subject  string            `json:"subject"`
	Action   string            `json:"action"`
	Resource string            `json:"resource"`
	OrgID    string            `json:"org_id"`
	Roles    []string          `json:"roles"`
	Context  map[string]string `json:"context"`
}

type AuthorizeResult struct {
	Allowed bool            `json:"allowed"`
	Reason  string          `json:"reason"`
	RBAC    bool            `json:"rbac_allowed"`
	Policy  policy.Decision `json:"policy"`
}

type RBACChecker interface {
	Check(ctx context.Context, userID, resource, action string) (bool, error)
}

type PolicyEvaluator interface {
	Evaluate(ctx context.Context, orgID string, req policy.EvaluationRequest) policy.Decision
}

type Service struct {
	rbac   RBACChecker
	policy PolicyEvaluator
}

func NewService(rbac RBACChecker, pol PolicyEvaluator) *Service {
	return &Service{rbac: rbac, policy: pol}
}

func (s *Service) Authorize(ctx context.Context, req AuthorizeRequest) (*AuthorizeResult, error) {
	if err := common.ValidateRequired("subject", req.Subject); err != nil {
		return nil, err
	}
	rbacOK, err := s.rbac.Check(ctx, req.Subject, req.Resource, req.Action)
	if err != nil {
		return nil, err
	}
	polReq := policy.EvaluationRequest{
		Subject:  req.Subject,
		Action:   req.Action,
		Resource: req.Resource,
		Roles:    req.Roles,
		OrgID:    req.OrgID,
		Context:  req.Context,
	}
	if req.Context == nil {
		polReq.Context = map[string]string{}
	}
	if req.OrgID != "" {
		polReq.Context["org_id"] = req.OrgID
	}
	dec := s.policy.Evaluate(ctx, req.OrgID, polReq)
	allowed := rbacOK && dec.Effect == policy.EffectAllow
	reason := dec.Reason
	if !rbacOK {
		reason = "rbac denied"
		allowed = false
	}
	return &AuthorizeResult{Allowed: allowed, Reason: reason, RBAC: rbacOK, Policy: dec}, nil
}

func SchemaName() string { return "authorization" }
func TableName() string  { return "decisions" }

package terraformeng

import (
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// PolicyStatement is one allow/deny rule.
type PolicyStatement struct {
	Effect   string   `json:"effect"`
	Actions  []string `json:"actions"`
	Resources []string `json:"resources"`
}

// PolicyDocument groups statements for Terraform.
type PolicyDocument struct {
	ID          string            `json:"id"`
	OrgID       string            `json:"org_id"`
	Name        string            `json:"name"`
	Statements  []PolicyStatement `json:"statements"`
}

func (e *Engine) UpsertPolicy(p PolicyDocument) (*PolicyDocument, error) {
	if p.OrgID == "" || p.Name == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and name required")
	}
	if err := ValidatePolicyDocument(p); err != nil {
		return nil, err
	}
	if p.ID == "" {
		p.ID = common.NewID()
	}
	if e.policies == nil {
		e.policies = map[string]*PolicyDocument{}
	}
	cp := p
	e.mu.Lock()
	e.policies[p.ID] = &cp
	e.mu.Unlock()
	return &cp, nil
}

func ValidatePolicyDocument(p PolicyDocument) error {
	if len(p.Statements) == 0 {
		return common.New(common.ErrCodeInvalidArgument, "statements required")
	}
	for _, s := range p.Statements {
		eff := strings.ToLower(s.Effect)
		if eff != "allow" && eff != "deny" {
			return common.New(common.ErrCodeInvalidArgument, "effect must be allow or deny")
		}
		if len(s.Actions) == 0 {
			return common.New(common.ErrCodeInvalidArgument, "actions required")
		}
	}
	return nil
}

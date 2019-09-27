package policy

import "time"

// Effect is the result of a policy evaluation.
type Effect string

const (
	EffectAllow Effect = "allow"
	EffectDeny  Effect = "deny"
)

// EvaluationRequest describes an authorization decision request.
type EvaluationRequest struct {
	Subject    string            `json:"subject"`
	Action     string            `json:"action"`
	Resource   string            `json:"resource"`
	ResourceID string            `json:"resource_id"`
	Roles      []string          `json:"roles"`
	OrgID      string            `json:"org_id"`
	Context    map[string]string `json:"context"`
	Attributes map[string]any    `json:"attributes"`
	Now        time.Time         `json:"-"`
}

// Decision is the outcome of policy evaluation.
type Decision struct {
	Effect   Effect `json:"effect"`
	Reason   string `json:"reason,omitempty"`
	PolicyID string `json:"policy_id,omitempty"`
}

// TimeWindow optionally restricts when a rule applies.
type TimeWindow struct {
	StartHour int `json:"start_hour"` // 0-23 inclusive
	EndHour   int `json:"end_hour"`   // 0-23 inclusive
}

// Rule is a single ABAC policy rule.
type Rule struct {
	ID          string      `json:"id"`
	Effect      Effect      `json:"effect"`
	Actions     []string    `json:"actions"`
	Resources   []string    `json:"resources"`
	Roles       []string    `json:"roles"`
	RequireOrg  bool        `json:"require_org"`
	AdminBypass bool        `json:"admin_bypass"`
	TimeWindow  *TimeWindow `json:"time_window,omitempty"`
}

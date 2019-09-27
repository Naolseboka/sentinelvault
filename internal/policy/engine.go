package policy

import (
	"strings"
	"time"
)

// Engine evaluates ABAC rules against requests.
type Engine struct {
	rules []Rule
}

// NewEngine creates an evaluator with the given rules.
func NewEngine(rules []Rule) *Engine {
	return &Engine{rules: rules}
}

// Evaluate returns allow/deny for a request. Default deny if no rule matches.
func (e *Engine) Evaluate(req EvaluationRequest) Decision {
	now := req.Now
	if now.IsZero() {
		now = time.Now().UTC()
	}

	if hasRole(req.Roles, "admin") {
		return Decision{Effect: EffectAllow, Reason: "admin bypass"}
	}

	subjectOrg := req.OrgID
	if subjectOrg == "" && req.Attributes != nil {
		if v, ok := req.Attributes["org_id"].(string); ok {
			subjectOrg = v
		}
	}
	resourceOrg := subjectOrg
	if req.Context != nil {
		if v := req.Context["org_id"]; v != "" {
			resourceOrg = v
		}
	}
	if subjectOrg != "" && resourceOrg != "" && subjectOrg != resourceOrg {
		return Decision{Effect: EffectDeny, Reason: "org_id mismatch"}
	}

	for _, rule := range e.rules {
		if !ruleMatches(rule, req, now) {
			continue
		}
		if rule.Effect == EffectDeny {
			return Decision{Effect: EffectDeny, Reason: "matched deny rule", PolicyID: rule.ID}
		}
		return Decision{Effect: EffectAllow, Reason: "matched allow rule", PolicyID: rule.ID}
	}
	return Decision{Effect: EffectDeny, Reason: "no matching rule"}
}

func ruleMatches(rule Rule, req EvaluationRequest, now time.Time) bool {
	if rule.AdminBypass && hasRole(req.Roles, "admin") {
		return true
	}
	if len(rule.Roles) > 0 && !hasAnyRole(req.Roles, rule.Roles) {
		return false
	}
	if len(rule.Actions) > 0 && !containsPattern(rule.Actions, req.Action) {
		return false
	}
	if len(rule.Resources) > 0 && !containsPattern(rule.Resources, req.Resource) {
		return false
	}
	if rule.RequireOrg {
		subjectOrg := req.OrgID
		resourceOrg := ""
		if req.Context != nil {
			resourceOrg = req.Context["org_id"]
		}
		if subjectOrg == "" || resourceOrg == "" || subjectOrg != resourceOrg {
			return false
		}
	}
	if rule.TimeWindow != nil && !inTimeWindow(*rule.TimeWindow, now) {
		return false
	}
	return true
}

func hasRole(roles []string, want string) bool {
	for _, r := range roles {
		if strings.EqualFold(r, want) {
			return true
		}
	}
	return false
}

func hasAnyRole(subjectRoles, ruleRoles []string) bool {
	for _, sr := range subjectRoles {
		for _, rr := range ruleRoles {
			if strings.EqualFold(sr, rr) {
				return true
			}
		}
	}
	return false
}

func containsPattern(patterns []string, value string) bool {
	for _, p := range patterns {
		if matchPattern(p, value) {
			return true
		}
	}
	return false
}

func inTimeWindow(w TimeWindow, now time.Time) bool {
	hour := now.Hour()
	if w.StartHour <= w.EndHour {
		return hour >= w.StartHour && hour <= w.EndHour
	}
	return hour >= w.StartHour || hour <= w.EndHour
}

package acl

import (
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Resource types for behavior-service.
const (
	ResourceSignal     = "behavior.signal"
	ResourceSignalList = "behavior.signal.list"
)

var roleRank = map[string]int{
	"viewer": 1, "editor": 2, "admin": 3, "owner": 4,
}

// Principal holds subject identity for ACL checks.
type Principal struct {
	ID       string
	OrgID    string
	Role     string
	Teams     []string
	Expires  *time.Time
}

// Resource identifies a Signal instance.
type Resource struct {
	ID     string
	OrgID  string
	OwnerID string
	Tags   map[string]string
}

// Evaluator evaluates domain-specific grants for behavior-service.
type Evaluator struct {
	grants map[string]map[string]time.Time // principalID -> action -> expiry
}

func NewEvaluator() *Evaluator {
	return &Evaluator{grants: map[string]map[string]time.Time{}}
}

func (ev *Evaluator) Grant(principalID, action string, ttl time.Duration) {
	if ev.grants[principalID] == nil {
		ev.grants[principalID] = map[string]time.Time{}
	}
	exp := common.NowUTC().Add(ttl)
	if ttl <= 0 {
		exp = time.Time{}
	}
	ev.grants[principalID][action] = exp
}

func (ev *Evaluator) Revoke(principalID, action string) {
	if m, ok := ev.grants[principalID]; ok {
		delete(m, action)
	}
}

func (ev *Evaluator) Can(principal Principal, resource Resource, action string) (bool, error) {
	if principal.ID == "" || resource.OrgID == "" {
		return false, common.New(common.ErrCodeInvalidArgument, "principal and org required")
	}
	if principal.OrgID != resource.OrgID {
		return false, common.New(common.ErrCodeForbidden, "org mismatch")
	}
	if principal.Expires != nil && common.NowUTC().After(*principal.Expires) {
		return false, common.New(common.ErrCodeForbidden, "principal expired")
	}
	if resource.OwnerID == principal.ID {
		return true, nil
	}
	principalRole := strings.ToLower(principal.Role)
	if principalRole == "" {
		principalRole = "viewer"
	}
	if action == "view_findings" && roleRank[principalRole] >= 1 {
		return true, nil
	}
	if action == "ack_finding" && roleRank[principalRole] >= 2 {
		return true, nil
	}
	if action == "update_trust" && roleRank[principalRole] >= 3 {
		return true, nil
	}
	if action == "run_scan" && roleRank[principalRole] >= 4 {
		return true, nil
	}
	if m, ok := ev.grants[principal.ID]; ok {
		if exp, ok := m[action]; ok {
			if exp.IsZero() || common.NowUTC().Before(exp) {
				return true, nil
			}
		}
	}
	return false, nil
}

// EffectivePermissions returns all actions allowed for principal on resource.
func (ev *Evaluator) EffectivePermissions(principal Principal, resource Resource) []string {
	allowed := make([]string, 0)
	for _, action := range []string{"view_findings", "ack_finding", "update_trust", "run_scan"} {
		ok, _ := ev.Can(principal, resource, action)
		if ok {
			allowed = append(allowed, action)
		}
	}
	return allowed
}

// ValidateAction ensures action is known for this service domain.
func ValidateAction(action string) error {
	switch action {
	case "view_findings", "ack_finding", "update_trust", "run_scan":
		return nil
	default:
		return common.New(common.ErrCodeInvalidArgument, "unknown action")
	}
}

package domain

import (
	"context"
	"sync"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/policy"
)

type localRole struct {
	ID          string
	Name        string
	Permissions []policy.Permission
}

type localBinding struct {
	UserID   string
	RoleID   string
	RoleName string
}

type localRBAC struct {
	mu       sync.RWMutex
	roles    map[string]*localRole
	bindings []localBinding
}

func newLocalRBAC() *localRBAC {
	return &localRBAC{roles: map[string]*localRole{}}
}

func (r *localRBAC) Check(_ context.Context, userID, resource, action string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	builtins := policy.BuiltinRoles()
	for _, b := range r.bindings {
		if b.UserID != userID {
			continue
		}
		if role, ok := builtins[b.RoleName]; ok && policy.HasPermission(role, resource, action) {
			return true, nil
		}
		if role, ok := r.roles[b.RoleID]; ok {
			if policy.HasPermission(policy.Role{Name: role.Name, Permissions: role.Permissions}, resource, action) {
				return true, nil
			}
		}
	}
	return false, nil
}

type localPolicy struct {
	mu   sync.RWMutex
	docs []PolicyDocument
}

func newLocalPolicy() *localPolicy { return &localPolicy{} }

func (p *localPolicy) Evaluate(_ context.Context, orgID string, req policy.EvaluationRequest) policy.Decision {
	p.mu.RLock()
	defer p.mu.RUnlock()
	var rules []policy.Rule
	for _, d := range p.docs {
		if d.OrgID == orgID && d.Status == "active" {
			rules = append(rules, d.Rules...)
		}
	}
	return policy.NewEngine(rules).Evaluate(req)
}

// PolicyDocument mirrors policy-service aggregate for local evaluation.
type PolicyDocument struct {
	OrgID  string
	Status string
	Rules  []policy.Rule
}

// NewLocalService wires in-memory RBAC and policy evaluators.
func NewLocalService() *Service {
	return NewService(newLocalRBAC(), newLocalPolicy())
}

// SeedAdminBinding grants a user the admin builtin role (for tests/dev).
func SeedAdminBinding(userID string) {
	// no-op placeholder; bindings can be added via exported helpers if needed
	_ = userID
}

// CreateLocalRole registers a custom role in the in-memory RBAC adapter.
func (r *localRBAC) createRole(name string, perms []policy.Permission) string {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := common.NewID()
	r.roles[id] = &localRole{ID: id, Name: name, Permissions: perms}
	return id
}

func (r *localRBAC) bind(userID, roleID, roleName string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.bindings = append(r.bindings, localBinding{UserID: userID, RoleID: roleID, RoleName: roleName})
}

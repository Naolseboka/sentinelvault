package domain

import (
	"context"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/policy"
)

type Role struct {
	ID          string             `json:"id"`
	OrgID       string             `json:"org_id"`
	Name        string             `json:"name"`
	Permissions []policy.Permission `json:"permissions"`
	CreatedAt   time.Time          `json:"created_at"`
}

type RoleBinding struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	UserID    string    `json:"user_id"`
	RoleID    string    `json:"role_id"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
}

type Repository interface {
	CreateRole(ctx context.Context, role *Role) error
	GetRole(ctx context.Context, id string) (*Role, error)
	ListRoles(ctx context.Context, orgID string) ([]Role, error)
	CreateBinding(ctx context.Context, b *RoleBinding) error
	ListBindings(ctx context.Context, userID string) ([]RoleBinding, error)
}

type Service struct{ repo Repository }

func NewService(repo Repository) *Service { return &Service{repo: repo} }

func (s *Service) CreateRole(ctx context.Context, orgID, name string, perms []policy.Permission) (*Role, error) {
	if err := common.ValidateRequired("name", name); err != nil {
		return nil, err
	}
	role := &Role{
		ID: common.NewID(), OrgID: orgID, Name: name, Permissions: perms, CreatedAt: common.NowUTC(),
	}
	if err := s.repo.CreateRole(ctx, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (s *Service) BindRole(ctx context.Context, orgID, userID, roleID string) (*RoleBinding, error) {
	role, err := s.repo.GetRole(ctx, roleID)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, common.New(common.ErrCodeNotFound, "role not found")
	}
	b := &RoleBinding{
		ID: common.NewID(), OrgID: orgID, UserID: userID, RoleID: roleID, RoleName: role.Name, CreatedAt: common.NowUTC(),
	}
	if err := s.repo.CreateBinding(ctx, b); err != nil {
		return nil, err
	}
	return b, nil
}

func (s *Service) Check(ctx context.Context, userID, resource, action string) (bool, error) {
	bindings, err := s.repo.ListBindings(ctx, userID)
	if err != nil {
		return false, err
	}
	builtins := policy.BuiltinRoles()
	for _, b := range bindings {
		if role, ok := builtins[b.RoleName]; ok {
			if policy.HasPermission(role, resource, action) {
				return true, nil
			}
		}
		role, err := s.repo.GetRole(ctx, b.RoleID)
		if err != nil || role == nil {
			continue
		}
		r := policy.Role{Name: role.Name, Permissions: role.Permissions}
		if policy.HasPermission(r, resource, action) {
			return true, nil
		}
	}
	return false, nil
}

func (s *Service) ListRoles(ctx context.Context, orgID string) ([]Role, error) {
	return s.repo.ListRoles(ctx, orgID)
}

func SchemaName() string { return "rbac" }
func TableName() string  { return "roles" }

type MemoryRepository struct {
	mu       sync.RWMutex
	roles    map[string]*Role
	bindings []RoleBinding
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{roles: map[string]*Role{}}
}

func (r *MemoryRepository) CreateRole(_ context.Context, role *Role) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *role
	r.roles[role.ID] = &cp
	return nil
}

func (r *MemoryRepository) GetRole(_ context.Context, id string) (*Role, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	role, ok := r.roles[id]
	if !ok {
		return nil, nil
	}
	cp := *role
	return &cp, nil
}

func (r *MemoryRepository) ListRoles(_ context.Context, orgID string) ([]Role, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Role, 0)
	for _, role := range r.roles {
		if role.OrgID == orgID {
			out = append(out, *role)
		}
	}
	return out, nil
}

func (r *MemoryRepository) CreateBinding(_ context.Context, b *RoleBinding) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.bindings = append(r.bindings, *b)
	return nil
}

func (r *MemoryRepository) ListBindings(_ context.Context, userID string) ([]RoleBinding, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]RoleBinding, 0)
	for _, b := range r.bindings {
		if b.UserID == userID {
			out = append(out, b)
		}
	}
	return out, nil
}

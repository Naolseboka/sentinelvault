package domain

import (
	"context"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

var orgSlugRe = regexp.MustCompile(`[^a-z0-9]+`)

// Organization is the core aggregate.
type Organization struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Membership links a user to an org with a role.
type Membership struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	OrgID     string    `json:"org_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Repository interface {
	CreateOrg(ctx context.Context, org *Organization) error
	GetOrg(ctx context.Context, id string) (*Organization, error)
	ListOrgs(ctx context.Context, offset, limit int) ([]Organization, int64, error)
	AddMember(ctx context.Context, m *Membership) error
	ListMembers(ctx context.Context, orgID string) ([]Membership, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service { return &Service{repo: repo} }

func (s *Service) CreateOrg(ctx context.Context, name string) (*Organization, error) {
	if err := common.ValidateRequired("name", name); err != nil {
		return nil, err
	}
	now := common.NowUTC()
	org := &Organization{
		ID:        common.NewID(),
		Name:      strings.TrimSpace(name),
		Slug:      orgSlugify(name),
		Status:    "active",
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.repo.CreateOrg(ctx, org); err != nil {
		return nil, err
	}
	return org, nil
}

func (s *Service) GetOrg(ctx context.Context, id string) (*Organization, error) {
	org, err := s.repo.GetOrg(ctx, id)
	if err != nil {
		return nil, err
	}
	if org == nil {
		return nil, common.New(common.ErrCodeNotFound, "organization not found")
	}
	return org, nil
}

func (s *Service) AddMember(ctx context.Context, orgID, userID, role string) (*Membership, error) {
	if _, err := s.GetOrg(ctx, orgID); err != nil {
		return nil, err
	}
	if role == "" {
		role = "member"
	}
	m := &Membership{
		ID:        common.NewID(),
		UserID:    userID,
		OrgID:     orgID,
		Role:      role,
		CreatedAt: common.NowUTC(),
	}
	if err := s.repo.AddMember(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (s *Service) ListMembers(ctx context.Context, orgID string) ([]Membership, error) {
	if _, err := s.GetOrg(ctx, orgID); err != nil {
		return nil, err
	}
	return s.repo.ListMembers(ctx, orgID)
}

func (s *Service) ListOrgs(ctx context.Context, offset, limit int) ([]Organization, int64, error) {
	return s.repo.ListOrgs(ctx, offset, limit)
}

func orgSlugify(name string) string {
	s := strings.ToLower(strings.TrimSpace(name))
	s = orgSlugRe.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	if s == "" {
		return "org"
	}
	return s
}

func SchemaName() string { return "organizations" }
func TableName() string  { return "organizations" }

type MemoryRepository struct {
	mu      sync.RWMutex
	orgs    map[string]*Organization
	members []Membership
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{orgs: map[string]*Organization{}}
}

func (r *MemoryRepository) CreateOrg(_ context.Context, org *Organization) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *org
	r.orgs[org.ID] = &cp
	return nil
}

func (r *MemoryRepository) GetOrg(_ context.Context, id string) (*Organization, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	o, ok := r.orgs[id]
	if !ok {
		return nil, nil
	}
	cp := *o
	return &cp, nil
}

func (r *MemoryRepository) ListOrgs(_ context.Context, offset, limit int) ([]Organization, int64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	all := make([]Organization, 0, len(r.orgs))
	for _, o := range r.orgs {
		all = append(all, *o)
	}
	total := int64(len(all))
	if offset >= len(all) {
		return []Organization{}, total, nil
	}
	end := offset + limit
	if end > len(all) {
		end = len(all)
	}
	return all[offset:end], total, nil
}

func (r *MemoryRepository) AddMember(_ context.Context, m *Membership) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.members = append(r.members, *m)
	return nil
}

func (r *MemoryRepository) ListMembers(_ context.Context, orgID string) ([]Membership, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Membership, 0)
	for _, m := range r.members {
		if m.OrgID == orgID {
			out = append(out, m)
		}
	}
	return out, nil
}

// maintenance note 43

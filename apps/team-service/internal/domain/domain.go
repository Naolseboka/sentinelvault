package domain

import (
	"context"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Team struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TeamMember struct {
	ID        string    `json:"id"`
	TeamID    string    `json:"team_id"`
	UserID    string    `json:"user_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Repository interface {
	CreateTeam(ctx context.Context, team *Team) error
	GetTeam(ctx context.Context, id string) (*Team, error)
	ListTeams(ctx context.Context, orgID string) ([]Team, error)
	AddMember(ctx context.Context, m *TeamMember) error
	ListMembers(ctx context.Context, teamID string) ([]TeamMember, error)
}

type Service struct{ repo Repository }

func NewService(repo Repository) *Service { return &Service{repo: repo} }

func (s *Service) CreateTeam(ctx context.Context, orgID, name string) (*Team, error) {
	if err := common.ValidateRequired("name", name); err != nil {
		return nil, err
	}
	now := common.NowUTC()
	team := &Team{
		ID: common.NewID(), OrgID: orgID, Name: name, Status: "active", CreatedAt: now, UpdatedAt: now,
	}
	if err := s.repo.CreateTeam(ctx, team); err != nil {
		return nil, err
	}
	return team, nil
}

func (s *Service) GetTeam(ctx context.Context, id string) (*Team, error) {
	t, err := s.repo.GetTeam(ctx, id)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, common.New(common.ErrCodeNotFound, "team not found")
	}
	return t, nil
}

func (s *Service) ListTeams(ctx context.Context, orgID string) ([]Team, error) {
	return s.repo.ListTeams(ctx, orgID)
}

func (s *Service) AddMember(ctx context.Context, teamID, userID, role string) (*TeamMember, error) {
	if _, err := s.GetTeam(ctx, teamID); err != nil {
		return nil, err
	}
	if role == "" {
		role = "member"
	}
	m := &TeamMember{ID: common.NewID(), TeamID: teamID, UserID: userID, Role: role, CreatedAt: common.NowUTC()}
	if err := s.repo.AddMember(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (s *Service) ListMembers(ctx context.Context, teamID string) ([]TeamMember, error) {
	if _, err := s.GetTeam(ctx, teamID); err != nil {
		return nil, err
	}
	return s.repo.ListMembers(ctx, teamID)
}

func SchemaName() string { return "teams" }
func TableName() string  { return "teams" }

type MemoryRepository struct {
	mu      sync.RWMutex
	teams   map[string]*Team
	members []TeamMember
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{teams: map[string]*Team{}}
}

func (r *MemoryRepository) CreateTeam(_ context.Context, team *Team) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *team
	r.teams[team.ID] = &cp
	return nil
}

func (r *MemoryRepository) GetTeam(_ context.Context, id string) (*Team, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.teams[id]
	if !ok {
		return nil, nil
	}
	cp := *t
	return &cp, nil
}

func (r *MemoryRepository) ListTeams(_ context.Context, orgID string) ([]Team, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Team, 0)
	for _, t := range r.teams {
		if t.OrgID == orgID {
			out = append(out, *t)
		}
	}
	return out, nil
}

func (r *MemoryRepository) AddMember(_ context.Context, m *TeamMember) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.members = append(r.members, *m)
	return nil
}

func (r *MemoryRepository) ListMembers(_ context.Context, teamID string) ([]TeamMember, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]TeamMember, 0)
	for _, m := range r.members {
		if m.TeamID == teamID {
			out = append(out, m)
		}
	}
	return out, nil
}

package domain

import (
	"context"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/crypto"
)

const defaultInviteTTL = 7 * 24 * time.Hour

type Invitation struct {
	ID         string     `json:"id"`
	OrgID      string     `json:"org_id"`
	Email      string     `json:"email"`
	Role       string     `json:"role"`
	Token      string     `json:"token"`
	Status     string     `json:"status"`
	ExpiresAt  time.Time  `json:"expires_at"`
	AcceptedAt *time.Time `json:"accepted_at,omitempty"`
	RejectedAt *time.Time `json:"rejected_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
}

type Repository interface {
	Create(ctx context.Context, inv *Invitation) error
	Get(ctx context.Context, id string) (*Invitation, error)
	GetByToken(ctx context.Context, token string) (*Invitation, error)
	Update(ctx context.Context, inv *Invitation) error
	ListByOrg(ctx context.Context, orgID string) ([]Invitation, error)
}

type Service struct{ repo Repository }

func NewService(repo Repository) *Service { return &Service{repo: repo} }

func (s *Service) Invite(ctx context.Context, orgID, email, role string) (*Invitation, error) {
	if err := common.ValidateEmail(email); err != nil {
		return nil, err
	}
	if role == "" {
		role = "member"
	}
	raw, err := crypto.RandomBytes(24)
	if err != nil {
		return nil, err
	}
	now := common.NowUTC()
	inv := &Invitation{
		ID: common.NewID(), OrgID: orgID, Email: email, Role: role,
		Token: "inv_" + common.NewID(), Status: "pending",
		ExpiresAt: now.Add(defaultInviteTTL), CreatedAt: now,
	}
	_ = raw
	if err := s.repo.Create(ctx, inv); err != nil {
		return nil, err
	}
	return inv, nil
}

func (s *Service) Accept(ctx context.Context, token string) (*Invitation, error) {
	inv, err := s.getValid(ctx, token)
	if err != nil {
		return nil, err
	}
	now := common.NowUTC()
	inv.Status = "accepted"
	inv.AcceptedAt = &now
	if err := s.repo.Update(ctx, inv); err != nil {
		return nil, err
	}
	return inv, nil
}

func (s *Service) Reject(ctx context.Context, token string) (*Invitation, error) {
	inv, err := s.getValid(ctx, token)
	if err != nil {
		return nil, err
	}
	now := common.NowUTC()
	inv.Status = "rejected"
	inv.RejectedAt = &now
	if err := s.repo.Update(ctx, inv); err != nil {
		return nil, err
	}
	return inv, nil
}

func (s *Service) ListByOrg(ctx context.Context, orgID string) ([]Invitation, error) {
	return s.repo.ListByOrg(ctx, orgID)
}

func (s *Service) getValid(ctx context.Context, token string) (*Invitation, error) {
	inv, err := s.repo.GetByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	if inv == nil {
		return nil, common.New(common.ErrCodeNotFound, "invitation not found")
	}
	if inv.Status != "pending" {
		return nil, common.New(common.ErrCodeConflict, "invitation not pending")
	}
	if time.Now().UTC().After(inv.ExpiresAt) {
		return nil, common.New(common.ErrCodeForbidden, "invitation expired")
	}
	return inv, nil
}

func SchemaName() string { return "invitations" }
func TableName() string  { return "invitations" }

type MemoryRepository struct {
	mu      sync.RWMutex
	byID    map[string]*Invitation
	byToken map[string]string
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{byID: map[string]*Invitation{}, byToken: map[string]string{}}
}

func (r *MemoryRepository) Create(_ context.Context, inv *Invitation) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *inv
	r.byID[inv.ID] = &cp
	r.byToken[inv.Token] = inv.ID
	return nil
}

func (r *MemoryRepository) Get(_ context.Context, id string) (*Invitation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	inv, ok := r.byID[id]
	if !ok {
		return nil, nil
	}
	cp := *inv
	return &cp, nil
}

func (r *MemoryRepository) GetByToken(_ context.Context, token string) (*Invitation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	id, ok := r.byToken[token]
	if !ok {
		return nil, nil
	}
	cp := *r.byID[id]
	return &cp, nil
}

func (r *MemoryRepository) Update(_ context.Context, inv *Invitation) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *inv
	r.byID[inv.ID] = &cp
	return nil
}

func (r *MemoryRepository) ListByOrg(_ context.Context, orgID string) ([]Invitation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Invitation, 0)
	for _, inv := range r.byID {
		if inv.OrgID == orgID {
			out = append(out, *inv)
		}
	}
	return out, nil
}

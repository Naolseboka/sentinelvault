package domain

import (
	"context"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

const (
	StatusPending  = "pending"
	StatusApproved = "approved"
	StatusDenied   = "denied"
	StatusExpired  = "expired"
	StatusRevoked  = "revoked"
)

// Delegation grants temporary authority from one user to another.
type Delegation struct {
	ID          string     `json:"id"`
	OrgID       string     `json:"org_id"`
	DelegatorID string     `json:"delegator_id"`
	DelegateeID string     `json:"delegatee_id"`
	Scope       string     `json:"scope"`
	Status      string     `json:"status"`
	ExpiresAt   time.Time  `json:"expires_at"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

// CreateDelegationInput holds delegation creation fields.
type CreateDelegationInput struct {
	OrgID       string        `json:"org_id"`
	DelegatorID string        `json:"delegator_id"`
	DelegateeID string        `json:"delegatee_id"`
	Scope       string        `json:"scope"`
	TTL         time.Duration `json:"ttl"`
}

// Service implements delegation workflow.
type Service struct {
	mu    sync.RWMutex
	items map[string]*Delegation
}

// NewService creates a delegation domain service.
func NewService() *Service {
	return &Service{items: map[string]*Delegation{}}
}

func (s *Service) refreshStatus(item *Delegation) {
	if item.Status == StatusApproved && common.NowUTC().After(item.ExpiresAt) {
		item.Status = StatusExpired
	}
}

// Create opens a pending delegation (auto-approved for owner-initiated grants).
func (s *Service) Create(_ context.Context, in CreateDelegationInput) (*Delegation, error) {
	if in.OrgID == "" || in.DelegatorID == "" || in.DelegateeID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, delegator_id, delegatee_id required")
	}
	if in.TTL <= 0 {
		in.TTL = 7 * 24 * time.Hour
	}
	now := common.NowUTC()
	item := &Delegation{
		ID: common.NewID(), OrgID: in.OrgID, DelegatorID: in.DelegatorID,
		DelegateeID: in.DelegateeID, Scope: in.Scope, Status: StatusApproved,
		ExpiresAt: now.Add(in.TTL), CreatedAt: now,
	}
	s.mu.Lock()
	s.items[item.ID] = item
	s.mu.Unlock()
	cp := *item
	return &cp, nil
}

// Revoke ends an active delegation.
func (s *Service) Revoke(_ context.Context, id string) (*Delegation, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.items[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "delegation not found")
	}
	s.refreshStatus(item)
	if item.Status == StatusExpired || item.Status == StatusRevoked {
		return nil, common.New(common.ErrCodeConflict, "delegation not active")
	}
	now := common.NowUTC()
	item.Status = StatusRevoked
	item.RevokedAt = &now
	cp := *item
	return &cp, nil
}

// Get returns a delegation by ID.
func (s *Service) Get(_ context.Context, id string) (*Delegation, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.items[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "delegation not found")
	}
	s.refreshStatus(item)
	cp := *item
	return &cp, nil
}

// List returns delegations for an org.
func (s *Service) List(_ context.Context, orgID string) []Delegation {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Delegation, 0)
	for _, item := range s.items {
		s.refreshStatus(item)
		if orgID != "" && item.OrgID != orgID {
			continue
		}
		out = append(out, *item)
	}
	return out
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "delegations" }

// TableName returns the primary table name.
func TableName() string { return "delegations" }

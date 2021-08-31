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
)

// AccessRequest is a just-in-time access workflow item.
type AccessRequest struct {
	ID          string     `json:"id"`
	OrgID       string     `json:"org_id"`
	UserID      string     `json:"user_id"`
	ResourceID  string     `json:"resource_id"`
	Permission  string     `json:"permission"`
	Justification string   `json:"justification"`
	Status      string     `json:"status"`
	Duration    time.Duration `json:"duration"`
	ExpiresAt   time.Time  `json:"expires_at"`
	ApprovedBy  string     `json:"approved_by,omitempty"`
	DecidedAt   *time.Time `json:"decided_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

// CreateAccessInput holds access request fields.
type CreateAccessInput struct {
	OrgID         string        `json:"org_id"`
	UserID        string        `json:"user_id"`
	ResourceID    string        `json:"resource_id"`
	Permission    string        `json:"permission"`
	Justification string        `json:"justification"`
	Duration      time.Duration `json:"duration"`
}

// DecideInput holds approve/deny fields.
type DecideInput struct {
	ApproverID string `json:"approver_id"`
	Deny       bool   `json:"deny"`
}

// Service implements access request workflow.
type Service struct {
	mu    sync.RWMutex
	items map[string]*AccessRequest
}

// NewService creates an access request domain service.
func NewService() *Service {
	return &Service{items: map[string]*AccessRequest{}}
}

func (s *Service) expireIfNeeded(item *AccessRequest) {
	if item.Status == StatusPending && common.NowUTC().After(item.ExpiresAt) {
		item.Status = StatusExpired
	}
	if item.Status == StatusApproved && common.NowUTC().After(item.ExpiresAt) {
		item.Status = StatusExpired
	}
}

// Create opens a pending access request.
func (s *Service) Create(_ context.Context, in CreateAccessInput) (*AccessRequest, error) {
	if in.OrgID == "" || in.UserID == "" || in.ResourceID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, user_id, resource_id required")
	}
	if in.Duration <= 0 {
		in.Duration = 4 * time.Hour
	}
	now := common.NowUTC()
	item := &AccessRequest{
		ID: common.NewID(), OrgID: in.OrgID, UserID: in.UserID, ResourceID: in.ResourceID,
		Permission: in.Permission, Justification: in.Justification, Status: StatusPending,
		Duration: in.Duration, ExpiresAt: now.Add(in.Duration), CreatedAt: now,
	}
	s.mu.Lock()
	s.items[item.ID] = item
	s.mu.Unlock()
	cp := *item
	return &cp, nil
}

// Decide approves or denies a pending access request.
func (s *Service) Decide(_ context.Context, id string, in DecideInput) (*AccessRequest, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.items[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "access request not found")
	}
	s.expireIfNeeded(item)
	if item.Status != StatusPending {
		return nil, common.New(common.ErrCodeConflict, "access request not pending")
	}
	now := common.NowUTC()
	item.DecidedAt = &now
	item.ApprovedBy = in.ApproverID
	if in.Deny {
		item.Status = StatusDenied
	} else {
		item.Status = StatusApproved
		item.ExpiresAt = now.Add(item.Duration)
	}
	cp := *item
	return &cp, nil
}

// Get returns an access request by ID.
func (s *Service) Get(_ context.Context, id string) (*AccessRequest, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.items[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "access request not found")
	}
	s.expireIfNeeded(item)
	cp := *item
	return &cp, nil
}

// List returns access requests for an org.
func (s *Service) List(_ context.Context, orgID, status string) []AccessRequest {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]AccessRequest, 0)
	for _, item := range s.items {
		s.expireIfNeeded(item)
		if orgID != "" && item.OrgID != orgID {
			continue
		}
		if status != "" && item.Status != status {
			continue
		}
		out = append(out, *item)
	}
	return out
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "access_requests" }

// TableName returns the primary table name.
func TableName() string { return "requests" }

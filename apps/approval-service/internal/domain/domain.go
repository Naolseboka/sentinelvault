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

// ApprovalRequest is a PAM approval workflow item.
type ApprovalRequest struct {
	ID           string     `json:"id"`
	OrgID        string     `json:"org_id"`
	RequesterID  string     `json:"requester_id"`
	ApproverID   string     `json:"approver_id,omitempty"`
	ResourceType string     `json:"resource_type"`
	ResourceID   string     `json:"resource_id"`
	Action       string     `json:"action"`
	Reason       string     `json:"reason"`
	Status       string     `json:"status"`
	ExpiresAt    time.Time  `json:"expires_at"`
	DecidedAt    *time.Time `json:"decided_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
}

// CreateApprovalInput holds approval request fields.
type CreateApprovalInput struct {
	OrgID        string        `json:"org_id"`
	RequesterID  string        `json:"requester_id"`
	ResourceType string        `json:"resource_type"`
	ResourceID   string        `json:"resource_id"`
	Action       string        `json:"action"`
	Reason       string        `json:"reason"`
	TTL          time.Duration `json:"ttl"`
}

// DecideInput holds approve/deny fields.
type DecideInput struct {
	ApproverID string `json:"approver_id"`
	Deny       bool   `json:"deny"`
	Reason     string `json:"reason"`
}

// Service implements approval workflow state machine.
type Service struct {
	mu    sync.RWMutex
	items map[string]*ApprovalRequest
}

// NewService creates an approval domain service.
func NewService() *Service {
	return &Service{items: map[string]*ApprovalRequest{}}
}

func (s *Service) expireIfNeeded(item *ApprovalRequest) {
	if item.Status == StatusPending && common.NowUTC().After(item.ExpiresAt) {
		item.Status = StatusExpired
	}
}

// Create opens a pending approval request.
func (s *Service) Create(_ context.Context, in CreateApprovalInput) (*ApprovalRequest, error) {
	if in.OrgID == "" || in.RequesterID == "" || in.ResourceID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, requester_id, resource_id required")
	}
	if in.TTL == 0 {
		in.TTL = 24 * time.Hour
	}
	now := common.NowUTC()
	item := &ApprovalRequest{
		ID: common.NewID(), OrgID: in.OrgID, RequesterID: in.RequesterID,
		ResourceType: in.ResourceType, ResourceID: in.ResourceID, Action: in.Action,
		Reason: in.Reason, Status: StatusPending, ExpiresAt: now.Add(in.TTL), CreatedAt: now,
	}
	s.mu.Lock()
	s.items[item.ID] = item
	s.mu.Unlock()
	cp := *item
	return &cp, nil
}

// Decide approves or denies a pending request.
func (s *Service) Decide(_ context.Context, id string, in DecideInput) (*ApprovalRequest, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.items[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "approval not found")
	}
	s.expireIfNeeded(item)
	if item.Status != StatusPending {
		return nil, common.New(common.ErrCodeConflict, "approval not pending")
	}
	now := common.NowUTC()
	item.DecidedAt = &now
	item.ApproverID = in.ApproverID
	if in.Deny {
		item.Status = StatusDenied
		if in.Reason != "" {
			item.Reason = in.Reason
		}
	} else {
		item.Status = StatusApproved
	}
	cp := *item
	return &cp, nil
}

// Get returns an approval by ID.
func (s *Service) Get(_ context.Context, id string) (*ApprovalRequest, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.items[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "approval not found")
	}
	s.expireIfNeeded(item)
	cp := *item
	return &cp, nil
}

// List returns approvals for an org.
func (s *Service) List(_ context.Context, orgID, status string) []ApprovalRequest {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]ApprovalRequest, 0)
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
func SchemaName() string { return "approvals" }

// TableName returns the primary table name.
func TableName() string { return "approval_requests" }

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
	StatusActive   = "active"
)

// EmergencyAccess is a break-glass access workflow.
type EmergencyAccess struct {
	ID          string     `json:"id"`
	OrgID       string     `json:"org_id"`
	RequesterID string     `json:"requester_id"`
	TargetOrgID string     `json:"target_org_id"`
	Reason      string     `json:"reason"`
	Status      string     `json:"status"`
	WaitPeriod  time.Duration `json:"wait_period"`
	ActiveUntil time.Time  `json:"active_until,omitempty"`
	ApprovedBy  string     `json:"approved_by,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

// CreateEmergencyInput holds break-glass request fields.
type CreateEmergencyInput struct {
	OrgID       string        `json:"org_id"`
	RequesterID string        `json:"requester_id"`
	TargetOrgID string        `json:"target_org_id"`
	Reason      string        `json:"reason"`
	WaitPeriod  time.Duration `json:"wait_period"`
	ActiveTTL   time.Duration `json:"active_ttl"`
}

// ActivateInput holds activation fields after wait period.
type ActivateInput struct {
	ApproverID string `json:"approver_id"`
	Deny       bool   `json:"deny"`
}

// Service implements emergency access workflow.
type Service struct {
	mu    sync.RWMutex
	items map[string]*EmergencyAccess
}

// NewService creates an emergency access domain service.
func NewService() *Service {
	return &Service{items: map[string]*EmergencyAccess{}}
}

func (s *Service) refreshStatus(item *EmergencyAccess) {
	now := common.NowUTC()
	if item.Status == StatusActive && !item.ActiveUntil.IsZero() && now.After(item.ActiveUntil) {
		item.Status = StatusExpired
	}
	if item.Status == StatusPending {
		eligible := item.CreatedAt.Add(item.WaitPeriod)
		if now.Before(eligible) {
			return
		}
	}
}

// Create opens a pending emergency access request.
func (s *Service) Create(_ context.Context, in CreateEmergencyInput) (*EmergencyAccess, error) {
	if in.OrgID == "" || in.RequesterID == "" || in.Reason == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, requester_id, reason required")
	}
	if in.WaitPeriod <= 0 {
		in.WaitPeriod = time.Hour
	}
	if in.ActiveTTL <= 0 {
		in.ActiveTTL = 8 * time.Hour
	}
	item := &EmergencyAccess{
		ID: common.NewID(), OrgID: in.OrgID, RequesterID: in.RequesterID,
		TargetOrgID: in.TargetOrgID, Reason: in.Reason, Status: StatusPending,
		WaitPeriod: in.WaitPeriod, CreatedAt: common.NowUTC(),
	}
	s.mu.Lock()
	s.items[item.ID] = item
	s.mu.Unlock()
	cp := *item
	return &cp, nil
}

// Activate approves or denies after the wait period.
func (s *Service) Activate(_ context.Context, id string, in ActivateInput, activeTTL time.Duration) (*EmergencyAccess, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.items[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "emergency access not found")
	}
	s.refreshStatus(item)
	if item.Status != StatusPending {
		return nil, common.New(common.ErrCodeConflict, "emergency access not pending")
	}
	if common.NowUTC().Before(item.CreatedAt.Add(item.WaitPeriod)) {
		return nil, common.New(common.ErrCodeConflict, "wait period not elapsed")
	}
	if in.Deny {
		item.Status = StatusDenied
	} else {
		if activeTTL <= 0 {
			activeTTL = 8 * time.Hour
		}
		item.Status = StatusActive
		item.ApprovedBy = in.ApproverID
		item.ActiveUntil = common.NowUTC().Add(activeTTL)
	}
	cp := *item
	return &cp, nil
}

// Get returns an emergency access record.
func (s *Service) Get(_ context.Context, id string) (*EmergencyAccess, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.items[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "emergency access not found")
	}
	s.refreshStatus(item)
	cp := *item
	return &cp, nil
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "emergency_access" }

// TableName returns the primary table name.
func TableName() string { return "sessions" }

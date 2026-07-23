package pam

import (
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Request statuses.
const (
	StatusPending  = "pending"
	StatusApproved = "approved"
	StatusDenied   = "denied"
	StatusExpired  = "expired"
	StatusRevoked  = "revoked"
)

// AccessRequest is a privileged access request.
type AccessRequest struct {
	ID            string        `json:"id"`
	OrgID         string        `json:"org_id"`
	RequesterID   string        `json:"requester_id"`
	ResourceType  string        `json:"resource_type"`
	ResourceID    string        `json:"resource_id"`
	Justification string        `json:"justification"`
	Status        string        `json:"status"`
	ApproverID    string        `json:"approver_id,omitempty"`
	TTL           time.Duration `json:"ttl"`
	ExpiresAt     *time.Time    `json:"expires_at,omitempty"`
	CreatedAt     time.Time     `json:"created_at"`
	DecidedAt     *time.Time    `json:"decided_at,omitempty"`
}

// EmergencyAccess is a break-glass grant.
type EmergencyAccess struct {
	ID         string    `json:"id"`
	OrgID      string    `json:"org_id"`
	ActorID    string    `json:"actor_id"`
	ResourceID string    `json:"resource_id"`
	Reason     string    `json:"reason"`
	Status     string    `json:"status"`
	ExpiresAt  time.Time `json:"expires_at"`
	CreatedAt  time.Time `json:"created_at"`
}

// Delegation grants another user temporary rights.
type Delegation struct {
	ID          string    `json:"id"`
	OrgID       string    `json:"org_id"`
	FromUserID  string    `json:"from_user_id"`
	ToUserID    string    `json:"to_user_id"`
	Permissions []string  `json:"permissions"`
	Status      string    `json:"status"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// Workflow manages PAM state machines.
type Workflow struct {
	mu          sync.Mutex
	requests    map[string]*AccessRequest
	emergencies map[string]*EmergencyAccess
	delegations map[string]*Delegation
}

// NewWorkflow creates a PAM workflow engine.
func NewWorkflow() *Workflow {
	return &Workflow{
		requests:    map[string]*AccessRequest{},
		emergencies: map[string]*EmergencyAccess{},
		delegations: map[string]*Delegation{},
	}
}

// CreateAccessRequest opens a pending request.
func (w *Workflow) CreateAccessRequest(orgID, requesterID, resourceType, resourceID, justification string, ttl time.Duration) (*AccessRequest, error) {
	if orgID == "" || requesterID == "" || resourceID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, requester_id, resource_id required")
	}
	if ttl <= 0 {
		ttl = 8 * time.Hour
	}
	req := &AccessRequest{
		ID: common.NewID(), OrgID: orgID, RequesterID: requesterID,
		ResourceType: resourceType, ResourceID: resourceID, Justification: justification,
		Status: StatusPending, TTL: ttl, CreatedAt: common.NowUTC(),
	}
	w.mu.Lock()
	w.requests[req.ID] = req
	w.mu.Unlock()
	cp := *req
	return &cp, nil
}

// DecideAccessRequest approves or denies.
func (w *Workflow) DecideAccessRequest(id, approverID, decision string) (*AccessRequest, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	req, ok := w.requests[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "request not found")
	}
	if req.Status != StatusPending {
		return nil, common.New(common.ErrCodeConflict, "request not pending")
	}
	now := common.NowUTC()
	req.ApproverID = approverID
	req.DecidedAt = &now
	switch decision {
	case StatusApproved:
		req.Status = StatusApproved
		exp := now.Add(req.TTL)
		req.ExpiresAt = &exp
	case StatusDenied:
		req.Status = StatusDenied
	default:
		return nil, common.New(common.ErrCodeInvalidArgument, "decision must be approved or denied")
	}
	cp := *req
	return &cp, nil
}

// CreateEmergency creates a break-glass grant.
func (w *Workflow) CreateEmergency(orgID, actorID, resourceID, reason string, ttl time.Duration) (*EmergencyAccess, error) {
	if orgID == "" || actorID == "" || resourceID == "" || reason == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, actor_id, resource_id, reason required")
	}
	if ttl <= 0 {
		ttl = time.Hour
	}
	now := common.NowUTC()
	ea := &EmergencyAccess{
		ID: common.NewID(), OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
		Reason: reason, Status: StatusApproved, ExpiresAt: now.Add(ttl), CreatedAt: now,
	}
	w.mu.Lock()
	w.emergencies[ea.ID] = ea
	w.mu.Unlock()
	cp := *ea
	return &cp, nil
}

// CreateDelegation creates a time-limited delegation.
func (w *Workflow) CreateDelegation(orgID, from, to string, perms []string, ttl time.Duration) (*Delegation, error) {
	if orgID == "" || from == "" || to == "" || len(perms) == 0 {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, from, to, permissions required")
	}
	if ttl <= 0 {
		ttl = 24 * time.Hour
	}
	now := common.NowUTC()
	d := &Delegation{
		ID: common.NewID(), OrgID: orgID, FromUserID: from, ToUserID: to,
		Permissions: perms, Status: StatusApproved, ExpiresAt: now.Add(ttl), CreatedAt: now,
	}
	w.mu.Lock()
	w.delegations[d.ID] = d
	w.mu.Unlock()
	cp := *d
	return &cp, nil
}

// ExpireStale marks expired approved requests/delegations/emergencies.
func (w *Workflow) ExpireStale(now time.Time) int {
	w.mu.Lock()
	defer w.mu.Unlock()
	n := 0
	for _, r := range w.requests {
		if r.Status == StatusApproved && r.ExpiresAt != nil && now.After(*r.ExpiresAt) {
			r.Status = StatusExpired
			n++
		}
	}
	for _, e := range w.emergencies {
		if e.Status == StatusApproved && now.After(e.ExpiresAt) {
			e.Status = StatusExpired
			n++
		}
	}
	for _, d := range w.delegations {
		if d.Status == StatusApproved && now.After(d.ExpiresAt) {
			d.Status = StatusExpired
			n++
		}
	}
	return n
}

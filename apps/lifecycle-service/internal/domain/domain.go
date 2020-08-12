package domain

import (
	"context"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

const (
	StateInvited     = "invited"
	StateActive      = "active"
	StateSuspended   = "suspended"
	StateOffboarded  = "offboarded"
)

// LifecycleState tracks user lifecycle in an org.
type LifecycleState struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	UserID    string    `json:"user_id"`
	State     string    `json:"state"`
	Reason    string    `json:"reason,omitempty"`
	ChangedAt time.Time `json:"changed_at"`
}

// TransitionInput holds lifecycle transition fields.
type TransitionInput struct {
	OrgID  string `json:"org_id"`
	UserID string `json:"user_id"`
	To     string `json:"to"`
	Reason string `json:"reason"`
}

// Service manages user lifecycle states.
type Service struct {
	mu     sync.RWMutex
	states map[string]*LifecycleState // key: orgID/userID
}

// NewService creates a lifecycle domain service.
func NewService() *Service {
	return &Service{states: map[string]*LifecycleState{}}
}

func key(orgID, userID string) string { return orgID + "/" + userID }

var allowedTransitions = map[string][]string{
	StateInvited:    {StateActive, StateOffboarded},
	StateActive:     {StateSuspended, StateOffboarded},
	StateSuspended:  {StateActive, StateOffboarded},
	StateOffboarded: {},
}

// Transition moves a user to a new lifecycle state.
func (s *Service) Transition(_ context.Context, in TransitionInput) (*LifecycleState, error) {
	if in.OrgID == "" || in.UserID == "" || in.To == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, user_id, to required")
	}
	if err := common.ValidateOneOf("to", in.To, StateInvited, StateActive, StateSuspended, StateOffboarded); err != nil {
		return nil, err
	}
	k := key(in.OrgID, in.UserID)
	s.mu.Lock()
	defer s.mu.Unlock()
	cur, ok := s.states[k]
	from := ""
	if ok {
		from = cur.State
	}
	if !canTransition(from, in.To) {
		return nil, common.New(common.ErrCodeConflict, "invalid lifecycle transition")
	}
	st := &LifecycleState{
		ID: common.NewID(), OrgID: in.OrgID, UserID: in.UserID,
		State: in.To, Reason: in.Reason, ChangedAt: common.NowUTC(),
	}
	s.states[k] = st
	cp := *st
	return &cp, nil
}

func canTransition(from, to string) bool {
	for _, t := range allowedTransitions[from] {
		if t == to {
			return true
		}
	}
	return from == "" && (to == StateInvited || to == StateActive)
}

// Invite marks a user as invited (initial state).
func (s *Service) Invite(ctx context.Context, orgID, userID string) (*LifecycleState, error) {
	return s.Transition(ctx, TransitionInput{OrgID: orgID, UserID: userID, To: StateInvited, Reason: "invited"})
}

// Get returns current lifecycle state.
func (s *Service) Get(_ context.Context, orgID, userID string) (*LifecycleState, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	st, ok := s.states[key(orgID, userID)]
	if !ok {
		return &LifecycleState{OrgID: orgID, UserID: userID, State: StateInvited, ChangedAt: common.NowUTC()}, nil
	}
	cp := *st
	return &cp, nil
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "lifecycle" }

// TableName returns the primary table name.
func TableName() string { return "user_states" }

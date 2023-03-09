package domain

import (
	"context"
	"sync"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// FeatureFlag is a platform-wide toggle.
type FeatureFlag struct {
	Name        string `json:"name"`
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
}

// OrgState tracks suspension for an organization.
type OrgState struct {
	OrgID     string `json:"org_id"`
	Suspended bool   `json:"suspended"`
	Reason    string `json:"reason,omitempty"`
}

// Service provides in-memory platform admin operations.
type Service struct {
	mu      sync.RWMutex
	flags   map[string]*FeatureFlag
	orgStates map[string]*OrgState
}

// NewService creates an admin domain service with default flags.
func NewService() *Service {
	s := &Service{
		flags: map[string]*FeatureFlag{
			"advanced_risk": {Name: "advanced_risk", Enabled: true, Description: "Adaptive risk scoring"},
			"scim_provisioning": {Name: "scim_provisioning", Enabled: false, Description: "SCIM user sync"},
			"break_glass": {Name: "break_glass", Enabled: true, Description: "Emergency access workflows"},
			"audit_export": {Name: "audit_export", Enabled: true, Description: "Audit log export API"},
		},
		orgStates: map[string]*OrgState{},
	}
	return s
}

// ListFeatureFlags returns all platform feature flags.
func (s *Service) ListFeatureFlags(_ context.Context) []FeatureFlag {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]FeatureFlag, 0, len(s.flags))
	for _, f := range s.flags {
		out = append(out, *f)
	}
	return out
}

// SetFeatureFlag updates a feature flag.
func (s *Service) SetFeatureFlag(_ context.Context, name string, enabled bool) (*FeatureFlag, error) {
	if name == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "name required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	f, ok := s.flags[name]
	if !ok {
		f = &FeatureFlag{Name: name}
		s.flags[name] = f
	}
	f.Enabled = enabled
	cp := *f
	return &cp, nil
}

// SuspendOrg marks an organization suspended.
func (s *Service) SuspendOrg(_ context.Context, orgID, reason string) (*OrgState, error) {
	if orgID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	st := &OrgState{OrgID: orgID, Suspended: true, Reason: reason}
	s.orgStates[orgID] = st
	cp := *st
	return &cp, nil
}

// GetOrgState returns org suspension state.
func (s *Service) GetOrgState(_ context.Context, orgID string) (*OrgState, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	st, ok := s.orgStates[orgID]
	if !ok {
		return &OrgState{OrgID: orgID, Suspended: false}, nil
	}
	cp := *st
	return &cp, nil
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "admin" }

// TableName returns the primary table name.
func TableName() string { return "feature_flags" }

package domain

import (
	"context"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/rotation"
)

// Policy re-exports rotation.Policy.
type Policy = rotation.Policy

// Job re-exports rotation.Job.
type Job = rotation.Job

// PolicyInput holds policy upsert fields.
type PolicyInput struct {
	ID               string        `json:"id"`
	OrgID            string        `json:"org_id"`
	Name             string        `json:"name"`
	TargetType       string        `json:"target_type"`
	TargetID         string        `json:"target_id"`
	Mode             string        `json:"mode"`
	Interval         time.Duration `json:"interval"`
	ApprovalRequired bool          `json:"approval_required"`
	Enabled          bool          `json:"enabled"`
}

// RunInput holds rotation run fields.
type RunInput struct {
	PolicyID string `json:"policy_id"`
	Mode     string `json:"mode"`
}

// Service wraps the rotation engine.
type Service struct {
	engine *rotation.Engine
}

// NewService creates a rotation domain service.
func NewService(engine *rotation.Engine) *Service {
	return &Service{engine: engine}
}

// NewEngine creates a rotation engine with a static rotator.
func NewEngine() *rotation.Engine {
	return rotation.NewEngine(rotation.StaticRotator{})
}

// UpsertPolicy creates or updates a rotation policy.
func (s *Service) UpsertPolicy(_ context.Context, in PolicyInput) (*Policy, error) {
	return s.engine.UpsertPolicy(Policy{
		ID: in.ID, OrgID: in.OrgID, Name: in.Name, TargetType: in.TargetType,
		TargetID: in.TargetID, Mode: in.Mode, Interval: in.Interval,
		ApprovalRequired: in.ApprovalRequired, Enabled: in.Enabled,
	})
}

// GetPolicy returns a policy by ID.
func (s *Service) GetPolicy(_ context.Context, id string) (*Policy, error) {
	return s.engine.GetPolicy(id)
}

// ListPolicies returns policies for an org.
func (s *Service) ListPolicies(_ context.Context, orgID string) []Policy {
	return s.engine.ListPolicies(orgID)
}

// DeletePolicy removes a policy.
func (s *Service) DeletePolicy(_ context.Context, id string) error {
	return s.engine.DeletePolicy(id)
}

// RunPolicy executes rotation for a policy.
func (s *Service) RunPolicy(_ context.Context, in RunInput) (*Job, error) {
	if in.PolicyID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "policy_id required")
	}
	return s.engine.RunPolicy(in.PolicyID, in.Mode)
}

// Jobs returns rotation job history.
func (s *Service) Jobs(_ context.Context) []Job {
	return s.engine.Jobs()
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "rotation" }

// TableName returns the primary table name.
func TableName() string { return "policies" }

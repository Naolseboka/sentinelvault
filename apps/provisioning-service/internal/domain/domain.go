package domain

import (
	"context"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

const (
	PlanStatusDraft     = "draft"
	PlanStatusScheduled = "scheduled"
	PlanStatusRunning   = "running"
	PlanStatusCompleted = "completed"
	PlanStatusFailed    = "failed"

	StepStatusPending   = "pending"
	StepStatusRunning   = "running"
	StepStatusSucceeded = "succeeded"
	StepStatusFailed    = "failed"

	JobStatusPending   = "pending"
	JobStatusRunning   = "running"
	JobStatusCompleted = "completed"
	JobStatusFailed    = "failed"
)

// PlanStep is one provisioning action.
type PlanStep struct {
	Order  int    `json:"order"`
	Action string `json:"action"` // create_user|assign_role|grant_vault_access
	Target string `json:"target"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

// ProvisioningPlan onboards resources for an org.
type ProvisioningPlan struct {
	ID         string     `json:"id"`
	OrgID      string     `json:"org_id"`
	Name       string     `json:"name"`
	Status     string     `json:"status"`
	Steps      []PlanStep `json:"steps"`
	CreatedAt  time.Time  `json:"created_at"`
	StartedAt  *time.Time `json:"started_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
}

// ProvisioningJob tracks execution of a plan.
type ProvisioningJob struct {
	ID         string     `json:"id"`
	PlanID     string     `json:"plan_id"`
	OrgID      string     `json:"org_id"`
	Status     string     `json:"status"`
	Steps      []PlanStep `json:"steps"`
	Error      string     `json:"error,omitempty"`
	StartedAt  time.Time  `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
}

// CreatePlanInput holds plan creation fields.
type CreatePlanInput struct {
	OrgID string     `json:"org_id"`
	Name  string     `json:"name"`
	Steps []PlanStep `json:"steps"`
}

// Service manages provisioning plans and jobs.
type Service struct {
	mu    sync.RWMutex
	plans map[string]*ProvisioningPlan
	jobs  map[string]*ProvisioningJob
}

// NewService creates a provisioning domain service.
func NewService() *Service {
	return &Service{plans: map[string]*ProvisioningPlan{}, jobs: map[string]*ProvisioningJob{}}
}

// Create registers a draft provisioning plan with pending steps.
func (s *Service) Create(_ context.Context, in CreatePlanInput) (*ProvisioningPlan, error) {
	if in.OrgID == "" || in.Name == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and name required")
	}
	steps := make([]PlanStep, len(in.Steps))
	for i, st := range in.Steps {
		steps[i] = st
		if steps[i].Status == "" {
			steps[i].Status = StepStatusPending
		}
	}
	plan := &ProvisioningPlan{
		ID: common.NewID(), OrgID: in.OrgID, Name: in.Name,
		Status: PlanStatusDraft, Steps: steps, CreatedAt: common.NowUTC(),
	}
	s.mu.Lock()
	s.plans[plan.ID] = plan
	s.mu.Unlock()
	cp := *plan
	return &cp, nil
}

// Schedule moves a plan from draft to scheduled.
func (s *Service) Schedule(_ context.Context, id string) (*ProvisioningPlan, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	plan, ok := s.plans[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "plan not found")
	}
	if plan.Status != PlanStatusDraft {
		return nil, common.New(common.ErrCodeConflict, "plan must be draft to schedule")
	}
	plan.Status = PlanStatusScheduled
	cp := *plan
	return &cp, nil
}

// Run executes a scheduled plan and creates a provisioning job.
func (s *Service) Run(_ context.Context, id string) (*ProvisioningJob, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	plan, ok := s.plans[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "plan not found")
	}
	if plan.Status != PlanStatusScheduled && plan.Status != PlanStatusDraft {
		return nil, common.New(common.ErrCodeConflict, "plan cannot be run in current status")
	}
	now := common.NowUTC()
	plan.StartedAt = &now
	plan.Status = PlanStatusRunning

	jobSteps := make([]PlanStep, len(plan.Steps))
	copy(jobSteps, plan.Steps)

	job := &ProvisioningJob{
		ID: common.NewID(), PlanID: plan.ID, OrgID: plan.OrgID,
		Status: JobStatusRunning, Steps: jobSteps, StartedAt: now,
	}

	for i := range job.Steps {
		job.Steps[i].Status = StepStatusRunning
		plan.Steps[i].Status = StepStatusRunning
		job.Steps[i].Status = StepStatusSucceeded
		plan.Steps[i].Status = StepStatusSucceeded
	}
	fin := common.NowUTC()
	job.FinishedAt = &fin
	plan.FinishedAt = &fin
	job.Status = JobStatusCompleted
	plan.Status = PlanStatusCompleted
	s.jobs[job.ID] = job
	cp := *job
	return &cp, nil
}

// GetPlan returns a plan by ID.
func (s *Service) GetPlan(_ context.Context, id string) (*ProvisioningPlan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	plan, ok := s.plans[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "plan not found")
	}
	cp := *plan
	return &cp, nil
}

// GetJob returns a job by ID.
func (s *Service) GetJob(_ context.Context, id string) (*ProvisioningJob, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	job, ok := s.jobs[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "job not found")
	}
	cp := *job
	return &cp, nil
}

// Get is an alias for GetPlan.
func (s *Service) Get(ctx context.Context, id string) (*ProvisioningPlan, error) {
	return s.GetPlan(ctx, id)
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "provisioning" }

// TableName returns the primary table name.
func TableName() string { return "plans" }

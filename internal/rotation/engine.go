package rotation

import (
	"fmt"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Mode of rotation.
const (
	ModeAutomatic = "automatic"
	ModeManual    = "manual"
	ModeScheduled = "scheduled"
	ModeApproval  = "approval"
	ModeEmergency = "emergency"
)

// Policy defines how a secret/credential is rotated.
type Policy struct {
	ID               string        `json:"id"`
	OrgID            string        `json:"org_id"`
	Name             string        `json:"name"`
	TargetType       string        `json:"target_type"` // secret|password|ssh|certificate|db
	TargetID         string        `json:"target_id"`
	Mode             string        `json:"mode"`
	Interval         time.Duration `json:"interval"`
	NextRunAt        time.Time     `json:"next_run_at"`
	ApprovalRequired bool          `json:"approval_required"`
	Enabled          bool          `json:"enabled"`
	CreatedAt        time.Time     `json:"created_at"`
}

// Job is a rotation execution attempt.
type Job struct {
	ID         string     `json:"id"`
	PolicyID   string     `json:"policy_id"`
	OrgID      string     `json:"org_id"`
	Status     string     `json:"status"` // pending|running|succeeded|failed|rolled_back
	Mode       string     `json:"mode"`
	Error      string     `json:"error,omitempty"`
	StartedAt  time.Time  `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	NewVersion string     `json:"new_version,omitempty"`
}

// Rotator generates new secret material.
type Rotator interface {
	Rotate(orgID, targetType, targetID string) (newVersion string, err error)
}

// Engine schedules and executes rotation policies.
type Engine struct {
	policies map[string]*Policy
	jobs     []Job
	rotator  Rotator
}

// NewEngine creates a rotation engine.
func NewEngine(r Rotator) *Engine {
	return &Engine{
		policies: map[string]*Policy{},
		jobs:     make([]Job, 0),
		rotator:  r,
	}
}

// UpsertPolicy creates or updates a policy.
func (e *Engine) UpsertPolicy(p Policy) (*Policy, error) {
	if p.OrgID == "" || p.Name == "" || p.TargetID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, name, target_id required")
	}
	switch p.Mode {
	case ModeAutomatic, ModeManual, ModeScheduled, ModeApproval, ModeEmergency:
	default:
		return nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("invalid mode %q", p.Mode))
	}
	if p.ID == "" {
		p.ID = common.NewID()
		p.CreatedAt = common.NowUTC()
	}
	if p.Mode == ModeScheduled && p.Interval > 0 && p.NextRunAt.IsZero() {
		p.NextRunAt = common.NowUTC().Add(p.Interval)
	}
	cp := p
	e.policies[p.ID] = &cp
	return &cp, nil
}

// RunPolicy executes a rotation for a policy.
func (e *Engine) RunPolicy(policyID string, mode string) (*Job, error) {
	p, ok := e.policies[policyID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "policy not found")
	}
	if !p.Enabled && mode != ModeEmergency {
		return nil, common.New(common.ErrCodeConflict, "policy disabled")
	}
	if mode == "" {
		mode = p.Mode
	}
	if (mode == ModeApproval || p.ApprovalRequired) && mode != ModeEmergency {
		job := Job{
			ID: common.NewID(), PolicyID: p.ID, OrgID: p.OrgID,
			Status: "pending", Mode: ModeApproval, StartedAt: common.NowUTC(),
		}
		e.jobs = append(e.jobs, job)
		return &job, nil
	}
	job := Job{
		ID: common.NewID(), PolicyID: p.ID, OrgID: p.OrgID,
		Status: "running", Mode: mode, StartedAt: common.NowUTC(),
	}
	ver, err := e.rotator.Rotate(p.OrgID, p.TargetType, p.TargetID)
	now := common.NowUTC()
	job.FinishedAt = &now
	if err != nil {
		job.Status = "failed"
		job.Error = err.Error()
		e.jobs = append(e.jobs, job)
		return &job, err
	}
	job.Status = "succeeded"
	job.NewVersion = ver
	if p.Interval > 0 {
		p.NextRunAt = now.Add(p.Interval)
	}
	e.jobs = append(e.jobs, job)
	return &job, nil
}

// DuePolicies returns scheduled policies ready to run.
func (e *Engine) DuePolicies(now time.Time) []*Policy {
	out := make([]*Policy, 0)
	for _, p := range e.policies {
		if p.Enabled && p.Mode == ModeScheduled && !p.NextRunAt.IsZero() && !p.NextRunAt.After(now) {
			cp := *p
			out = append(out, &cp)
		}
	}
	return out
}

// Jobs returns job history.
func (e *Engine) Jobs() []Job {
	out := make([]Job, len(e.jobs))
	copy(out, e.jobs)
	return out
}

// GetPolicy returns a policy by ID.
func (e *Engine) GetPolicy(id string) (*Policy, error) {
	p, ok := e.policies[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "policy not found")
	}
	cp := *p
	return &cp, nil
}

// ListPolicies returns policies optionally filtered by org.
func (e *Engine) ListPolicies(orgID string) []Policy {
	out := make([]Policy, 0, len(e.policies))
	for _, p := range e.policies {
		if orgID == "" || p.OrgID == orgID {
			out = append(out, *p)
		}
	}
	return out
}

// DeletePolicy removes a policy.
func (e *Engine) DeletePolicy(id string) error {
	if _, ok := e.policies[id]; !ok {
		return common.New(common.ErrCodeNotFound, "policy not found")
	}
	delete(e.policies, id)
	return nil
}

// StaticRotator is a test/dev rotator that returns a new version ID.
type StaticRotator struct{}

func (StaticRotator) Rotate(_, _, _ string) (string, error) {
	return common.NewID(), nil
}

package saga

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

const (
	StatusPending      = "pending"
	StatusRunning      = "running"
	StatusCompensating = "compensating"
	StatusCompleted    = "completed"
	StatusFailed       = "failed"
)

type StepFunc func(ctx context.Context, state map[string]any) error
type CompensateFunc func(ctx context.Context, state map[string]any) error

type Step struct {
	Name       string
	Execute    StepFunc
	Compensate CompensateFunc
}

type Definition struct {
	Name  string
	Steps []Step
}

type Instance struct {
	ID          string
	Name        string
	Status      string
	CurrentStep int
	State       map[string]any
	Error       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Engine struct {
	mu   sync.Mutex
	defs map[string]Definition
	inst map[string]*Instance
}

func NewEngine() *Engine {
	return &Engine{defs: map[string]Definition{}, inst: map[string]*Instance{}}
}

func (e *Engine) Register(def Definition) error {
	if def.Name == "" || len(def.Steps) == 0 {
		return common.New(common.ErrCodeInvalidArgument, "name and steps required")
	}
	e.mu.Lock()
	e.defs[def.Name] = def
	e.mu.Unlock()
	return nil
}

func (e *Engine) Start(ctx context.Context, name string, state map[string]any) (*Instance, error) {
	e.mu.Lock()
	def, ok := e.defs[name]
	if !ok {
		e.mu.Unlock()
		return nil, common.New(common.ErrCodeNotFound, "saga not registered")
	}
	if state == nil {
		state = map[string]any{}
	}
	inst := &Instance{ID: common.NewID(), Name: name, Status: StatusRunning, State: state, CreatedAt: common.NowUTC(), UpdatedAt: common.NowUTC()}
	e.inst[inst.ID] = inst
	e.mu.Unlock()
	if err := e.run(ctx, inst.ID, def); err != nil {
		return e.Get(inst.ID)
	}
	return e.Get(inst.ID)
}

func (e *Engine) Get(id string) (*Instance, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	inst, ok := e.inst[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "saga not found")
	}
	cp := *inst
	cp.State = copyState(inst.State)
	return &cp, nil
}

func (e *Engine) run(ctx context.Context, id string, def Definition) error {
	for i, step := range def.Steps {
		e.mu.Lock()
		inst := e.inst[id]
		inst.CurrentStep = i
		inst.Status = StatusRunning
		inst.UpdatedAt = common.NowUTC()
		state := copyState(inst.State)
		e.mu.Unlock()
		if err := step.Execute(ctx, state); err != nil {
			e.mu.Lock()
			inst.Error = err.Error()
			inst.Status = StatusCompensating
			inst.State = state
			e.mu.Unlock()
			_ = e.compensate(ctx, id, def, i-1)
			e.mu.Lock()
			inst.Status = StatusFailed
			inst.UpdatedAt = common.NowUTC()
			e.mu.Unlock()
			return err
		}
		e.mu.Lock()
		inst.State = state
		inst.UpdatedAt = common.NowUTC()
		e.mu.Unlock()
	}
	e.mu.Lock()
	inst := e.inst[id]
	inst.Status = StatusCompleted
	inst.UpdatedAt = common.NowUTC()
	e.mu.Unlock()
	return nil
}

func (e *Engine) compensate(ctx context.Context, id string, def Definition, from int) error {
	for i := from; i >= 0; i-- {
		step := def.Steps[i]
		if step.Compensate == nil {
			continue
		}
		e.mu.Lock()
		state := copyState(e.inst[id].State)
		e.mu.Unlock()
		if err := step.Compensate(ctx, state); err != nil {
			return fmt.Errorf("compensate %s: %w", step.Name, err)
		}
		e.mu.Lock()
		e.inst[id].State = state
		e.mu.Unlock()
	}
	return nil
}

func copyState(in map[string]any) map[string]any {
	out := make(map[string]any, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

// Builtins registers standard platform sagas.
func Builtins(e *Engine) {
	_ = e.Register(Definition{Name: "UserOffboarding", Steps: []Step{
		{Name: "revoke_sessions", Execute: func(ctx context.Context, s map[string]any) error { s["sessions_revoked"] = true; return nil }, Compensate: func(ctx context.Context, s map[string]any) error { s["sessions_revoked"] = false; return nil }},
		{Name: "disable_user", Execute: func(ctx context.Context, s map[string]any) error { s["user_disabled"] = true; return nil }, Compensate: func(ctx context.Context, s map[string]any) error { s["user_disabled"] = false; return nil }},
		{Name: "reassign_owned", Execute: func(ctx context.Context, s map[string]any) error { s["reassigned"] = true; return nil }},
	}})
	_ = e.Register(Definition{Name: "SecretRotation", Steps: []Step{
		{Name: "generate", Execute: func(ctx context.Context, s map[string]any) error { s["new_version"] = "v2"; return nil }},
		{Name: "propagate", Execute: func(ctx context.Context, s map[string]any) error { s["propagated"] = true; return nil }, Compensate: func(ctx context.Context, s map[string]any) error { s["propagated"] = false; return nil }},
		{Name: "retire_old", Execute: func(ctx context.Context, s map[string]any) error { s["retired"] = true; return nil }},
	}})
	_ = e.Register(Definition{Name: "OrgProvision", Steps: []Step{
		{Name: "create_org", Execute: func(ctx context.Context, s map[string]any) error { s["org"] = true; return nil }, Compensate: func(ctx context.Context, s map[string]any) error { s["org"] = false; return nil }},
		{Name: "seed_rbac", Execute: func(ctx context.Context, s map[string]any) error { s["rbac"] = true; return nil }},
		{Name: "issue_license", Execute: func(ctx context.Context, s map[string]any) error { s["license"] = true; return nil }},
	}})
	_ = e.Register(Definition{Name: "EmergencyAccess", Steps: []Step{
		{Name: "validate_reason", Execute: func(ctx context.Context, s map[string]any) error {
			if s["reason"] == nil || s["reason"] == "" {
				return common.New(common.ErrCodeInvalidArgument, "reason required")
			}
			return nil
		}},
		{Name: "grant", Execute: func(ctx context.Context, s map[string]any) error { s["granted"] = true; return nil }, Compensate: func(ctx context.Context, s map[string]any) error { s["granted"] = false; return nil }},
		{Name: "audit", Execute: func(ctx context.Context, s map[string]any) error { s["audited"] = true; return nil }},
	}})
}

// maintenance note 16

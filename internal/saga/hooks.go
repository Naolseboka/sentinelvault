package saga

import "context"

// StepHandler executes one saga step with typed compensation.
type StepHandler interface {
    Name() string
    Execute(ctx context.Context, state map[string]any) error
    Compensate(ctx context.Context, state map[string]any) error
}

// noopHandler skips compensation.
type noopHandler struct{ name string }

func (n noopHandler) Name() string { return n.name }
func (n noopHandler) Execute(ctx context.Context, s map[string]any) error { return nil }
func (n noopHandler) Compensate(ctx context.Context, s map[string]any) error { return nil }

// FromHandler converts StepHandler to Step.
func FromHandler(h StepHandler) Step {
    return Step{Name: h.Name(), Execute: h.Execute, Compensate: h.Compensate}
}

// RegisterHandlers registers handlers as a saga definition.
func RegisterHandlers(e *Engine, name string, handlers ...StepHandler) error {
    steps := make([]Step, len(handlers))
    for i, h := range handlers {
        steps[i] = FromHandler(h)
    }
    return e.Register(Definition{Name: name, Steps: steps})
}
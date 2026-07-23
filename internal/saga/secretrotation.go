package saga

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type SecretRotationStep1 struct{ fail bool }

func (s SecretRotationStep1) Name() string { return "secretrotation_step_1" }
func (s SecretRotationStep1) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["secretrotation_step_1"] = true
	return nil
}
func (s SecretRotationStep1) Compensate(ctx context.Context, st map[string]any) error {
	st["secretrotation_step_1"] = false
	return nil
}

type SecretRotationStep2 struct{ fail bool }

func (s SecretRotationStep2) Name() string { return "secretrotation_step_2" }
func (s SecretRotationStep2) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["secretrotation_step_2"] = true
	return nil
}
func (s SecretRotationStep2) Compensate(ctx context.Context, st map[string]any) error {
	st["secretrotation_step_2"] = false
	return nil
}

type SecretRotationStep3 struct{ fail bool }

func (s SecretRotationStep3) Name() string { return "secretrotation_step_3" }
func (s SecretRotationStep3) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["secretrotation_step_3"] = true
	return nil
}
func (s SecretRotationStep3) Compensate(ctx context.Context, st map[string]any) error {
	st["secretrotation_step_3"] = false
	return nil
}

type SecretRotationStep4 struct{ fail bool }

func (s SecretRotationStep4) Name() string { return "secretrotation_step_4" }
func (s SecretRotationStep4) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["secretrotation_step_4"] = true
	return nil
}
func (s SecretRotationStep4) Compensate(ctx context.Context, st map[string]any) error {
	st["secretrotation_step_4"] = false
	return nil
}

// RegisterSecretRotation wires SecretRotation saga with compensation hooks.
func RegisterSecretRotation(e *Engine) error {
	return RegisterHandlers(e, "SecretRotation", &SecretRotationStep1{}, &SecretRotationStep2{}, &SecretRotationStep3{}, &SecretRotationStep4{})
}

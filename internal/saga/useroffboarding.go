package saga

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/common"
)


type UserOffboardingStep1 struct{ fail bool }

func (s UserOffboardingStep1) Name() string { return "useroffboarding_step_1" }
func (s UserOffboardingStep1) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["useroffboarding_step_1"] = true
	return nil
}
func (s UserOffboardingStep1) Compensate(ctx context.Context, st map[string]any) error {
	st["useroffboarding_step_1"] = false
	return nil
}

type UserOffboardingStep2 struct{ fail bool }

func (s UserOffboardingStep2) Name() string { return "useroffboarding_step_2" }
func (s UserOffboardingStep2) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["useroffboarding_step_2"] = true
	return nil
}
func (s UserOffboardingStep2) Compensate(ctx context.Context, st map[string]any) error {
	st["useroffboarding_step_2"] = false
	return nil
}

type UserOffboardingStep3 struct{ fail bool }

func (s UserOffboardingStep3) Name() string { return "useroffboarding_step_3" }
func (s UserOffboardingStep3) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["useroffboarding_step_3"] = true
	return nil
}
func (s UserOffboardingStep3) Compensate(ctx context.Context, st map[string]any) error {
	st["useroffboarding_step_3"] = false
	return nil
}

type UserOffboardingStep4 struct{ fail bool }

func (s UserOffboardingStep4) Name() string { return "useroffboarding_step_4" }
func (s UserOffboardingStep4) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["useroffboarding_step_4"] = true
	return nil
}
func (s UserOffboardingStep4) Compensate(ctx context.Context, st map[string]any) error {
	st["useroffboarding_step_4"] = false
	return nil
}


// RegisterUserOffboarding wires UserOffboarding saga with compensation hooks.
func RegisterUserOffboarding(e *Engine) error {
	return RegisterHandlers(e, "UserOffboarding", &UserOffboardingStep1{}, &UserOffboardingStep2{}, &UserOffboardingStep3{}, &UserOffboardingStep4{})
}

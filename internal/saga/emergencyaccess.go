package saga

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type EmergencyAccessStep1 struct{ fail bool }

func (s EmergencyAccessStep1) Name() string { return "emergencyaccess_step_1" }
func (s EmergencyAccessStep1) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["emergencyaccess_step_1"] = true
	return nil
}
func (s EmergencyAccessStep1) Compensate(ctx context.Context, st map[string]any) error {
	st["emergencyaccess_step_1"] = false
	return nil
}

type EmergencyAccessStep2 struct{ fail bool }

func (s EmergencyAccessStep2) Name() string { return "emergencyaccess_step_2" }
func (s EmergencyAccessStep2) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["emergencyaccess_step_2"] = true
	return nil
}
func (s EmergencyAccessStep2) Compensate(ctx context.Context, st map[string]any) error {
	st["emergencyaccess_step_2"] = false
	return nil
}

type EmergencyAccessStep3 struct{ fail bool }

func (s EmergencyAccessStep3) Name() string { return "emergencyaccess_step_3" }
func (s EmergencyAccessStep3) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["emergencyaccess_step_3"] = true
	return nil
}
func (s EmergencyAccessStep3) Compensate(ctx context.Context, st map[string]any) error {
	st["emergencyaccess_step_3"] = false
	return nil
}

type EmergencyAccessStep4 struct{ fail bool }

func (s EmergencyAccessStep4) Name() string { return "emergencyaccess_step_4" }
func (s EmergencyAccessStep4) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["emergencyaccess_step_4"] = true
	return nil
}
func (s EmergencyAccessStep4) Compensate(ctx context.Context, st map[string]any) error {
	st["emergencyaccess_step_4"] = false
	return nil
}

// RegisterEmergencyAccess wires EmergencyAccess saga with compensation hooks.
func RegisterEmergencyAccess(e *Engine) error {
	return RegisterHandlers(e, "EmergencyAccess", &EmergencyAccessStep1{}, &EmergencyAccessStep2{}, &EmergencyAccessStep3{}, &EmergencyAccessStep4{})
}

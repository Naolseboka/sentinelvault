package saga

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type OrgProvisionStep1 struct{ fail bool }

func (s OrgProvisionStep1) Name() string { return "orgprovision_step_1" }
func (s OrgProvisionStep1) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["orgprovision_step_1"] = true
	return nil
}
func (s OrgProvisionStep1) Compensate(ctx context.Context, st map[string]any) error {
	st["orgprovision_step_1"] = false
	return nil
}

type OrgProvisionStep2 struct{ fail bool }

func (s OrgProvisionStep2) Name() string { return "orgprovision_step_2" }
func (s OrgProvisionStep2) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["orgprovision_step_2"] = true
	return nil
}
func (s OrgProvisionStep2) Compensate(ctx context.Context, st map[string]any) error {
	st["orgprovision_step_2"] = false
	return nil
}

type OrgProvisionStep3 struct{ fail bool }

func (s OrgProvisionStep3) Name() string { return "orgprovision_step_3" }
func (s OrgProvisionStep3) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["orgprovision_step_3"] = true
	return nil
}
func (s OrgProvisionStep3) Compensate(ctx context.Context, st map[string]any) error {
	st["orgprovision_step_3"] = false
	return nil
}

type OrgProvisionStep4 struct{ fail bool }

func (s OrgProvisionStep4) Name() string { return "orgprovision_step_4" }
func (s OrgProvisionStep4) Execute(ctx context.Context, st map[string]any) error {
	if s.fail {
		return common.New(common.ErrCodeInternal, "forced failure")
	}
	st["orgprovision_step_4"] = true
	return nil
}
func (s OrgProvisionStep4) Compensate(ctx context.Context, st map[string]any) error {
	st["orgprovision_step_4"] = false
	return nil
}

// RegisterOrgProvision wires OrgProvision saga with compensation hooks.
func RegisterOrgProvision(e *Engine) error {
	return RegisterHandlers(e, "OrgProvision", &OrgProvisionStep1{}, &OrgProvisionStep2{}, &OrgProvisionStep3{}, &OrgProvisionStep4{})
}

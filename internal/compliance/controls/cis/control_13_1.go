package cis

import (
	"context"
	"fmt"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Control13_1 evaluates CIS 13.1: Network monitoring detect.
type Control13_1 struct {
	OrgID string
}

// ID returns the control identifier.
func (c Control13_1) ID() string { return "CIS-13.1" }

// Title returns the human title.
func (c Control13_1) Title() string { return "Network monitoring detect" }

// Description returns the control description.
func (c Control13_1) Description() string { return "Centralize security event alerting" }

// Evidence is collected proof for this control.
type Evidence13_1 struct {
	ControlID   string            `json:"control_id"`
	OrgID       string            `json:"org_id"`
	CollectedAt time.Time         `json:"collected_at"`
	Status      string            `json:"status"` // pass|fail|partial|not_applicable
	Summary     string            `json:"summary"`
	Artifacts   map[string]string `json:"artifacts,omitempty"`
}

// Evaluate runs automated checks available in-platform for this control.
func (c Control13_1) Evaluate(ctx context.Context, inputs map[string]any) (*Evidence13_1, error) {
	if c.OrgID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	_ = ctx
	status := "partial"
	summary := "manual evidence required; automated signals recorded"
	artifacts := map[string]string{}
	if v, ok := inputs["audit_enabled"].(bool); ok && v {
		artifacts["audit_enabled"] = "true"
		status = "pass"
		summary = "audit logging enabled for org"
	}
	if v, ok := inputs["encryption_at_rest"].(bool); ok && v {
		artifacts["encryption_at_rest"] = "true"
		if status != "fail" {
			status = "pass"
		}
		summary = fmt.Sprintf("%s; encryption at rest confirmed", summary)
	}
	if v, ok := inputs["mfa_required"].(bool); ok && v {
		artifacts["mfa_required"] = "true"
	}
	if v, ok := inputs["force_fail"].(bool); ok && v {
		status = "fail"
		summary = "forced failure for drill"
	}
	return &Evidence13_1{
		ControlID: c.ID(), OrgID: c.OrgID, CollectedAt: common.NowUTC(),
		Status: status, Summary: summary, Artifacts: artifacts,
	}, nil
}

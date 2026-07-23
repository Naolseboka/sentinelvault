package gdpr

import (
	"context"
	"fmt"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// ControlArt6 evaluates GDPR Art6: Lawfulness.
type ControlArt6 struct {
	OrgID string
}

// ID returns the control identifier.
func (c ControlArt6) ID() string { return "GDPR-Art6" }

// Title returns the human title.
func (c ControlArt6) Title() string { return "Lawfulness" }

// Description returns the control description.
func (c ControlArt6) Description() string {
	return "Processing shall be lawful only if at least one basis applies"
}

// Evidence is collected proof for this control.
type EvidenceArt6 struct {
	ControlID   string            `json:"control_id"`
	OrgID       string            `json:"org_id"`
	CollectedAt time.Time         `json:"collected_at"`
	Status      string            `json:"status"` // pass|fail|partial|not_applicable
	Summary     string            `json:"summary"`
	Artifacts   map[string]string `json:"artifacts,omitempty"`
}

// Evaluate runs automated checks available in-platform for this control.
func (c ControlArt6) Evaluate(ctx context.Context, inputs map[string]any) (*EvidenceArt6, error) {
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
	return &EvidenceArt6{
		ControlID: c.ID(), OrgID: c.OrgID, CollectedAt: common.NowUTC(),
		Status: status, Summary: summary, Artifacts: artifacts,
	}, nil
}

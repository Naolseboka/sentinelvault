package compliance

import (
	"fmt"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Framework identifiers.
const (
	FrameworkSOC2     = "SOC2"
	FrameworkHIPAA    = "HIPAA"
	FrameworkGDPR     = "GDPR"
	FrameworkISO27001 = "ISO27001"
	FrameworkNIST     = "NIST"
	FrameworkPCIDSS   = "PCI_DSS"
	FrameworkFedRAMP  = "FedRAMP"
	FrameworkCIS      = "CIS"
)

// Control is a compliance control requirement.
type Control struct {
	ID          string `json:"id"`
	Framework   string `json:"framework"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

// Evidence is collected proof for a control.
type Evidence struct {
	ID          string    `json:"id"`
	ControlID   string    `json:"control_id"`
	OrgID       string    `json:"org_id"`
	Source      string    `json:"source"`
	Summary     string    `json:"summary"`
	CollectedAt time.Time `json:"collected_at"`
}

// Report summarizes framework compliance posture.
type Report struct {
	ID            string    `json:"id"`
	OrgID         string    `json:"org_id"`
	Framework     string    `json:"framework"`
	GeneratedAt   time.Time `json:"generated_at"`
	TotalControls int       `json:"total_controls"`
	Covered       int       `json:"covered"`
	CoveragePct   float64   `json:"coverage_pct"`
	Gaps          []string  `json:"gaps"`
	EvidenceIDs   []string  `json:"evidence_ids"`
}

// Catalog returns built-in controls for supported frameworks.
func Catalog(framework string) []Control {
	all := []Control{
		{ID: "SOC2-CC6.1", Framework: FrameworkSOC2, Title: "Logical access controls", Category: "security", Description: "Restrict logical access to systems"},
		{ID: "SOC2-CC7.2", Framework: FrameworkSOC2, Title: "System monitoring", Category: "security", Description: "Monitor system components for anomalies"},
		{ID: "SOC2-CC8.1", Framework: FrameworkSOC2, Title: "Change management", Category: "change", Description: "Authorize and track changes"},
		{ID: "GDPR-32", Framework: FrameworkGDPR, Title: "Security of processing", Category: "privacy", Description: "Implement appropriate technical measures"},
		{ID: "GDPR-33", Framework: FrameworkGDPR, Title: "Breach notification", Category: "privacy", Description: "Notify authority of personal data breaches"},
		{ID: "GDPR-17", Framework: FrameworkGDPR, Title: "Right to erasure", Category: "privacy", Description: "Support data subject erasure requests"},
		{ID: "HIPAA-164.312", Framework: FrameworkHIPAA, Title: "Technical safeguards", Category: "security", Description: "Access control, audit, integrity, auth"},
		{ID: "ISO-A.9.4", Framework: FrameworkISO27001, Title: "System access control", Category: "access", Description: "Secure log-on and privileged access"},
		{ID: "NIST-AC-2", Framework: FrameworkNIST, Title: "Account management", Category: "access", Description: "Manage system accounts"},
		{ID: "PCI-3.4", Framework: FrameworkPCIDSS, Title: "Render PAN unreadable", Category: "crypto", Description: "Encrypt stored cardholder data"},
		{ID: "FED-AC-3", Framework: FrameworkFedRAMP, Title: "Access enforcement", Category: "access", Description: "Enforce approved authorizations"},
		{ID: "CIS-5.2", Framework: FrameworkCIS, Title: "MFA", Category: "identity", Description: "Require MFA for administrative access"},
	}
	if framework == "" {
		return all
	}
	out := make([]Control, 0)
	for _, c := range all {
		if c.Framework == framework {
			out = append(out, c)
		}
	}
	return out
}

// Collector gathers evidence and builds reports.
type Collector struct {
	evidence map[string][]Evidence // orgID -> evidence
}

// NewCollector creates an empty collector.
func NewCollector() *Collector {
	return &Collector{evidence: map[string][]Evidence{}}
}

// AddEvidence stores evidence for an org/control.
func (c *Collector) AddEvidence(orgID, controlID, source, summary string) (*Evidence, error) {
	if orgID == "" || controlID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and control_id required")
	}
	ev := &Evidence{
		ID:          common.NewID(),
		ControlID:   controlID,
		OrgID:       orgID,
		Source:      source,
		Summary:     summary,
		CollectedAt: common.NowUTC(),
	}
	c.evidence[orgID] = append(c.evidence[orgID], *ev)
	return ev, nil
}

// GenerateReport builds a coverage report for a framework.
func (c *Collector) GenerateReport(orgID, framework string) (*Report, error) {
	controls := Catalog(framework)
	if len(controls) == 0 {
		return nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("unknown framework %q", framework))
	}
	coveredSet := map[string]struct{}{}
	ids := make([]string, 0)
	for _, ev := range c.evidence[orgID] {
		for _, ctrl := range controls {
			if ev.ControlID == ctrl.ID {
				coveredSet[ctrl.ID] = struct{}{}
				ids = append(ids, ev.ID)
			}
		}
	}
	gaps := make([]string, 0)
	for _, ctrl := range controls {
		if _, ok := coveredSet[ctrl.ID]; !ok {
			gaps = append(gaps, ctrl.ID)
		}
	}
	covered := len(coveredSet)
	pct := 0.0
	if len(controls) > 0 {
		pct = float64(covered) / float64(len(controls)) * 100
	}
	return &Report{
		ID:            common.NewID(),
		OrgID:         orgID,
		Framework:     framework,
		GeneratedAt:   common.NowUTC(),
		TotalControls: len(controls),
		Covered:       covered,
		CoveragePct:   pct,
		Gaps:          gaps,
		EvidenceIDs:   ids,
	}, nil
}

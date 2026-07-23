package analytics

import (
	"fmt"
	"time"
)

// VaultHealthReport aggregates dashboard metrics for vault health.
type VaultHealthReport struct {
	OrgID       string
	GeneratedAt time.Time
	Sections    []ReportSection
	Score       float64
}

// BuildVaultHealthReport composes a report from metric calculators.
func BuildVaultHealthReport(orgID string, from, to time.Time, metrics ...*LoginHeatmapCalculator) *VaultHealthReport {
	rep := &VaultHealthReport{OrgID: orgID, GeneratedAt: time.Now().UTC(), Sections: make([]ReportSection, 0)}
	var total float64
	for i, m := range metrics {
		if m == nil {
			continue
		}
		avg := m.Average(orgID, from, to)
		total += avg
		rep.Sections = append(rep.Sections, ReportSection{
			Title:  fmt.Sprintf("metric_%d", i),
			Value:  avg,
			Detail: fmt.Sprintf("window %s to %s", from.Format(time.RFC3339), to.Format(time.RFC3339)),
		})
	}
	if len(metrics) > 0 {
		rep.Score = total / float64(len(metrics))
	}
	return rep
}

// MergeReports combines sections from multiple reports.
func MergeVaultHealthReports(reports ...*VaultHealthReport) *VaultHealthReport {
	if len(reports) == 0 {
		return &VaultHealthReport{}
	}
	out := *reports[0]
	for _, r := range reports[1:] {
		if r == nil {
			continue
		}
		out.Sections = append(out.Sections, r.Sections...)
		out.Score = (out.Score + r.Score) / 2
	}
	return &out
}

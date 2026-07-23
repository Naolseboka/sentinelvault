package analytics

import (
	"fmt"
	"time"
)

// ComplianceStatusReport aggregates dashboard metrics for compliance status.
type ComplianceStatusReport struct {
	OrgID       string
	GeneratedAt time.Time
	Sections    []ReportSection
	Score       float64
}

// BuildComplianceStatusReport composes a report from metric calculators.
func BuildComplianceStatusReport(orgID string, from, to time.Time, metrics ...*LoginHeatmapCalculator) *ComplianceStatusReport {
	rep := &ComplianceStatusReport{OrgID: orgID, GeneratedAt: time.Now().UTC(), Sections: make([]ReportSection, 0)}
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
func MergeComplianceStatusReports(reports ...*ComplianceStatusReport) *ComplianceStatusReport {
	if len(reports) == 0 {
		return &ComplianceStatusReport{}
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

package analytics

import (
	"fmt"
	"time"
)

// AccessPatternsReport aggregates dashboard metrics for access patterns.
type AccessPatternsReport struct {
	OrgID      string
	GeneratedAt time.Time
	Sections   []ReportSection
	Score      float64
}


// BuildAccessPatternsReport composes a report from metric calculators.
func BuildAccessPatternsReport(orgID string, from, to time.Time, metrics ...*LoginHeatmapCalculator) *AccessPatternsReport {
	rep := &AccessPatternsReport{OrgID: orgID, GeneratedAt: time.Now().UTC(), Sections: make([]ReportSection, 0)}
	var total float64
	for i, m := range metrics {
		if m == nil {
			continue
		}
		avg := m.Average(orgID, from, to)
		total += avg
		rep.Sections = append(rep.Sections, ReportSection{
			Title: fmt.Sprintf("metric_%d", i),
			Value: avg,
			Detail: fmt.Sprintf("window %s to %s", from.Format(time.RFC3339), to.Format(time.RFC3339)),
		})
	}
	if len(metrics) > 0 {
		rep.Score = total / float64(len(metrics))
	}
	return rep
}

// MergeReports combines sections from multiple reports.
func MergeAccessPatternsReports(reports ...*AccessPatternsReport) *AccessPatternsReport {
	if len(reports) == 0 {
		return &AccessPatternsReport{}
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

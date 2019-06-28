package analytics

import "time"

// ReportSection is one row in a dashboard report.
type ReportSection struct {
	Title  string
	Value  float64
	Detail string
}

// BaseReport holds common report metadata.
type BaseReport struct {
	OrgID       string
	GeneratedAt time.Time
	Score       float64
}

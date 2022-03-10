package domain

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/compliance"
)

// Control re-exports compliance.Control.
type Control = compliance.Control

// Evidence re-exports compliance.Evidence.
type Evidence = compliance.Evidence

// Report re-exports compliance.Report.
type Report = compliance.Report

// EvidenceInput holds evidence collection fields.
type EvidenceInput struct {
	OrgID     string `json:"org_id"`
	ControlID string `json:"control_id"`
	Source    string `json:"source"`
	Summary   string `json:"summary"`
}

// ReportInput holds report generation fields.
type ReportInput struct {
	OrgID     string `json:"org_id"`
	Framework string `json:"framework"`
}

// Service wraps the compliance collector.
type Service struct {
	collector *compliance.Collector
}

// NewService creates a compliance domain service.
func NewService(c *compliance.Collector) *Service {
	return &Service{collector: c}
}

// NewCollector creates an empty evidence collector.
func NewCollector() *compliance.Collector {
	return compliance.NewCollector()
}

// AddEvidence stores compliance evidence.
func (s *Service) AddEvidence(_ context.Context, in EvidenceInput) (*Evidence, error) {
	return s.collector.AddEvidence(in.OrgID, in.ControlID, in.Source, in.Summary)
}

// GenerateReport builds a framework coverage report.
func (s *Service) GenerateReport(_ context.Context, in ReportInput) (*Report, error) {
	return s.collector.GenerateReport(in.OrgID, in.Framework)
}

// ListControls returns catalog controls for a framework.
func (s *Service) ListControls(_ context.Context, framework string) []Control {
	return compliance.Catalog(framework)
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "compliance" }

// TableName returns the primary table name.
func TableName() string { return "evidence" }

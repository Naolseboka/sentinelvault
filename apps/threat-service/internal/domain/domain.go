package domain

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/threat"
)

// Finding re-exports threat.Finding.
type Finding = threat.Finding

// AnalyzePasswordInput holds password analysis fields.
type AnalyzePasswordInput struct {
	OrgID      string `json:"org_id"`
	ResourceID string `json:"resource_id"`
	Password   string `json:"password"`
}

// DetectReuseInput holds reuse detection fields.
type DetectReuseInput struct {
	OrgID         string            `json:"org_id"`
	ItemPasswords map[string]string `json:"item_passwords"`
}

// Service wraps the threat analyzer.
type Service struct {
	analyzer *threat.Analyzer
}

// NewService creates a threat domain service.
func NewService(a *threat.Analyzer) *Service {
	return &Service{analyzer: a}
}

// NewAnalyzer creates a threat analyzer.
func NewAnalyzer() *threat.Analyzer {
	return threat.NewAnalyzer()
}

// AnalyzePassword inspects password strength and weaknesses.
func (s *Service) AnalyzePassword(_ context.Context, in AnalyzePasswordInput) []Finding {
	return s.analyzer.AnalyzePassword(in.OrgID, in.ResourceID, in.Password)
}

// DetectReuse finds passwords reused across vault items.
func (s *Service) DetectReuse(_ context.Context, in DetectReuseInput) []Finding {
	return s.analyzer.DetectReuse(in.OrgID, in.ItemPasswords)
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "threat" }

// TableName returns the primary table name.
func TableName() string { return "findings" }

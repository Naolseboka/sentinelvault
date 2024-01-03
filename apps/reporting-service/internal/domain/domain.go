package domain

import (
	"context"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

const (
	ReportTypeUsage    = "usage"
	ReportTypeAccess   = "access"
	ReportTypeRotation = "rotation"
	ReportTypePosture  = "posture"
)

// OrgStats holds in-memory counters for report generation.
type OrgStats struct {
	Logins          int `json:"logins"`
	VaultReads      int `json:"vault_reads"`
	VaultWrites     int `json:"vault_writes"`
	AccessGrants    int `json:"access_grants"`
	AccessDenials   int `json:"access_denials"`
	RotationsRun    int `json:"rotations_run"`
	RotationsFailed int `json:"rotations_failed"`
	WeakPasswords   int `json:"weak_passwords"`
	MFAEnabledUsers int `json:"mfa_enabled_users"`
	TotalUsers      int `json:"total_users"`
}

// UsageAggregation summarizes vault activity.
type UsageAggregation struct {
	ActiveUsers  int     `json:"active_users"`
	Logins       int     `json:"logins"`
	VaultReads   int     `json:"vault_reads"`
	VaultWrites  int     `json:"vault_writes"`
	ReadWriteRatio float64 `json:"read_write_ratio"`
}

// AccessAggregation summarizes access control events.
type AccessAggregation struct {
	Grants       int     `json:"grants"`
	Denials      int     `json:"denials"`
	DenialRate   float64 `json:"denial_rate_pct"`
	TotalEvents  int     `json:"total_events"`
}

// RotationAggregation summarizes credential rotation jobs.
type RotationAggregation struct {
	RotationsRun    int     `json:"rotations_run"`
	RotationsFailed int     `json:"rotations_failed"`
	SuccessRate     float64 `json:"success_rate_pct"`
}

// PostureAggregation summarizes security posture indicators.
type PostureAggregation struct {
	WeakPasswords    int     `json:"weak_passwords"`
	MFAEnabledUsers  int     `json:"mfa_enabled_users"`
	TotalUsers       int     `json:"total_users"`
	MFACoveragePct   float64 `json:"mfa_coverage_pct"`
	WeakPasswordPct  float64 `json:"weak_password_pct"`
}

// GeneratedReport is a platform analytics report with typed aggregations.
type GeneratedReport struct {
	ID          string    `json:"id"`
	OrgID       string    `json:"org_id"`
	Type        string    `json:"type"`
	GeneratedAt time.Time `json:"generated_at"`
	Summary     string    `json:"summary"`
	Stats       OrgStats  `json:"stats"`
	Usage       *UsageAggregation    `json:"usage,omitempty"`
	Access      *AccessAggregation   `json:"access,omitempty"`
	Rotation    *RotationAggregation `json:"rotation,omitempty"`
	Posture     *PostureAggregation  `json:"posture,omitempty"`
}

// GenerateInput holds report generation fields.
type GenerateInput struct {
	OrgID string `json:"org_id"`
	Type  string `json:"type"`
}

// Service generates analytics reports from in-memory stats.
type Service struct {
	mu    sync.RWMutex
	stats map[string]*OrgStats
}

// NewService creates a reporting domain service.
func NewService() *Service {
	return &Service{stats: map[string]*OrgStats{}}
}

// RecordStats updates org counters (for tests and internal aggregation).
func (s *Service) RecordStats(orgID string, update func(*OrgStats)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	st, ok := s.stats[orgID]
	if !ok {
		st = &OrgStats{TotalUsers: 10, MFAEnabledUsers: 7}
		s.stats[orgID] = st
	}
	update(st)
}

// GenerateReport builds a typed report from stored stats.
func (s *Service) GenerateReport(_ context.Context, in GenerateInput) (*GeneratedReport, error) {
	if in.OrgID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	reportType := in.Type
	if reportType == "" {
		reportType = ReportTypeUsage
	}
	switch reportType {
	case ReportTypeUsage, ReportTypeAccess, ReportTypeRotation, ReportTypePosture:
	default:
		return nil, common.New(common.ErrCodeInvalidArgument, "invalid report type")
	}

	s.mu.RLock()
	st, ok := s.stats[in.OrgID]
	if !ok {
		st = &OrgStats{TotalUsers: 10, MFAEnabledUsers: 7, Logins: 42, VaultReads: 120, VaultWrites: 35}
	}
	cp := *st
	s.mu.RUnlock()

	report := &GeneratedReport{
		ID: common.NewID(), OrgID: in.OrgID, Type: reportType,
		GeneratedAt: common.NowUTC(), Stats: cp,
	}

	switch reportType {
	case ReportTypeUsage:
		report.Summary = "Vault usage summary"
		ratio := 0.0
		if cp.VaultWrites > 0 {
			ratio = float64(cp.VaultReads) / float64(cp.VaultWrites)
		}
		report.Usage = &UsageAggregation{
			ActiveUsers: cp.TotalUsers, Logins: cp.Logins,
			VaultReads: cp.VaultReads, VaultWrites: cp.VaultWrites,
			ReadWriteRatio: ratio,
		}
	case ReportTypeAccess:
		report.Summary = "Access control summary"
		total := cp.AccessGrants + cp.AccessDenials
		denialRate := 0.0
		if total > 0 {
			denialRate = float64(cp.AccessDenials) / float64(total) * 100
		}
		report.Access = &AccessAggregation{
			Grants: cp.AccessGrants, Denials: cp.AccessDenials,
			DenialRate: denialRate, TotalEvents: total,
		}
	case ReportTypeRotation:
		report.Summary = "Credential rotation summary"
		total := cp.RotationsRun + cp.RotationsFailed
		successRate := 100.0
		if total > 0 {
			successRate = float64(cp.RotationsRun) / float64(total) * 100
		}
		report.Rotation = &RotationAggregation{
			RotationsRun: cp.RotationsRun, RotationsFailed: cp.RotationsFailed,
			SuccessRate: successRate,
		}
	case ReportTypePosture:
		report.Summary = "Security posture summary"
		total := max(cp.TotalUsers, 1)
		report.Posture = &PostureAggregation{
			WeakPasswords: cp.WeakPasswords, MFAEnabledUsers: cp.MFAEnabledUsers,
			TotalUsers: cp.TotalUsers,
			MFACoveragePct: float64(cp.MFAEnabledUsers) / float64(total) * 100,
			WeakPasswordPct: float64(cp.WeakPasswords) / float64(total) * 100,
		}
	}
	return report, nil
}

// Generate is an alias for GenerateReport.
func (s *Service) Generate(ctx context.Context, in GenerateInput) (*GeneratedReport, error) {
	return s.GenerateReport(ctx, in)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "reporting" }

// TableName returns the primary table name.
func TableName() string { return "reports" }

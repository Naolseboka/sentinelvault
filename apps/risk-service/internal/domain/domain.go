package domain

import (
	"context"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/risk"
)

// Factors re-exports risk.Factors.
type Factors = risk.Factors

// Score re-exports risk.Score.
type Score = risk.Score

// EvaluateInput holds risk evaluation fields.
type EvaluateInput struct {
	ImpossibleTravel bool   `json:"impossible_travel"`
	NewDevice        bool   `json:"new_device"`
	SuspiciousIP     bool   `json:"suspicious_ip"`
	FailedAttempts   int    `json:"failed_attempts"`
	OffHours         bool   `json:"off_hours"`
	GeoAnomaly       bool   `json:"geo_anomaly"`
	CredentialStuff  bool   `json:"credential_stuffing"`
	DeviceTrust      string `json:"device_trust"`
}

// Service wraps the risk engine.
type Service struct {
	engine *risk.Engine
}

// NewService creates a risk domain service.
func NewService(engine *risk.Engine) *Service {
	return &Service{engine: engine}
}

// NewEngine creates a risk engine with defaults.
func NewEngine() *risk.Engine {
	return risk.NewEngine()
}

// Evaluate computes a risk score.
func (s *Service) Evaluate(_ context.Context, in EvaluateInput) Score {
	return s.engine.Evaluate(Factors{
		ImpossibleTravel: in.ImpossibleTravel, NewDevice: in.NewDevice,
		SuspiciousIP: in.SuspiciousIP, FailedAttempts: in.FailedAttempts,
		OffHours: in.OffHours, GeoAnomaly: in.GeoAnomaly,
		CredentialStuff: in.CredentialStuff, DeviceTrust: in.DeviceTrust,
	})
}

// CheckImpossibleTravel detects implausible travel between login points.
func CheckImpossibleTravel(lat1, lon1 float64, t1 time.Time, lat2, lon2 float64, t2 time.Time) bool {
	return risk.ImpossibleTravel(lat1, lon1, t1, lat2, lon2, t2, 0)
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "risk" }

// TableName returns the primary table name.
func TableName() string { return "evaluations" }

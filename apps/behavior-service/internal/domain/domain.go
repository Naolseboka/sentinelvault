package domain

import (
	"context"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/behavior"
)

// Signal re-exports behavior.Signal.
type Signal = behavior.Signal

// Anomaly re-exports behavior.Anomaly.
type Anomaly = behavior.Anomaly

// ObserveInput holds behavioral observation fields.
type ObserveInput struct {
	UserID    string    `json:"user_id"`
	OrgID     string    `json:"org_id"`
	Type      string    `json:"type"`
	HourUTC   int       `json:"hour_utc"`
	Timestamp time.Time `json:"timestamp"`
}

// Service wraps the behavior tracker.
type Service struct {
	tracker *behavior.Tracker
}

// NewService creates a behavior domain service.
func NewService(t *behavior.Tracker) *Service {
	return &Service{tracker: t}
}

// NewTracker creates a behavior tracker.
func NewTracker() *behavior.Tracker {
	return behavior.NewTracker()
}

// Observe records a signal and returns anomalies.
func (s *Service) Observe(_ context.Context, in ObserveInput) []Anomaly {
	return s.tracker.Observe(Signal{
		UserID: in.UserID, OrgID: in.OrgID, Type: in.Type,
		HourUTC: in.HourUTC, Timestamp: in.Timestamp,
	})
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "behavior" }

// TableName returns the primary table name.
func TableName() string { return "signals" }

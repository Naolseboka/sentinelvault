package domain

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/risk"
)

const (
	TrustUntrusted = "untrusted"
	TrustPending   = "pending"
	TrustTrusted   = "trusted"
)

// DeviceTrustRecord holds device trust assessment bound to a fingerprint.
type DeviceTrustRecord struct {
	ID          string       `json:"id"`
	OrgID       string       `json:"org_id"`
	DeviceID    string       `json:"device_id"`
	UserID      string       `json:"user_id"`
	Fingerprint string       `json:"fingerprint"`
	TrustLevel  string       `json:"trust_level"`
	Score       int          `json:"score"`
	RiskLevel   string       `json:"risk_level"`
	Factors     risk.Factors `json:"factors"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// EvaluateTrustInput holds device trust evaluation fields.
type EvaluateTrustInput struct {
	OrgID          string `json:"org_id"`
	DeviceID       string `json:"device_id"`
	UserID         string `json:"user_id"`
	Fingerprint    string `json:"fingerprint"`
	NewDevice      bool   `json:"new_device"`
	SuspiciousIP   bool   `json:"suspicious_ip"`
	FailedAttempts int    `json:"failed_attempts"`
	ImpossibleTravel bool `json:"impossible_travel"`
	CurrentLevel   string `json:"current_level"`
}

// Service scores device trust using the shared risk engine.
type Service struct {
	mu      sync.RWMutex
	engine  *risk.Engine
	records map[string]*DeviceTrustRecord // fingerprint key -> record
}

// NewService creates a device trust domain service.
func NewService(engine *risk.Engine) *Service {
	return &Service{engine: engine, records: map[string]*DeviceTrustRecord{}}
}

// NewEngine creates a risk engine.
func NewEngine() *risk.Engine {
	return risk.NewEngine()
}

func fingerprintKey(userID, fingerprint string) string {
	if fingerprint == "" {
		return userID
	}
	sum := sha256.Sum256([]byte(userID + ":" + fingerprint))
	return hex.EncodeToString(sum[:])
}

// EvaluateTrust computes trust level from risk factors and binds to fingerprint.
func (s *Service) EvaluateTrust(_ context.Context, in EvaluateTrustInput) (*DeviceTrustRecord, error) {
	if in.OrgID == "" || in.DeviceID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and device_id required")
	}
	if in.Fingerprint == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "fingerprint required")
	}
	level := in.CurrentLevel
	if level == "" {
		if prev, ok := s.lookupByFingerprint(in.UserID, in.Fingerprint); ok {
			level = prev.TrustLevel
		} else {
			level = TrustPending
		}
	}
	factors := risk.Factors{
		NewDevice: in.NewDevice, SuspiciousIP: in.SuspiciousIP,
		FailedAttempts: in.FailedAttempts, DeviceTrust: level,
		ImpossibleTravel: in.ImpossibleTravel,
	}
	score := s.engine.Evaluate(factors)

	trustLevel := TrustTrusted
	switch {
	case score.Block || score.Level == "critical":
		trustLevel = TrustUntrusted
	case score.RequireMFA || level == TrustPending || in.NewDevice || score.Level == "high":
		trustLevel = TrustPending
	}

	rec := &DeviceTrustRecord{
		ID: common.NewID(), OrgID: in.OrgID, DeviceID: in.DeviceID, UserID: in.UserID,
		Fingerprint: in.Fingerprint, TrustLevel: trustLevel, Score: score.Value,
		RiskLevel: score.Level, Factors: factors, UpdatedAt: common.NowUTC(),
	}
	fpKey := fingerprintKey(in.UserID, in.Fingerprint)
	s.mu.Lock()
	s.records[fpKey] = rec
	s.records[in.DeviceID] = rec
	s.mu.Unlock()
	cp := *rec
	return &cp, nil
}

func (s *Service) lookupByFingerprint(userID, fingerprint string) (*DeviceTrustRecord, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	rec, ok := s.records[fingerprintKey(userID, fingerprint)]
	return rec, ok
}

// Get returns the latest trust record for a device.
func (s *Service) Get(_ context.Context, deviceID string) (*DeviceTrustRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	rec, ok := s.records[deviceID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "device trust record not found")
	}
	cp := *rec
	return &cp, nil
}

// GetByFingerprint returns trust for a user/fingerprint pair.
func (s *Service) GetByFingerprint(_ context.Context, userID, fingerprint string) (*DeviceTrustRecord, error) {
	rec, ok := s.lookupByFingerprint(userID, fingerprint)
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "device trust record not found")
	}
	cp := *rec
	return &cp, nil
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "device_trust" }

// TableName returns the primary table name.
func TableName() string { return "trust_records" }

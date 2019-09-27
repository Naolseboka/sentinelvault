package risk

import (
	"math"
	"strings"
	"time"
)

// Factors contributing to a session/login risk score.
type Factors struct {
	ImpossibleTravel bool
	NewDevice        bool
	SuspiciousIP     bool
	FailedAttempts   int
	OffHours         bool
	GeoAnomaly       bool
	CredentialStuff  bool
	DeviceTrust      string // untrusted|pending|trusted
}

// Score is a 0-100 risk assessment.
type Score struct {
	Value      int      `json:"value"`
	Level      string   `json:"level"` // low|medium|high|critical
	Reasons    []string `json:"reasons"`
	RequireMFA bool     `json:"require_mfa"`
	Block      bool     `json:"block"`
}

// Engine computes adaptive authentication risk.
type Engine struct {
	HighThreshold     int
	CriticalThreshold int
}

// NewEngine returns defaults.
func NewEngine() *Engine {
	return &Engine{HighThreshold: 60, CriticalThreshold: 85}
}

// Evaluate produces a risk score from factors.
func (e *Engine) Evaluate(f Factors) Score {
	score := 0
	reasons := make([]string, 0, 8)

	add := func(pts int, reason string) {
		score += pts
		reasons = append(reasons, reason)
	}

	if f.ImpossibleTravel {
		add(40, "impossible_travel")
	}
	if f.NewDevice {
		add(15, "new_device")
	}
	if f.SuspiciousIP {
		add(25, "suspicious_ip")
	}
	if f.FailedAttempts > 0 {
		pts := int(math.Min(30, float64(f.FailedAttempts*5)))
		add(pts, "failed_attempts")
	}
	if f.OffHours {
		add(10, "off_hours")
	}
	if f.GeoAnomaly {
		add(20, "geo_anomaly")
	}
	if f.CredentialStuff {
		add(35, "credential_stuffing_pattern")
	}
	switch strings.ToLower(f.DeviceTrust) {
	case "untrusted":
		add(20, "untrusted_device")
	case "pending":
		add(10, "pending_device")
	}

	if score > 100 {
		score = 100
	}
	level := "low"
	requireMFA := score >= 40
	block := false
	switch {
	case score >= e.CriticalThreshold:
		level = "critical"
		block = true
		requireMFA = true
	case score >= e.HighThreshold:
		level = "high"
		requireMFA = true
	case score >= 40:
		level = "medium"
		requireMFA = true
	}
	return Score{Value: score, Level: level, Reasons: reasons, RequireMFA: requireMFA, Block: block}
}

// ImpossibleTravel detects travel exceeding max speed between two login points.
func ImpossibleTravel(lat1, lon1 float64, t1 time.Time, lat2, lon2 float64, t2 time.Time, maxKmh float64) bool {
	if maxKmh <= 0 {
		maxKmh = 900 // commercial flight
	}
	hours := math.Abs(t2.Sub(t1).Hours())
	if hours < 0.05 {
		hours = 0.05
	}
	dist := haversineKm(lat1, lon1, lat2, lon2)
	return dist/hours > maxKmh
}

func haversineKm(lat1, lon1, lat2, lon2 float64) float64 {
	const r = 6371.0
	toRad := func(d float64) float64 { return d * math.Pi / 180 }
	dLat := toRad(lat2 - lat1)
	dLon := toRad(lon2 - lon1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(toRad(lat1))*math.Cos(toRad(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)
	return 2 * r * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

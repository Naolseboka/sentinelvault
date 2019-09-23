package behavior

import (
	"math"
	"sync"
	"time"
)

// EventType for behavioral signals.
const (
	EventLogin      = "login"
	EventVaultRead  = "vault_read"
	EventVaultWrite = "vault_write"
	EventExport     = "export"
	EventAdmin      = "admin"
)

// Signal is a recorded user action.
type Signal struct {
	UserID    string
	OrgID     string
	Type      string
	HourUTC   int
	Timestamp time.Time
}

// Baseline captures typical activity hours and rates.
type Baseline struct {
	UserID       string             `json:"user_id"`
	HourHistogram [24]int           `json:"hour_histogram"`
	TypeCounts   map[string]int     `json:"type_counts"`
	Total        int                `json:"total"`
}

// Anomaly describes unusual behavior.
type Anomaly struct {
	UserID  string  `json:"user_id"`
	OrgID   string  `json:"org_id"`
	Type    string  `json:"type"`
	Score   float64 `json:"score"`
	Reason  string  `json:"reason"`
}

// Tracker maintains per-user baselines.
type Tracker struct {
	mu        sync.Mutex
	baselines map[string]*Baseline
}

// NewTracker creates a tracker.
func NewTracker() *Tracker {
	return &Tracker{baselines: map[string]*Baseline{}}
}

// Observe records a signal and returns anomalies if any.
func (t *Tracker) Observe(s Signal) []Anomaly {
	t.mu.Lock()
	defer t.mu.Unlock()
	b, ok := t.baselines[s.UserID]
	if !ok {
		b = &Baseline{UserID: s.UserID, TypeCounts: map[string]int{}}
		t.baselines[s.UserID] = b
	}
	hour := s.HourUTC
	if s.Timestamp.IsZero() {
		hour = time.Now().UTC().Hour()
	} else if hour == 0 && s.Timestamp.Hour() != 0 {
		hour = s.Timestamp.Hour()
	} else if !s.Timestamp.IsZero() {
		hour = s.Timestamp.Hour()
	}
	anoms := make([]Anomaly, 0)
	if b.Total >= 20 {
		// Unusual hour: less than 5% of historical activity
		hist := float64(b.HourHistogram[hour]) / float64(b.Total)
		if hist < 0.05 {
			anoms = append(anoms, Anomaly{
				UserID: s.UserID, OrgID: s.OrgID, Type: "unusual_hour",
				Score: math.Min(1, (0.05-hist)*10), Reason: "activity outside normal hours",
			})
		}
		if s.Type == EventExport && b.TypeCounts[EventExport] == 0 {
			anoms = append(anoms, Anomaly{
				UserID: s.UserID, OrgID: s.OrgID, Type: "first_export",
				Score: 0.8, Reason: "first export for user",
			})
		}
	}
	b.HourHistogram[hour]++
	b.TypeCounts[s.Type]++
	b.Total++
	return anoms
}

// BaselineOf returns a copy of user baseline.
func (t *Tracker) BaselineOf(userID string) (Baseline, bool) {
	t.mu.Lock()
	defer t.mu.Unlock()
	b, ok := t.baselines[userID]
	if !ok {
		return Baseline{}, false
	}
	cp := *b
	cp.TypeCounts = map[string]int{}
	for k, v := range b.TypeCounts {
		cp.TypeCounts[k] = v
	}
	return cp, true
}

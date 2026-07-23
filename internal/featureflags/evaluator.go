package featureflags

import (
	"sync"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// BooleanFlag is a simple on/off toggle.
type BooleanFlag struct {
	Key     string
	Enabled bool
}

// PercentageFlag rolls out to a fraction of subjects.
type PercentageFlag struct {
	Key        string
	Percentage int
	Enabled    bool
}

// OrgTargetedFlag allows per-org overrides.
type OrgTargetedFlag struct {
	Key      string
	Enabled  bool
	OrgAllow map[string]bool
	OrgDeny  map[string]bool
}

// Evaluator combines registry evaluation helpers.
type Evaluator struct {
	reg   *Registry
	mu    sync.RWMutex
	cache map[string]bool
}

func NewEvaluator(reg *Registry) *Evaluator {
	return &Evaluator{reg: reg, cache: map[string]bool{}}
}

func (ev *Evaluator) IsEnabled(key, orgID, subjectID string) (bool, error) {
	ck := key + ":" + orgID + ":" + subjectID
	ev.mu.RLock()
	if v, ok := ev.cache[ck]; ok {
		ev.mu.RUnlock()
		return v, nil
	}
	ev.mu.RUnlock()
	ok, err := ev.reg.Evaluate(key, orgID, subjectID)
	if err != nil {
		return false, err
	}
	ev.mu.Lock()
	ev.cache[ck] = ok
	ev.mu.Unlock()
	return ok, nil
}

func UpsertBoolean(reg *Registry, f BooleanFlag) error {
	return reg.Upsert(Flag{Key: f.Key, Enabled: f.Enabled, Percentage: 100})
}

func UpsertPercentage(reg *Registry, f PercentageFlag) error {
	if !f.Enabled {
		return reg.Upsert(Flag{Key: f.Key, Enabled: false, Percentage: 0})
	}
	return reg.Upsert(Flag{Key: f.Key, Enabled: true, Percentage: f.Percentage})
}

func UpsertOrgTargeted(reg *Registry, f OrgTargetedFlag) error {
	if f.Key == "" {
		return common.New(common.ErrCodeInvalidArgument, "key required")
	}
	return reg.Upsert(Flag{
		Key: f.Key, Enabled: f.Enabled, Percentage: 100,
		OrgAllow: f.OrgAllow, OrgDeny: f.OrgDeny,
	})
}

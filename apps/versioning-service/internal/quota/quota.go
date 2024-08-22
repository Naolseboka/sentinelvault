package quota

import (
	"sync"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Limits struct {
	MaxResources int
	MaxCreatesPerHour int
	MaxExportsPerDay int
	MaxShares int
}

func DefaultLimits() Limits {
	return Limits{MaxResources: 10000, MaxCreatesPerHour: 1000, MaxExportsPerDay: 100, MaxShares: 500}
}

type Usage struct {
	Resources int
	CreatesHour int
	ExportsDay int
	Shares int
}

type Tracker struct {
	mu sync.Mutex
	limits map[string]Limits
	usage map[string]*Usage
}

func NewTracker() *Tracker {
	return &Tracker{limits: map[string]Limits{}, usage: map[string]*Usage{}}
}

func (t *Tracker) SetLimits(orgID string, l Limits) {
	t.mu.Lock()
	t.limits[orgID] = l
	t.mu.Unlock()
}

func (t *Tracker) get(orgID string) (Limits, *Usage) {
	l, ok := t.limits[orgID]
	if !ok {
		l = DefaultLimits()
	}
	u := t.usage[orgID]
	if u == nil {
		u = &Usage{}
		t.usage[orgID] = u
	}
	return l, u
}

func (t *Tracker) CheckCreate(orgID string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	l, u := t.get(orgID)
	if u.Resources >= l.MaxResources {
		return common.New(common.ErrCodeForbidden, "versioning: resource quota exceeded")
	}
	if u.CreatesHour >= l.MaxCreatesPerHour {
		return common.New(common.ErrCodeForbidden, "versioning: create rate exceeded")
	}
	return nil
}

func (t *Tracker) RecordCreate(orgID string) {
	t.mu.Lock()
	_, u := t.get(orgID)
	u.Resources++
	u.CreatesHour++
	t.mu.Unlock()
}

func (t *Tracker) CheckExport(orgID string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	l, u := t.get(orgID)
	if u.ExportsDay >= l.MaxExportsPerDay {
		return common.New(common.ErrCodeForbidden, "versioning: export quota exceeded")
	}
	return nil
}

func (t *Tracker) RecordExport(orgID string) {
	t.mu.Lock()
	_, u := t.get(orgID)
	u.ExportsDay++
	t.mu.Unlock()
}

func (t *Tracker) CheckShare(orgID string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	l, u := t.get(orgID)
	if u.Shares >= l.MaxShares {
		return common.New(common.ErrCodeForbidden, "versioning: share quota exceeded")
	}
	return nil
}

func (t *Tracker) RecordShare(orgID string) {
	t.mu.Lock()
	_, u := t.get(orgID)
	u.Shares++
	t.mu.Unlock()
}

func (t *Tracker) Snapshot(orgID string) (Limits, Usage) {
	t.mu.Lock()
	defer t.mu.Unlock()
	l, u := t.get(orgID)
	return l, *u
}

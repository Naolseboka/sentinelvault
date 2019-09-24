package featureflags

import (
	"hash/fnv"
	"sync"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Flag struct {
	Key         string
	Enabled     bool
	Percentage  int            // 0-100 sticky rollout
	OrgAllow    map[string]bool
	OrgDeny     map[string]bool
}

type Registry struct {
	mu    sync.RWMutex
	flags map[string]*Flag
}

func NewRegistry() *Registry { return &Registry{flags: map[string]*Flag{}} }

func (r *Registry) Upsert(f Flag) error {
	if f.Key == "" {
		return common.New(common.ErrCodeInvalidArgument, "key required")
	}
	if f.Percentage < 0 || f.Percentage > 100 {
		return common.New(common.ErrCodeInvalidArgument, "percentage out of range")
	}
	if f.OrgAllow == nil {
		f.OrgAllow = map[string]bool{}
	}
	if f.OrgDeny == nil {
		f.OrgDeny = map[string]bool{}
	}
	cp := f
	r.mu.Lock()
	r.flags[f.Key] = &cp
	r.mu.Unlock()
	return nil
}

func (r *Registry) Evaluate(key, orgID, subjectID string) (bool, error) {
	r.mu.RLock()
	f, ok := r.flags[key]
	r.mu.RUnlock()
	if !ok {
		return false, common.New(common.ErrCodeNotFound, "flag not found")
	}
	if f.OrgDeny[orgID] {
		return false, nil
	}
	if f.OrgAllow[orgID] {
		return true, nil
	}
	if !f.Enabled {
		return false, nil
	}
	if f.Percentage >= 100 {
		return true, nil
	}
	if f.Percentage <= 0 {
		return false, nil
	}
	h := fnv.New32a()
	_, _ = h.Write([]byte(key + ":" + orgID + ":" + subjectID))
	bucket := int(h.Sum32() % 100)
	return bucket < f.Percentage, nil
}

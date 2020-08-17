package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// LifecycleStateView is a query-side read model for lifecycle-service.
type LifecycleStateView struct {
	ID        string            `json:"id"`
	OrgID     string            `json:"org_id"`
	Name      string            `json:"name"`
	Status    string            `json:"status"`
	ActorID   string            `json:"actor_id,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Score     float64           `json:"score,omitempty"`
}

// LifecycleStateQuery holds filter/sort/page options.
type LifecycleStateQuery struct {
	OrgID        string
	Status       string
	NameContains string
	ActorID      string
	LabelKey     string
	LabelValue   string
	SortBy       string // name|created_at|updated_at|status
	SortDesc     bool
	Offset       int
	Limit        int
}

// LifecycleStateIndex is an in-memory query index (CQRS read side).
type LifecycleStateIndex struct {
	items map[string]LifecycleStateView
}

// NewLifecycleStateIndex creates an empty index.
func NewLifecycleStateIndex() *LifecycleStateIndex {
	return &LifecycleStateIndex{items: map[string]LifecycleStateView{}}
}

// Upsert indexes a view.
func (idx *LifecycleStateIndex) Upsert(v LifecycleStateView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *LifecycleStateIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *LifecycleStateIndex) Get(id string) (LifecycleStateView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *LifecycleStateIndex) Search(q LifecycleStateQuery) ([]LifecycleStateView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]LifecycleStateView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterLifecycleStateByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterLifecycleStateByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterLifecycleStateByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]LifecycleStateView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortLifecycleStateViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []LifecycleStateView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterLifecycleStateByStatus(in []LifecycleStateView, status string) []LifecycleStateView {
	o := make([]LifecycleStateView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterLifecycleStateByNameContains(in []LifecycleStateView, needle string) []LifecycleStateView {
	n := strings.ToLower(needle)
	o := make([]LifecycleStateView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterLifecycleStateByActorID(in []LifecycleStateView, actor string) []LifecycleStateView {
	o := make([]LifecycleStateView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortLifecycleStateViews(in []LifecycleStateView, by string, desc bool) {
	sort.SliceStable(in, func(i, j int) bool {
		var less bool
		switch by {
		case "name":
			less = in[i].Name < in[j].Name
		case "status":
			less = in[i].Status < in[j].Status
		case "updated_at":
			less = in[i].UpdatedAt.Before(in[j].UpdatedAt)
		default:
			less = in[i].CreatedAt.Before(in[j].CreatedAt)
		}
		if desc {
			return !less
		}
		return less
	})
}

// KnownLifecycleStateStatuses returns allowed statuses for lifecycle-service.
func KnownLifecycleStateStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidLifecycleStateStatus reports whether s is allowed.
func IsValidLifecycleStateStatus(s string) bool {
	for _, k := range KnownLifecycleStateStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

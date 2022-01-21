package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// SignalView is a query-side read model for behavior-service.
type SignalView struct {
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

// SignalQuery holds filter/sort/page options.
type SignalQuery struct {
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

// SignalIndex is an in-memory query index (CQRS read side).
type SignalIndex struct {
	items map[string]SignalView
}

// NewSignalIndex creates an empty index.
func NewSignalIndex() *SignalIndex {
	return &SignalIndex{items: map[string]SignalView{}}
}

// Upsert indexes a view.
func (idx *SignalIndex) Upsert(v SignalView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *SignalIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *SignalIndex) Get(id string) (SignalView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *SignalIndex) Search(q SignalQuery) ([]SignalView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]SignalView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterSignalByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterSignalByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterSignalByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]SignalView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortSignalViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []SignalView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterSignalByStatus(in []SignalView, status string) []SignalView {
	o := make([]SignalView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterSignalByNameContains(in []SignalView, needle string) []SignalView {
	n := strings.ToLower(needle)
	o := make([]SignalView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterSignalByActorID(in []SignalView, actor string) []SignalView {
	o := make([]SignalView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortSignalViews(in []SignalView, by string, desc bool) {
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

// KnownSignalStatuses returns allowed statuses for behavior-service.
func KnownSignalStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidSignalStatus reports whether s is allowed.
func IsValidSignalStatus(s string) bool {
	for _, k := range KnownSignalStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

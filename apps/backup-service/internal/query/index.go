package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// SnapshotView is a query-side read model for backup-service.
type SnapshotView struct {
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

// SnapshotQuery holds filter/sort/page options.
type SnapshotQuery struct {
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

// SnapshotIndex is an in-memory query index (CQRS read side).
type SnapshotIndex struct {
	items map[string]SnapshotView
}

// NewSnapshotIndex creates an empty index.
func NewSnapshotIndex() *SnapshotIndex {
	return &SnapshotIndex{items: map[string]SnapshotView{}}
}

// Upsert indexes a view.
func (idx *SnapshotIndex) Upsert(v SnapshotView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *SnapshotIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *SnapshotIndex) Get(id string) (SnapshotView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *SnapshotIndex) Search(q SnapshotQuery) ([]SnapshotView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]SnapshotView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterSnapshotByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterSnapshotByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterSnapshotByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]SnapshotView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortSnapshotViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []SnapshotView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterSnapshotByStatus(in []SnapshotView, status string) []SnapshotView {
	o := make([]SnapshotView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterSnapshotByNameContains(in []SnapshotView, needle string) []SnapshotView {
	n := strings.ToLower(needle)
	o := make([]SnapshotView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterSnapshotByActorID(in []SnapshotView, actor string) []SnapshotView {
	o := make([]SnapshotView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortSnapshotViews(in []SnapshotView, by string, desc bool) {
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

// KnownSnapshotStatuses returns allowed statuses for backup-service.
func KnownSnapshotStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidSnapshotStatus reports whether s is allowed.
func IsValidSnapshotStatus(s string) bool {
	for _, k := range KnownSnapshotStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

// maintenance note 56

package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// SyncJobView is a query-side read model for directory-sync-service.
type SyncJobView struct {
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

// SyncJobQuery holds filter/sort/page options.
type SyncJobQuery struct {
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

// SyncJobIndex is an in-memory query index (CQRS read side).
type SyncJobIndex struct {
	items map[string]SyncJobView
}

// NewSyncJobIndex creates an empty index.
func NewSyncJobIndex() *SyncJobIndex {
	return &SyncJobIndex{items: map[string]SyncJobView{}}
}

// Upsert indexes a view.
func (idx *SyncJobIndex) Upsert(v SyncJobView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *SyncJobIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *SyncJobIndex) Get(id string) (SyncJobView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *SyncJobIndex) Search(q SyncJobQuery) ([]SyncJobView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]SyncJobView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterSyncJobByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterSyncJobByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterSyncJobByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]SyncJobView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortSyncJobViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []SyncJobView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterSyncJobByStatus(in []SyncJobView, status string) []SyncJobView {
	o := make([]SyncJobView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterSyncJobByNameContains(in []SyncJobView, needle string) []SyncJobView {
	n := strings.ToLower(needle)
	o := make([]SyncJobView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterSyncJobByActorID(in []SyncJobView, actor string) []SyncJobView {
	o := make([]SyncJobView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortSyncJobViews(in []SyncJobView, by string, desc bool) {
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

// KnownSyncJobStatuses returns allowed statuses for directory-sync-service.
func KnownSyncJobStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidSyncJobStatus reports whether s is allowed.
func IsValidSyncJobStatus(s string) bool {
	for _, k := range KnownSyncJobStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

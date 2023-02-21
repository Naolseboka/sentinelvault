package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// FindingView is a query-side read model for threat-service.
type FindingView struct {
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

// FindingQuery holds filter/sort/page options.
type FindingQuery struct {
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

// FindingIndex is an in-memory query index (CQRS read side).
type FindingIndex struct {
	items map[string]FindingView
}

// NewFindingIndex creates an empty index.
func NewFindingIndex() *FindingIndex {
	return &FindingIndex{items: map[string]FindingView{}}
}

// Upsert indexes a view.
func (idx *FindingIndex) Upsert(v FindingView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *FindingIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *FindingIndex) Get(id string) (FindingView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *FindingIndex) Search(q FindingQuery) ([]FindingView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]FindingView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterFindingByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterFindingByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterFindingByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]FindingView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortFindingViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []FindingView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterFindingByStatus(in []FindingView, status string) []FindingView {
	o := make([]FindingView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterFindingByNameContains(in []FindingView, needle string) []FindingView {
	n := strings.ToLower(needle)
	o := make([]FindingView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterFindingByActorID(in []FindingView, actor string) []FindingView {
	o := make([]FindingView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortFindingViews(in []FindingView, by string, desc bool) {
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

// KnownFindingStatuses returns allowed statuses for threat-service.
func KnownFindingStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidFindingStatus reports whether s is allowed.
func IsValidFindingStatus(s string) bool {
	for _, k := range KnownFindingStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

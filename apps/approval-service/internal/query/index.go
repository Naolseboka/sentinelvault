package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// ApprovalView is a query-side read model for approval-service.
type ApprovalView struct {
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

// ApprovalQuery holds filter/sort/page options.
type ApprovalQuery struct {
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

// ApprovalIndex is an in-memory query index (CQRS read side).
type ApprovalIndex struct {
	items map[string]ApprovalView
}

// NewApprovalIndex creates an empty index.
func NewApprovalIndex() *ApprovalIndex {
	return &ApprovalIndex{items: map[string]ApprovalView{}}
}

// Upsert indexes a view.
func (idx *ApprovalIndex) Upsert(v ApprovalView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *ApprovalIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *ApprovalIndex) Get(id string) (ApprovalView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *ApprovalIndex) Search(q ApprovalQuery) ([]ApprovalView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]ApprovalView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterApprovalByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterApprovalByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterApprovalByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]ApprovalView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortApprovalViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []ApprovalView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterApprovalByStatus(in []ApprovalView, status string) []ApprovalView {
	o := make([]ApprovalView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterApprovalByNameContains(in []ApprovalView, needle string) []ApprovalView {
	n := strings.ToLower(needle)
	o := make([]ApprovalView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterApprovalByActorID(in []ApprovalView, actor string) []ApprovalView {
	o := make([]ApprovalView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortApprovalViews(in []ApprovalView, by string, desc bool) {
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

// KnownApprovalStatuses returns allowed statuses for approval-service.
func KnownApprovalStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidApprovalStatus reports whether s is allowed.
func IsValidApprovalStatus(s string) bool {
	for _, k := range KnownApprovalStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

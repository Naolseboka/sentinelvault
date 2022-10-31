package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// PolicyDocView is a query-side read model for policy-service.
type PolicyDocView struct {
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

// PolicyDocQuery holds filter/sort/page options.
type PolicyDocQuery struct {
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

// PolicyDocIndex is an in-memory query index (CQRS read side).
type PolicyDocIndex struct {
	items map[string]PolicyDocView
}

// NewPolicyDocIndex creates an empty index.
func NewPolicyDocIndex() *PolicyDocIndex {
	return &PolicyDocIndex{items: map[string]PolicyDocView{}}
}

// Upsert indexes a view.
func (idx *PolicyDocIndex) Upsert(v PolicyDocView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *PolicyDocIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *PolicyDocIndex) Get(id string) (PolicyDocView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *PolicyDocIndex) Search(q PolicyDocQuery) ([]PolicyDocView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]PolicyDocView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterPolicyDocByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterPolicyDocByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterPolicyDocByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]PolicyDocView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortPolicyDocViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []PolicyDocView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterPolicyDocByStatus(in []PolicyDocView, status string) []PolicyDocView {
	o := make([]PolicyDocView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterPolicyDocByNameContains(in []PolicyDocView, needle string) []PolicyDocView {
	n := strings.ToLower(needle)
	o := make([]PolicyDocView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterPolicyDocByActorID(in []PolicyDocView, actor string) []PolicyDocView {
	o := make([]PolicyDocView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortPolicyDocViews(in []PolicyDocView, by string, desc bool) {
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

// KnownPolicyDocStatuses returns allowed statuses for policy-service.
func KnownPolicyDocStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidPolicyDocStatus reports whether s is allowed.
func IsValidPolicyDocStatus(s string) bool {
	for _, k := range KnownPolicyDocStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

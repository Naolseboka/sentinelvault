package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// RoleBindingView is a query-side read model for rbac-service.
type RoleBindingView struct {
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

// RoleBindingQuery holds filter/sort/page options.
type RoleBindingQuery struct {
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

// RoleBindingIndex is an in-memory query index (CQRS read side).
type RoleBindingIndex struct {
	items map[string]RoleBindingView
}

// NewRoleBindingIndex creates an empty index.
func NewRoleBindingIndex() *RoleBindingIndex {
	return &RoleBindingIndex{items: map[string]RoleBindingView{}}
}

// Upsert indexes a view.
func (idx *RoleBindingIndex) Upsert(v RoleBindingView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *RoleBindingIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *RoleBindingIndex) Get(id string) (RoleBindingView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *RoleBindingIndex) Search(q RoleBindingQuery) ([]RoleBindingView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]RoleBindingView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterRoleBindingByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterRoleBindingByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterRoleBindingByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]RoleBindingView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortRoleBindingViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []RoleBindingView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterRoleBindingByStatus(in []RoleBindingView, status string) []RoleBindingView {
	o := make([]RoleBindingView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterRoleBindingByNameContains(in []RoleBindingView, needle string) []RoleBindingView {
	n := strings.ToLower(needle)
	o := make([]RoleBindingView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterRoleBindingByActorID(in []RoleBindingView, actor string) []RoleBindingView {
	o := make([]RoleBindingView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortRoleBindingViews(in []RoleBindingView, by string, desc bool) {
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

// KnownRoleBindingStatuses returns allowed statuses for rbac-service.
func KnownRoleBindingStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidRoleBindingStatus reports whether s is allowed.
func IsValidRoleBindingStatus(s string) bool {
	for _, k := range KnownRoleBindingStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

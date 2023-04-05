package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// AdminActionView is a query-side read model for admin-service.
type AdminActionView struct {
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

// AdminActionQuery holds filter/sort/page options.
type AdminActionQuery struct {
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

// AdminActionIndex is an in-memory query index (CQRS read side).
type AdminActionIndex struct {
	items map[string]AdminActionView
}

// NewAdminActionIndex creates an empty index.
func NewAdminActionIndex() *AdminActionIndex {
	return &AdminActionIndex{items: map[string]AdminActionView{}}
}

// Upsert indexes a view.
func (idx *AdminActionIndex) Upsert(v AdminActionView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *AdminActionIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *AdminActionIndex) Get(id string) (AdminActionView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *AdminActionIndex) Search(q AdminActionQuery) ([]AdminActionView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]AdminActionView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterAdminActionByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterAdminActionByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterAdminActionByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]AdminActionView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortAdminActionViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []AdminActionView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterAdminActionByStatus(in []AdminActionView, status string) []AdminActionView {
	o := make([]AdminActionView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterAdminActionByNameContains(in []AdminActionView, needle string) []AdminActionView {
	n := strings.ToLower(needle)
	o := make([]AdminActionView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterAdminActionByActorID(in []AdminActionView, actor string) []AdminActionView {
	o := make([]AdminActionView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortAdminActionViews(in []AdminActionView, by string, desc bool) {
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

// KnownAdminActionStatuses returns allowed statuses for admin-service.
func KnownAdminActionStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidAdminActionStatus reports whether s is allowed.
func IsValidAdminActionStatus(s string) bool {
	for _, k := range KnownAdminActionStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

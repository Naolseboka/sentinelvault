package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// UserView is a query-side read model for user-service.
type UserView struct {
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

// UserQuery holds filter/sort/page options.
type UserQuery struct {
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

// UserIndex is an in-memory query index (CQRS read side).
type UserIndex struct {
	items map[string]UserView
}

// NewUserIndex creates an empty index.
func NewUserIndex() *UserIndex {
	return &UserIndex{items: map[string]UserView{}}
}

// Upsert indexes a view.
func (idx *UserIndex) Upsert(v UserView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *UserIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *UserIndex) Get(id string) (UserView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *UserIndex) Search(q UserQuery) ([]UserView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]UserView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterUserByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterUserByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterUserByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]UserView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortUserViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []UserView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterUserByStatus(in []UserView, status string) []UserView {
	o := make([]UserView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterUserByNameContains(in []UserView, needle string) []UserView {
	n := strings.ToLower(needle)
	o := make([]UserView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterUserByActorID(in []UserView, actor string) []UserView {
	o := make([]UserView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortUserViews(in []UserView, by string, desc bool) {
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

// KnownUserStatuses returns allowed statuses for user-service.
func KnownUserStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidUserStatus reports whether s is allowed.
func IsValidUserStatus(s string) bool {
	for _, k := range KnownUserStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

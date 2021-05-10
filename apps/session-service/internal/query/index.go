package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// SessionView is a query-side read model for session-service.
type SessionView struct {
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

// SessionQuery holds filter/sort/page options.
type SessionQuery struct {
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

// SessionIndex is an in-memory query index (CQRS read side).
type SessionIndex struct {
	items map[string]SessionView
}

// NewSessionIndex creates an empty index.
func NewSessionIndex() *SessionIndex {
	return &SessionIndex{items: map[string]SessionView{}}
}

// Upsert indexes a view.
func (idx *SessionIndex) Upsert(v SessionView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *SessionIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *SessionIndex) Get(id string) (SessionView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *SessionIndex) Search(q SessionQuery) ([]SessionView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]SessionView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterSessionByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterSessionByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterSessionByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]SessionView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortSessionViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []SessionView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterSessionByStatus(in []SessionView, status string) []SessionView {
	o := make([]SessionView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterSessionByNameContains(in []SessionView, needle string) []SessionView {
	n := strings.ToLower(needle)
	o := make([]SessionView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterSessionByActorID(in []SessionView, actor string) []SessionView {
	o := make([]SessionView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortSessionViews(in []SessionView, by string, desc bool) {
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

// KnownSessionStatuses returns allowed statuses for session-service.
func KnownSessionStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidSessionStatus reports whether s is allowed.
func IsValidSessionStatus(s string) bool {
	for _, k := range KnownSessionStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

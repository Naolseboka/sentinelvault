package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// PasskeyView is a query-side read model for passkey-service.
type PasskeyView struct {
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

// PasskeyQuery holds filter/sort/page options.
type PasskeyQuery struct {
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

// PasskeyIndex is an in-memory query index (CQRS read side).
type PasskeyIndex struct {
	items map[string]PasskeyView
}

// NewPasskeyIndex creates an empty index.
func NewPasskeyIndex() *PasskeyIndex {
	return &PasskeyIndex{items: map[string]PasskeyView{}}
}

// Upsert indexes a view.
func (idx *PasskeyIndex) Upsert(v PasskeyView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *PasskeyIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *PasskeyIndex) Get(id string) (PasskeyView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *PasskeyIndex) Search(q PasskeyQuery) ([]PasskeyView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]PasskeyView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterPasskeyByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterPasskeyByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterPasskeyByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]PasskeyView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortPasskeyViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []PasskeyView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterPasskeyByStatus(in []PasskeyView, status string) []PasskeyView {
	o := make([]PasskeyView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterPasskeyByNameContains(in []PasskeyView, needle string) []PasskeyView {
	n := strings.ToLower(needle)
	o := make([]PasskeyView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterPasskeyByActorID(in []PasskeyView, actor string) []PasskeyView {
	o := make([]PasskeyView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortPasskeyViews(in []PasskeyView, by string, desc bool) {
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

// KnownPasskeyStatuses returns allowed statuses for passkey-service.
func KnownPasskeyStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidPasskeyStatus reports whether s is allowed.
func IsValidPasskeyStatus(s string) bool {
	for _, k := range KnownPasskeyStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// SecretItemView is a query-side read model for secret-service.
type SecretItemView struct {
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

// SecretItemQuery holds filter/sort/page options.
type SecretItemQuery struct {
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

// SecretItemIndex is an in-memory query index (CQRS read side).
type SecretItemIndex struct {
	items map[string]SecretItemView
}

// NewSecretItemIndex creates an empty index.
func NewSecretItemIndex() *SecretItemIndex {
	return &SecretItemIndex{items: map[string]SecretItemView{}}
}

// Upsert indexes a view.
func (idx *SecretItemIndex) Upsert(v SecretItemView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *SecretItemIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *SecretItemIndex) Get(id string) (SecretItemView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *SecretItemIndex) Search(q SecretItemQuery) ([]SecretItemView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]SecretItemView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterSecretItemByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterSecretItemByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterSecretItemByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]SecretItemView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortSecretItemViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []SecretItemView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterSecretItemByStatus(in []SecretItemView, status string) []SecretItemView {
	o := make([]SecretItemView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterSecretItemByNameContains(in []SecretItemView, needle string) []SecretItemView {
	n := strings.ToLower(needle)
	o := make([]SecretItemView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterSecretItemByActorID(in []SecretItemView, actor string) []SecretItemView {
	o := make([]SecretItemView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortSecretItemViews(in []SecretItemView, by string, desc bool) {
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

// KnownSecretItemStatuses returns allowed statuses for secret-service.
func KnownSecretItemStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidSecretItemStatus reports whether s is allowed.
func IsValidSecretItemStatus(s string) bool {
	for _, k := range KnownSecretItemStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

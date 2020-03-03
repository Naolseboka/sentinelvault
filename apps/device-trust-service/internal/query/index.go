package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// TrustRecordView is a query-side read model for device-trust-service.
type TrustRecordView struct {
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

// TrustRecordQuery holds filter/sort/page options.
type TrustRecordQuery struct {
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

// TrustRecordIndex is an in-memory query index (CQRS read side).
type TrustRecordIndex struct {
	items map[string]TrustRecordView
}

// NewTrustRecordIndex creates an empty index.
func NewTrustRecordIndex() *TrustRecordIndex {
	return &TrustRecordIndex{items: map[string]TrustRecordView{}}
}

// Upsert indexes a view.
func (idx *TrustRecordIndex) Upsert(v TrustRecordView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *TrustRecordIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *TrustRecordIndex) Get(id string) (TrustRecordView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *TrustRecordIndex) Search(q TrustRecordQuery) ([]TrustRecordView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]TrustRecordView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterTrustRecordByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterTrustRecordByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterTrustRecordByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]TrustRecordView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortTrustRecordViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []TrustRecordView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterTrustRecordByStatus(in []TrustRecordView, status string) []TrustRecordView {
	o := make([]TrustRecordView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterTrustRecordByNameContains(in []TrustRecordView, needle string) []TrustRecordView {
	n := strings.ToLower(needle)
	o := make([]TrustRecordView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterTrustRecordByActorID(in []TrustRecordView, actor string) []TrustRecordView {
	o := make([]TrustRecordView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortTrustRecordViews(in []TrustRecordView, by string, desc bool) {
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

// KnownTrustRecordStatuses returns allowed statuses for device-trust-service.
func KnownTrustRecordStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidTrustRecordStatus reports whether s is allowed.
func IsValidTrustRecordStatus(s string) bool {
	for _, k := range KnownTrustRecordStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

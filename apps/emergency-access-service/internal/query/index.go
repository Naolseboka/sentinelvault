package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// EmergencyGrantView is a query-side read model for emergency-access-service.
type EmergencyGrantView struct {
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

// EmergencyGrantQuery holds filter/sort/page options.
type EmergencyGrantQuery struct {
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

// EmergencyGrantIndex is an in-memory query index (CQRS read side).
type EmergencyGrantIndex struct {
	items map[string]EmergencyGrantView
}

// NewEmergencyGrantIndex creates an empty index.
func NewEmergencyGrantIndex() *EmergencyGrantIndex {
	return &EmergencyGrantIndex{items: map[string]EmergencyGrantView{}}
}

// Upsert indexes a view.
func (idx *EmergencyGrantIndex) Upsert(v EmergencyGrantView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *EmergencyGrantIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *EmergencyGrantIndex) Get(id string) (EmergencyGrantView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *EmergencyGrantIndex) Search(q EmergencyGrantQuery) ([]EmergencyGrantView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]EmergencyGrantView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterEmergencyGrantByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterEmergencyGrantByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterEmergencyGrantByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]EmergencyGrantView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortEmergencyGrantViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []EmergencyGrantView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterEmergencyGrantByStatus(in []EmergencyGrantView, status string) []EmergencyGrantView {
	o := make([]EmergencyGrantView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterEmergencyGrantByNameContains(in []EmergencyGrantView, needle string) []EmergencyGrantView {
	n := strings.ToLower(needle)
	o := make([]EmergencyGrantView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterEmergencyGrantByActorID(in []EmergencyGrantView, actor string) []EmergencyGrantView {
	o := make([]EmergencyGrantView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortEmergencyGrantViews(in []EmergencyGrantView, by string, desc bool) {
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

// KnownEmergencyGrantStatuses returns allowed statuses for emergency-access-service.
func KnownEmergencyGrantStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidEmergencyGrantStatus reports whether s is allowed.
func IsValidEmergencyGrantStatus(s string) bool {
	for _, k := range KnownEmergencyGrantStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

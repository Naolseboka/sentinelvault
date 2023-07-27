package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// LicenseView is a query-side read model for billing-service.
type LicenseView struct {
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

// LicenseQuery holds filter/sort/page options.
type LicenseQuery struct {
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

// LicenseIndex is an in-memory query index (CQRS read side).
type LicenseIndex struct {
	items map[string]LicenseView
}

// NewLicenseIndex creates an empty index.
func NewLicenseIndex() *LicenseIndex {
	return &LicenseIndex{items: map[string]LicenseView{}}
}

// Upsert indexes a view.
func (idx *LicenseIndex) Upsert(v LicenseView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *LicenseIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *LicenseIndex) Get(id string) (LicenseView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *LicenseIndex) Search(q LicenseQuery) ([]LicenseView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]LicenseView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterLicenseByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterLicenseByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterLicenseByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]LicenseView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortLicenseViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []LicenseView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterLicenseByStatus(in []LicenseView, status string) []LicenseView {
	o := make([]LicenseView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterLicenseByNameContains(in []LicenseView, needle string) []LicenseView {
	n := strings.ToLower(needle)
	o := make([]LicenseView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterLicenseByActorID(in []LicenseView, actor string) []LicenseView {
	o := make([]LicenseView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortLicenseViews(in []LicenseView, by string, desc bool) {
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

// KnownLicenseStatuses returns allowed statuses for billing-service.
func KnownLicenseStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidLicenseStatus reports whether s is allowed.
func IsValidLicenseStatus(s string) bool {
	for _, k := range KnownLicenseStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

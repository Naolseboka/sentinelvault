package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// AuditEventView is a query-side read model for audit-service.
type AuditEventView struct {
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

// AuditEventQuery holds filter/sort/page options.
type AuditEventQuery struct {
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

// AuditEventIndex is an in-memory query index (CQRS read side).
type AuditEventIndex struct {
	items map[string]AuditEventView
}

// NewAuditEventIndex creates an empty index.
func NewAuditEventIndex() *AuditEventIndex {
	return &AuditEventIndex{items: map[string]AuditEventView{}}
}

// Upsert indexes a view.
func (idx *AuditEventIndex) Upsert(v AuditEventView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *AuditEventIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *AuditEventIndex) Get(id string) (AuditEventView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *AuditEventIndex) Search(q AuditEventQuery) ([]AuditEventView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]AuditEventView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterAuditEventByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterAuditEventByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterAuditEventByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]AuditEventView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortAuditEventViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []AuditEventView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterAuditEventByStatus(in []AuditEventView, status string) []AuditEventView {
	o := make([]AuditEventView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterAuditEventByNameContains(in []AuditEventView, needle string) []AuditEventView {
	n := strings.ToLower(needle)
	o := make([]AuditEventView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterAuditEventByActorID(in []AuditEventView, actor string) []AuditEventView {
	o := make([]AuditEventView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortAuditEventViews(in []AuditEventView, by string, desc bool) {
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

// KnownAuditEventStatuses returns allowed statuses for audit-service.
func KnownAuditEventStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidAuditEventStatus reports whether s is allowed.
func IsValidAuditEventStatus(s string) bool {
	for _, k := range KnownAuditEventStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

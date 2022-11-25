package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// RiskScoreView is a query-side read model for risk-service.
type RiskScoreView struct {
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

// RiskScoreQuery holds filter/sort/page options.
type RiskScoreQuery struct {
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

// RiskScoreIndex is an in-memory query index (CQRS read side).
type RiskScoreIndex struct {
	items map[string]RiskScoreView
}

// NewRiskScoreIndex creates an empty index.
func NewRiskScoreIndex() *RiskScoreIndex {
	return &RiskScoreIndex{items: map[string]RiskScoreView{}}
}

// Upsert indexes a view.
func (idx *RiskScoreIndex) Upsert(v RiskScoreView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *RiskScoreIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *RiskScoreIndex) Get(id string) (RiskScoreView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *RiskScoreIndex) Search(q RiskScoreQuery) ([]RiskScoreView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]RiskScoreView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterRiskScoreByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterRiskScoreByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterRiskScoreByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]RiskScoreView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortRiskScoreViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []RiskScoreView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterRiskScoreByStatus(in []RiskScoreView, status string) []RiskScoreView {
	o := make([]RiskScoreView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterRiskScoreByNameContains(in []RiskScoreView, needle string) []RiskScoreView {
	n := strings.ToLower(needle)
	o := make([]RiskScoreView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterRiskScoreByActorID(in []RiskScoreView, actor string) []RiskScoreView {
	o := make([]RiskScoreView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortRiskScoreViews(in []RiskScoreView, by string, desc bool) {
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

// KnownRiskScoreStatuses returns allowed statuses for risk-service.
func KnownRiskScoreStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidRiskScoreStatus reports whether s is allowed.
func IsValidRiskScoreStatus(s string) bool {
	for _, k := range KnownRiskScoreStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

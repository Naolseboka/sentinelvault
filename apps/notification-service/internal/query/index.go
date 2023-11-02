package query

import (
	"sort"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// MessageView is a query-side read model for notification-service.
type MessageView struct {
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

// MessageQuery holds filter/sort/page options.
type MessageQuery struct {
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

// MessageIndex is an in-memory query index (CQRS read side).
type MessageIndex struct {
	items map[string]MessageView
}

// NewMessageIndex creates an empty index.
func NewMessageIndex() *MessageIndex {
	return &MessageIndex{items: map[string]MessageView{}}
}

// Upsert indexes a view.
func (idx *MessageIndex) Upsert(v MessageView) {
	if v.Labels == nil {
		v.Labels = map[string]string{}
	}
	idx.items[v.ID] = v
}

// Delete removes a view.
func (idx *MessageIndex) Delete(id string) {
	delete(idx.items, id)
}

// Get returns a view by id.
func (idx *MessageIndex) Get(id string) (MessageView, bool) {
	v, ok := idx.items[id]
	return v, ok
}

// Search applies filters, sorting, and pagination.
func (idx *MessageIndex) Search(q MessageQuery) ([]MessageView, int64, error) {
	if q.OrgID == "" {
		return nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 500 {
		q.Limit = 500
	}
	out := make([]MessageView, 0)
	for _, v := range idx.items {
		if v.OrgID != q.OrgID {
			continue
		}
		out = append(out, v)
	}

	if q.Status != "" {
		out = filterMessageByStatus(out, q.Status)
	}

	if q.NameContains != "" {
		out = filterMessageByNameContains(out, q.NameContains)
	}

	if q.ActorID != "" {
		out = filterMessageByActorID(out, q.ActorID)
	}
	if q.LabelKey != "" {
		filtered := make([]MessageView, 0)
		for _, v := range out {
			if val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {
				filtered = append(filtered, v)
			}
		}
		out = filtered
	}
	sortMessageViews(out, q.SortBy, q.SortDesc)
	total := int64(len(out))
	if q.Offset >= len(out) {
		return []MessageView{}, total, nil
	}
	end := q.Offset + q.Limit
	if end > len(out) {
		end = len(out)
	}
	return out[q.Offset:end], total, nil
}

func filterMessageByStatus(in []MessageView, status string) []MessageView {
	o := make([]MessageView, 0, len(in))
	for _, v := range in {
		if v.Status == status {
			o = append(o, v)
		}
	}
	return o
}

func filterMessageByNameContains(in []MessageView, needle string) []MessageView {
	n := strings.ToLower(needle)
	o := make([]MessageView, 0, len(in))
	for _, v := range in {
		if strings.Contains(strings.ToLower(v.Name), n) {
			o = append(o, v)
		}
	}
	return o
}

func filterMessageByActorID(in []MessageView, actor string) []MessageView {
	o := make([]MessageView, 0, len(in))
	for _, v := range in {
		if v.ActorID == actor {
			o = append(o, v)
		}
	}
	return o
}

func sortMessageViews(in []MessageView, by string, desc bool) {
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

// KnownMessageStatuses returns allowed statuses for notification-service.
func KnownMessageStatuses() []string {
	return []string{"active", "inactive", "pending", "approved", "denied", "expired", "revoked", "archived", "deleted", "draft", "scheduled", "running", "failed", "completed"}
}

// IsValidMessageStatus reports whether s is allowed.
func IsValidMessageStatus(s string) bool {
	for _, k := range KnownMessageStatuses() {
		if k == s {
			return true
		}
	}
	return false
}

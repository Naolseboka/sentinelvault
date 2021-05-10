package filters

import (
	"strings"
	"time"
)

// Filter defines query filters for session-service.
type Filter struct {
	Id string `json:"id"`
	OrgId string `json:"org_id"`
	Name string `json:"name"`
	Status string `json:"status"`
	OwnerId string `json:"owner_id"`
	CollectionId string `json:"collection_id"`
	FolderId string `json:"folder_id"`
	Tag string `json:"tag"`
	Label string `json:"label"`
	CreatedBefore string `json:"created_before"`
	CreatedAfter string `json:"created_after"`
	UpdatedBefore string `json:"updated_before"`
	UpdatedAfter string `json:"updated_after"`
	Risk string `json:"risk"`
	Trust string `json:"trust"`
	Classification string `json:"classification"`
	ActorId string `json:"actor_id"`
	ResourceType string `json:"resource_type"`
	ResourceId string `json:"resource_id"`
	Severity string `json:"severity"`
	Channel string `json:"channel"`
	Mode string `json:"mode"` // exact|prefix|contains
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

// Item is a generic filterable record.
type Item struct {
	Fields map[string]string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func getMeta(item Item, key string) string {
	if item.Fields == nil {
		return ""
	}
	return item.Fields[key]
}

func matchString(have, want, mode string) bool {
	switch mode {
	case "prefix":
		return strings.HasPrefix(strings.ToLower(have), strings.ToLower(want))
	case "contains":
		return strings.Contains(strings.ToLower(have), strings.ToLower(want))
	default:
		return strings.EqualFold(have, want)
	}
}

func Match(item Item, f Filter) bool {
	if f.Id != "" && !matchString(getMeta(item, "id"), f.Id, f.Mode) {
		return false
	}
	if f.OrgId != "" && !matchString(getMeta(item, "org_id"), f.OrgId, f.Mode) {
		return false
	}
	if f.Name != "" && !matchString(getMeta(item, "name"), f.Name, f.Mode) {
		return false
	}
	if f.Status != "" && !matchString(getMeta(item, "status"), f.Status, f.Mode) {
		return false
	}
	if f.OwnerId != "" && !matchString(getMeta(item, "owner_id"), f.OwnerId, f.Mode) {
		return false
	}
	if f.CollectionId != "" && !matchString(getMeta(item, "collection_id"), f.CollectionId, f.Mode) {
		return false
	}
	if f.FolderId != "" && !matchString(getMeta(item, "folder_id"), f.FolderId, f.Mode) {
		return false
	}
	if f.Tag != "" && !matchString(getMeta(item, "tag"), f.Tag, f.Mode) {
		return false
	}
	if f.Label != "" && !matchString(getMeta(item, "label"), f.Label, f.Mode) {
		return false
	}
	if f.CreatedBefore != "" && !matchString(getMeta(item, "created_before"), f.CreatedBefore, f.Mode) {
		return false
	}
	if f.CreatedAfter != "" && !matchString(getMeta(item, "created_after"), f.CreatedAfter, f.Mode) {
		return false
	}
	if f.UpdatedBefore != "" && !matchString(getMeta(item, "updated_before"), f.UpdatedBefore, f.Mode) {
		return false
	}
	if f.UpdatedAfter != "" && !matchString(getMeta(item, "updated_after"), f.UpdatedAfter, f.Mode) {
		return false
	}
	if f.Risk != "" && !matchString(getMeta(item, "risk"), f.Risk, f.Mode) {
		return false
	}
	if f.Trust != "" && !matchString(getMeta(item, "trust"), f.Trust, f.Mode) {
		return false
	}
	if f.Classification != "" && !matchString(getMeta(item, "classification"), f.Classification, f.Mode) {
		return false
	}
	if f.ActorId != "" && !matchString(getMeta(item, "actor_id"), f.ActorId, f.Mode) {
		return false
	}
	if f.ResourceType != "" && !matchString(getMeta(item, "resource_type"), f.ResourceType, f.Mode) {
		return false
	}
	if f.ResourceId != "" && !matchString(getMeta(item, "resource_id"), f.ResourceId, f.Mode) {
		return false
	}
	if f.Severity != "" && !matchString(getMeta(item, "severity"), f.Severity, f.Mode) {
		return false
	}
	if f.Channel != "" && !matchString(getMeta(item, "channel"), f.Channel, f.Mode) {
		return false
	}
	if f.CreatedBefore != "" {
		if ts, err := time.Parse(time.RFC3339, f.CreatedBefore); err == nil && item.CreatedAt.After(ts) {
			return false
		}
	}
	if f.CreatedAfter != "" {
		if ts, err := time.Parse(time.RFC3339, f.CreatedAfter); err == nil && item.CreatedAt.Before(ts) {
			return false
		}
	}
	if f.UpdatedBefore != "" {
		if ts, err := time.Parse(time.RFC3339, f.UpdatedBefore); err == nil && item.UpdatedAt.After(ts) {
			return false
		}
	}
	if f.UpdatedAfter != "" {
		if ts, err := time.Parse(time.RFC3339, f.UpdatedAfter); err == nil && item.UpdatedAt.Before(ts) {
			return false
		}
	}
	return true
}

func Apply(items []Item, f Filter) []Item {
	out := make([]Item, 0)
	for _, it := range items {
		if Match(it, f) {
			out = append(out, it)
		}
	}
	if f.Offset > len(out) {
		return []Item{}
	}
	out = out[f.Offset:]
	if f.Limit > 0 && len(out) > f.Limit {
		out = out[:f.Limit]
	}
	return out
}

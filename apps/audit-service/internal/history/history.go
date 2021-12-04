package history

import (
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type HistoryEntry struct {
	ID string `json:"id"`
	OrgID string `json:"org_id"`
	ActorID string `json:"actor_id"`
	ResourceID string `json:"resource_id"`
	Action string `json:"action"`
	Details map[string]string `json:"details,omitempty"`
	At time.Time `json:"at"`
}

type AuditEventHistory struct {
	mu sync.RWMutex
	entries []HistoryEntry
	max int
}

func NewAuditEventHistory(max int) *AuditEventHistory {
	if max <= 0 {
		max = 10000
	}
	return &AuditEventHistory{entries: make([]HistoryEntry, 0, 128), max: max}
}

func (h *AuditEventHistory) append(orgID, actorID, resourceID, action string, details map[string]string) HistoryEntry {
	e := HistoryEntry{ID: common.NewID(), OrgID: orgID, ActorID: actorID, ResourceID: resourceID, Action: action, Details: details, At: common.NowUTC()}
	h.mu.Lock()
	h.entries = append(h.entries, e)
	if len(h.entries) > h.max {
		h.entries = h.entries[len(h.entries)-h.max:]
	}
	h.mu.Unlock()
	return e
}

func (h *AuditEventHistory) List(orgID, resourceID string, limit int) []HistoryEntry {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if limit <= 0 {
		limit = 100
	}
	out := make([]HistoryEntry, 0, limit)
	for i := len(h.entries) - 1; i >= 0 && len(out) < limit; i-- {
		e := h.entries[i]
		if orgID != "" && e.OrgID != orgID {
			continue
		}
		if resourceID != "" && e.ResourceID != resourceID {
			continue
		}
		out = append(out, e)
	}
	return out
}

func (h *AuditEventHistory) Count(orgID string) int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	n := 0
	for _, e := range h.entries {
		if orgID == "" || e.OrgID == orgID {
			n++
		}
	}
	return n
}

func (h *AuditEventHistory) RecordNameChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_name", map[string]string{"from": from, "to": to, "field": "name"})
}

func (h *AuditEventHistory) RecordStatusChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_status", map[string]string{"from": from, "to": to, "field": "status"})
}

func (h *AuditEventHistory) RecordOwnerChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_owner", map[string]string{"from": from, "to": to, "field": "owner"})
}

func (h *AuditEventHistory) RecordTagsChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_tags", map[string]string{"from": from, "to": to, "field": "tags"})
}

func (h *AuditEventHistory) RecordMetadataChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_metadata", map[string]string{"from": from, "to": to, "field": "metadata"})
}

func (h *AuditEventHistory) RecordPayloadChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_payload", map[string]string{"from": from, "to": to, "field": "payload"})
}

func (h *AuditEventHistory) RecordPermissionChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_permission", map[string]string{"from": from, "to": to, "field": "permission"})
}

func (h *AuditEventHistory) RecordTtlChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_ttl", map[string]string{"from": from, "to": to, "field": "ttl"})
}

func (h *AuditEventHistory) RecordEnabledChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_enabled", map[string]string{"from": from, "to": to, "field": "enabled"})
}

func (h *AuditEventHistory) RecordMfaChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_mfa", map[string]string{"from": from, "to": to, "field": "mfa"})
}

func (h *AuditEventHistory) RecordDeviceChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_device", map[string]string{"from": from, "to": to, "field": "device"})
}

func (h *AuditEventHistory) RecordSessionChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_session", map[string]string{"from": from, "to": to, "field": "session"})
}

func (h *AuditEventHistory) RecordCollectionChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_collection", map[string]string{"from": from, "to": to, "field": "collection"})
}

func (h *AuditEventHistory) RecordFolderChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_folder", map[string]string{"from": from, "to": to, "field": "folder"})
}

func (h *AuditEventHistory) RecordShareChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_share", map[string]string{"from": from, "to": to, "field": "share"})
}

func (h *AuditEventHistory) RecordPolicyChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_policy", map[string]string{"from": from, "to": to, "field": "policy"})
}

func (h *AuditEventHistory) RecordRoleChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_role", map[string]string{"from": from, "to": to, "field": "role"})
}

func (h *AuditEventHistory) RecordSecretChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_secret", map[string]string{"from": from, "to": to, "field": "secret"})
}

func (h *AuditEventHistory) RecordKeyChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_key", map[string]string{"from": from, "to": to, "field": "key"})
}

func (h *AuditEventHistory) RecordCertificateChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_certificate", map[string]string{"from": from, "to": to, "field": "certificate"})
}

func (h *AuditEventHistory) RecordLeaseChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_lease", map[string]string{"from": from, "to": to, "field": "lease"})
}

func (h *AuditEventHistory) RecordScheduleChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_schedule", map[string]string{"from": from, "to": to, "field": "schedule"})
}

func (h *AuditEventHistory) RecordApproverChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_approver", map[string]string{"from": from, "to": to, "field": "approver"})
}

func (h *AuditEventHistory) RecordRiskChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_risk", map[string]string{"from": from, "to": to, "field": "risk"})
}

func (h *AuditEventHistory) RecordTrustChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_trust", map[string]string{"from": from, "to": to, "field": "trust"})
}

func (h *AuditEventHistory) RecordLicenseChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_license", map[string]string{"from": from, "to": to, "field": "license"})
}

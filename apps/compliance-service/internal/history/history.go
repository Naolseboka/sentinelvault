package history

import (
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type HistoryEntry struct {
	ID         string            `json:"id"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	Action     string            `json:"action"`
	Details    map[string]string `json:"details,omitempty"`
	At         time.Time         `json:"at"`
}

type EvidenceHistory struct {
	mu      sync.RWMutex
	entries []HistoryEntry
	max     int
}

func NewEvidenceHistory(max int) *EvidenceHistory {
	if max <= 0 {
		max = 10000
	}
	return &EvidenceHistory{entries: make([]HistoryEntry, 0, 128), max: max}
}

func (h *EvidenceHistory) append(orgID, actorID, resourceID, action string, details map[string]string) HistoryEntry {
	e := HistoryEntry{ID: common.NewID(), OrgID: orgID, ActorID: actorID, ResourceID: resourceID, Action: action, Details: details, At: common.NowUTC()}
	h.mu.Lock()
	h.entries = append(h.entries, e)
	if len(h.entries) > h.max {
		h.entries = h.entries[len(h.entries)-h.max:]
	}
	h.mu.Unlock()
	return e
}

func (h *EvidenceHistory) List(orgID, resourceID string, limit int) []HistoryEntry {
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

func (h *EvidenceHistory) Count(orgID string) int {
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

func (h *EvidenceHistory) RecordNameChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_name", map[string]string{"from": from, "to": to, "field": "name"})
}

func (h *EvidenceHistory) RecordStatusChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_status", map[string]string{"from": from, "to": to, "field": "status"})
}

func (h *EvidenceHistory) RecordOwnerChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_owner", map[string]string{"from": from, "to": to, "field": "owner"})
}

func (h *EvidenceHistory) RecordTagsChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_tags", map[string]string{"from": from, "to": to, "field": "tags"})
}

func (h *EvidenceHistory) RecordMetadataChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_metadata", map[string]string{"from": from, "to": to, "field": "metadata"})
}

func (h *EvidenceHistory) RecordPayloadChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_payload", map[string]string{"from": from, "to": to, "field": "payload"})
}

func (h *EvidenceHistory) RecordPermissionChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_permission", map[string]string{"from": from, "to": to, "field": "permission"})
}

func (h *EvidenceHistory) RecordTtlChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_ttl", map[string]string{"from": from, "to": to, "field": "ttl"})
}

func (h *EvidenceHistory) RecordEnabledChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_enabled", map[string]string{"from": from, "to": to, "field": "enabled"})
}

func (h *EvidenceHistory) RecordMfaChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_mfa", map[string]string{"from": from, "to": to, "field": "mfa"})
}

func (h *EvidenceHistory) RecordDeviceChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_device", map[string]string{"from": from, "to": to, "field": "device"})
}

func (h *EvidenceHistory) RecordSessionChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_session", map[string]string{"from": from, "to": to, "field": "session"})
}

func (h *EvidenceHistory) RecordCollectionChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_collection", map[string]string{"from": from, "to": to, "field": "collection"})
}

func (h *EvidenceHistory) RecordFolderChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_folder", map[string]string{"from": from, "to": to, "field": "folder"})
}

func (h *EvidenceHistory) RecordShareChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_share", map[string]string{"from": from, "to": to, "field": "share"})
}

func (h *EvidenceHistory) RecordPolicyChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_policy", map[string]string{"from": from, "to": to, "field": "policy"})
}

func (h *EvidenceHistory) RecordRoleChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_role", map[string]string{"from": from, "to": to, "field": "role"})
}

func (h *EvidenceHistory) RecordSecretChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_secret", map[string]string{"from": from, "to": to, "field": "secret"})
}

func (h *EvidenceHistory) RecordKeyChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_key", map[string]string{"from": from, "to": to, "field": "key"})
}

func (h *EvidenceHistory) RecordCertificateChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_certificate", map[string]string{"from": from, "to": to, "field": "certificate"})
}

func (h *EvidenceHistory) RecordLeaseChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_lease", map[string]string{"from": from, "to": to, "field": "lease"})
}

func (h *EvidenceHistory) RecordScheduleChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_schedule", map[string]string{"from": from, "to": to, "field": "schedule"})
}

func (h *EvidenceHistory) RecordApproverChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_approver", map[string]string{"from": from, "to": to, "field": "approver"})
}

func (h *EvidenceHistory) RecordRiskChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_risk", map[string]string{"from": from, "to": to, "field": "risk"})
}

func (h *EvidenceHistory) RecordTrustChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_trust", map[string]string{"from": from, "to": to, "field": "trust"})
}

func (h *EvidenceHistory) RecordLicenseChange(orgID, actorID, resourceID, from, to string) HistoryEntry {
	return h.append(orgID, actorID, resourceID, "change_license", map[string]string{"from": from, "to": to, "field": "license"})
}

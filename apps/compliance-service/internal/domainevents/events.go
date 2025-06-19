package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for compliance-service.
const (
	EventEvidenceCreated = "compliance.created"
	EventEvidenceUpdated = "compliance.updated"
	EventEvidenceDeleted = "compliance.deleted"
	EventEvidenceArchived = "compliance.archived"
	EventEvidenceRestored = "compliance.restored"
	EventEvidenceShared = "compliance.shared"
	EventEvidenceUnshared = "compliance.unshared"
	EventEvidenceExported = "compliance.exported"
	EventEvidenceImported = "compliance.imported"
	EventEvidenceRotated = "compliance.rotated"
	EventEvidenceAccessed = "compliance.accessed"
	EventEvidenceDenied = "compliance.denied"
	EventEvidenceApproved = "compliance.approved"
	EventEvidenceRevoked = "compliance.revoked"
	EventEvidenceExpired = "compliance.expired"
	EventEvidenceLocked = "compliance.locked"
	EventEvidenceUnlocked = "compliance.unlocked"
	EventEvidenceVerified = "compliance.verified"
	EventEvidenceFailed = "compliance.failed"
	EventEvidenceSucceeded = "compliance.succeeded"
)

// Event is a compliance-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownEvidenceEvent reports whether t is a known compliance event type.
func IsKnownEvidenceEvent(t string) bool {
	switch t {
	case EventEvidenceCreated:
		return true
	case EventEvidenceUpdated:
		return true
	case EventEvidenceDeleted:
		return true
	case EventEvidenceArchived:
		return true
	case EventEvidenceRestored:
		return true
	case EventEvidenceShared:
		return true
	case EventEvidenceUnshared:
		return true
	case EventEvidenceExported:
		return true
	case EventEvidenceImported:
		return true
	case EventEvidenceRotated:
		return true
	case EventEvidenceAccessed:
		return true
	case EventEvidenceDenied:
		return true
	case EventEvidenceApproved:
		return true
	case EventEvidenceRevoked:
		return true
	case EventEvidenceExpired:
		return true
	case EventEvidenceLocked:
		return true
	case EventEvidenceUnlocked:
		return true
	case EventEvidenceVerified:
		return true
	case EventEvidenceFailed:
		return true
	case EventEvidenceSucceeded:
		return true
	default:
		return false
	}
}


// NewEvidenceCreatedEvent builds a domain event.
func NewEvidenceCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventEvidenceCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceUpdatedEvent builds a domain event.
func NewEvidenceUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventEvidenceUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceDeletedEvent builds a domain event.
func NewEvidenceDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventEvidenceDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceArchivedEvent builds a domain event.
func NewEvidenceArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventEvidenceArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceRestoredEvent builds a domain event.
func NewEvidenceRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventEvidenceRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceSharedEvent builds a domain event.
func NewEvidenceSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventEvidenceShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceUnsharedEvent builds a domain event.
func NewEvidenceUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventEvidenceUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceExportedEvent builds a domain event.
func NewEvidenceExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventEvidenceExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceImportedEvent builds a domain event.
func NewEvidenceImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventEvidenceImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceRotatedEvent builds a domain event.
func NewEvidenceRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventEvidenceRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceAccessedEvent builds a domain event.
func NewEvidenceAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventEvidenceAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceDeniedEvent builds a domain event.
func NewEvidenceDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventEvidenceDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceApprovedEvent builds a domain event.
func NewEvidenceApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventEvidenceApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceRevokedEvent builds a domain event.
func NewEvidenceRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventEvidenceRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceExpiredEvent builds a domain event.
func NewEvidenceExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventEvidenceExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceLockedEvent builds a domain event.
func NewEvidenceLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventEvidenceLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceUnlockedEvent builds a domain event.
func NewEvidenceUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventEvidenceUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceVerifiedEvent builds a domain event.
func NewEvidenceVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventEvidenceVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceFailedEvent builds a domain event.
func NewEvidenceFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventEvidenceFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEvidenceSucceededEvent builds a domain event.
func NewEvidenceSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Evidence"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventEvidenceSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownEvidenceEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}
// maintenance note 32

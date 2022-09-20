package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for kms-service.
const (
	EventKeyCreated = "kms.created"
	EventKeyUpdated = "kms.updated"
	EventKeyDeleted = "kms.deleted"
	EventKeyArchived = "kms.archived"
	EventKeyRestored = "kms.restored"
	EventKeyShared = "kms.shared"
	EventKeyUnshared = "kms.unshared"
	EventKeyExported = "kms.exported"
	EventKeyImported = "kms.imported"
	EventKeyRotated = "kms.rotated"
	EventKeyAccessed = "kms.accessed"
	EventKeyDenied = "kms.denied"
	EventKeyApproved = "kms.approved"
	EventKeyRevoked = "kms.revoked"
	EventKeyExpired = "kms.expired"
	EventKeyLocked = "kms.locked"
	EventKeyUnlocked = "kms.unlocked"
	EventKeyVerified = "kms.verified"
	EventKeyFailed = "kms.failed"
	EventKeySucceeded = "kms.succeeded"
)

// Event is a kms-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownKeyEvent reports whether t is a known kms event type.
func IsKnownKeyEvent(t string) bool {
	switch t {
	case EventKeyCreated:
		return true
	case EventKeyUpdated:
		return true
	case EventKeyDeleted:
		return true
	case EventKeyArchived:
		return true
	case EventKeyRestored:
		return true
	case EventKeyShared:
		return true
	case EventKeyUnshared:
		return true
	case EventKeyExported:
		return true
	case EventKeyImported:
		return true
	case EventKeyRotated:
		return true
	case EventKeyAccessed:
		return true
	case EventKeyDenied:
		return true
	case EventKeyApproved:
		return true
	case EventKeyRevoked:
		return true
	case EventKeyExpired:
		return true
	case EventKeyLocked:
		return true
	case EventKeyUnlocked:
		return true
	case EventKeyVerified:
		return true
	case EventKeyFailed:
		return true
	case EventKeySucceeded:
		return true
	default:
		return false
	}
}


// NewKeyCreatedEvent builds a domain event.
func NewKeyCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventKeyCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyUpdatedEvent builds a domain event.
func NewKeyUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventKeyUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyDeletedEvent builds a domain event.
func NewKeyDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventKeyDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyArchivedEvent builds a domain event.
func NewKeyArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventKeyArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyRestoredEvent builds a domain event.
func NewKeyRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventKeyRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeySharedEvent builds a domain event.
func NewKeySharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventKeyShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyUnsharedEvent builds a domain event.
func NewKeyUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventKeyUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyExportedEvent builds a domain event.
func NewKeyExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventKeyExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyImportedEvent builds a domain event.
func NewKeyImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventKeyImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyRotatedEvent builds a domain event.
func NewKeyRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventKeyRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyAccessedEvent builds a domain event.
func NewKeyAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventKeyAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyDeniedEvent builds a domain event.
func NewKeyDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventKeyDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyApprovedEvent builds a domain event.
func NewKeyApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventKeyApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyRevokedEvent builds a domain event.
func NewKeyRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventKeyRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyExpiredEvent builds a domain event.
func NewKeyExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventKeyExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyLockedEvent builds a domain event.
func NewKeyLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventKeyLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyUnlockedEvent builds a domain event.
func NewKeyUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventKeyUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyVerifiedEvent builds a domain event.
func NewKeyVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventKeyVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeyFailedEvent builds a domain event.
func NewKeyFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventKeyFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewKeySucceededEvent builds a domain event.
func NewKeySucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Key"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventKeySucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownKeyEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

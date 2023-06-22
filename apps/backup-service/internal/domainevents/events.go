package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for backup-service.
const (
	EventSnapshotCreated = "backup.created"
	EventSnapshotUpdated = "backup.updated"
	EventSnapshotDeleted = "backup.deleted"
	EventSnapshotArchived = "backup.archived"
	EventSnapshotRestored = "backup.restored"
	EventSnapshotShared = "backup.shared"
	EventSnapshotUnshared = "backup.unshared"
	EventSnapshotExported = "backup.exported"
	EventSnapshotImported = "backup.imported"
	EventSnapshotRotated = "backup.rotated"
	EventSnapshotAccessed = "backup.accessed"
	EventSnapshotDenied = "backup.denied"
	EventSnapshotApproved = "backup.approved"
	EventSnapshotRevoked = "backup.revoked"
	EventSnapshotExpired = "backup.expired"
	EventSnapshotLocked = "backup.locked"
	EventSnapshotUnlocked = "backup.unlocked"
	EventSnapshotVerified = "backup.verified"
	EventSnapshotFailed = "backup.failed"
	EventSnapshotSucceeded = "backup.succeeded"
)

// Event is a backup-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownSnapshotEvent reports whether t is a known backup event type.
func IsKnownSnapshotEvent(t string) bool {
	switch t {
	case EventSnapshotCreated:
		return true
	case EventSnapshotUpdated:
		return true
	case EventSnapshotDeleted:
		return true
	case EventSnapshotArchived:
		return true
	case EventSnapshotRestored:
		return true
	case EventSnapshotShared:
		return true
	case EventSnapshotUnshared:
		return true
	case EventSnapshotExported:
		return true
	case EventSnapshotImported:
		return true
	case EventSnapshotRotated:
		return true
	case EventSnapshotAccessed:
		return true
	case EventSnapshotDenied:
		return true
	case EventSnapshotApproved:
		return true
	case EventSnapshotRevoked:
		return true
	case EventSnapshotExpired:
		return true
	case EventSnapshotLocked:
		return true
	case EventSnapshotUnlocked:
		return true
	case EventSnapshotVerified:
		return true
	case EventSnapshotFailed:
		return true
	case EventSnapshotSucceeded:
		return true
	default:
		return false
	}
}


// NewSnapshotCreatedEvent builds a domain event.
func NewSnapshotCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventSnapshotCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotUpdatedEvent builds a domain event.
func NewSnapshotUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventSnapshotUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotDeletedEvent builds a domain event.
func NewSnapshotDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventSnapshotDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotArchivedEvent builds a domain event.
func NewSnapshotArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventSnapshotArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotRestoredEvent builds a domain event.
func NewSnapshotRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventSnapshotRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotSharedEvent builds a domain event.
func NewSnapshotSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventSnapshotShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotUnsharedEvent builds a domain event.
func NewSnapshotUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventSnapshotUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotExportedEvent builds a domain event.
func NewSnapshotExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventSnapshotExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotImportedEvent builds a domain event.
func NewSnapshotImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventSnapshotImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotRotatedEvent builds a domain event.
func NewSnapshotRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventSnapshotRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotAccessedEvent builds a domain event.
func NewSnapshotAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventSnapshotAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotDeniedEvent builds a domain event.
func NewSnapshotDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventSnapshotDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotApprovedEvent builds a domain event.
func NewSnapshotApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventSnapshotApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotRevokedEvent builds a domain event.
func NewSnapshotRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventSnapshotRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotExpiredEvent builds a domain event.
func NewSnapshotExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventSnapshotExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotLockedEvent builds a domain event.
func NewSnapshotLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventSnapshotLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotUnlockedEvent builds a domain event.
func NewSnapshotUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventSnapshotUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotVerifiedEvent builds a domain event.
func NewSnapshotVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventSnapshotVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotFailedEvent builds a domain event.
func NewSnapshotFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventSnapshotFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSnapshotSucceededEvent builds a domain event.
func NewSnapshotSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Snapshot"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventSnapshotSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownSnapshotEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

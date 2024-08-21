package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for versioning-service.
const (
	EventVersionCreated = "versioning.created"
	EventVersionUpdated = "versioning.updated"
	EventVersionDeleted = "versioning.deleted"
	EventVersionArchived = "versioning.archived"
	EventVersionRestored = "versioning.restored"
	EventVersionShared = "versioning.shared"
	EventVersionUnshared = "versioning.unshared"
	EventVersionExported = "versioning.exported"
	EventVersionImported = "versioning.imported"
	EventVersionRotated = "versioning.rotated"
	EventVersionAccessed = "versioning.accessed"
	EventVersionDenied = "versioning.denied"
	EventVersionApproved = "versioning.approved"
	EventVersionRevoked = "versioning.revoked"
	EventVersionExpired = "versioning.expired"
	EventVersionLocked = "versioning.locked"
	EventVersionUnlocked = "versioning.unlocked"
	EventVersionVerified = "versioning.verified"
	EventVersionFailed = "versioning.failed"
	EventVersionSucceeded = "versioning.succeeded"
)

// Event is a versioning-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownVersionEvent reports whether t is a known versioning event type.
func IsKnownVersionEvent(t string) bool {
	switch t {
	case EventVersionCreated:
		return true
	case EventVersionUpdated:
		return true
	case EventVersionDeleted:
		return true
	case EventVersionArchived:
		return true
	case EventVersionRestored:
		return true
	case EventVersionShared:
		return true
	case EventVersionUnshared:
		return true
	case EventVersionExported:
		return true
	case EventVersionImported:
		return true
	case EventVersionRotated:
		return true
	case EventVersionAccessed:
		return true
	case EventVersionDenied:
		return true
	case EventVersionApproved:
		return true
	case EventVersionRevoked:
		return true
	case EventVersionExpired:
		return true
	case EventVersionLocked:
		return true
	case EventVersionUnlocked:
		return true
	case EventVersionVerified:
		return true
	case EventVersionFailed:
		return true
	case EventVersionSucceeded:
		return true
	default:
		return false
	}
}


// NewVersionCreatedEvent builds a domain event.
func NewVersionCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventVersionCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionUpdatedEvent builds a domain event.
func NewVersionUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventVersionUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionDeletedEvent builds a domain event.
func NewVersionDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventVersionDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionArchivedEvent builds a domain event.
func NewVersionArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventVersionArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionRestoredEvent builds a domain event.
func NewVersionRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventVersionRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionSharedEvent builds a domain event.
func NewVersionSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventVersionShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionUnsharedEvent builds a domain event.
func NewVersionUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventVersionUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionExportedEvent builds a domain event.
func NewVersionExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventVersionExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionImportedEvent builds a domain event.
func NewVersionImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventVersionImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionRotatedEvent builds a domain event.
func NewVersionRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventVersionRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionAccessedEvent builds a domain event.
func NewVersionAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventVersionAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionDeniedEvent builds a domain event.
func NewVersionDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventVersionDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionApprovedEvent builds a domain event.
func NewVersionApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventVersionApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionRevokedEvent builds a domain event.
func NewVersionRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventVersionRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionExpiredEvent builds a domain event.
func NewVersionExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventVersionExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionLockedEvent builds a domain event.
func NewVersionLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventVersionLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionUnlockedEvent builds a domain event.
func NewVersionUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventVersionUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionVerifiedEvent builds a domain event.
func NewVersionVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventVersionVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionFailedEvent builds a domain event.
func NewVersionFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventVersionFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVersionSucceededEvent builds a domain event.
func NewVersionSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Version"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventVersionSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownVersionEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

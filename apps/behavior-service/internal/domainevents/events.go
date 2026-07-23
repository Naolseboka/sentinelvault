package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for behavior-service.
const (
	EventSignalCreated   = "behavior.created"
	EventSignalUpdated   = "behavior.updated"
	EventSignalDeleted   = "behavior.deleted"
	EventSignalArchived  = "behavior.archived"
	EventSignalRestored  = "behavior.restored"
	EventSignalShared    = "behavior.shared"
	EventSignalUnshared  = "behavior.unshared"
	EventSignalExported  = "behavior.exported"
	EventSignalImported  = "behavior.imported"
	EventSignalRotated   = "behavior.rotated"
	EventSignalAccessed  = "behavior.accessed"
	EventSignalDenied    = "behavior.denied"
	EventSignalApproved  = "behavior.approved"
	EventSignalRevoked   = "behavior.revoked"
	EventSignalExpired   = "behavior.expired"
	EventSignalLocked    = "behavior.locked"
	EventSignalUnlocked  = "behavior.unlocked"
	EventSignalVerified  = "behavior.verified"
	EventSignalFailed    = "behavior.failed"
	EventSignalSucceeded = "behavior.succeeded"
)

// Event is a behavior-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownSignalEvent reports whether t is a known behavior event type.
func IsKnownSignalEvent(t string) bool {
	switch t {
	case EventSignalCreated:
		return true
	case EventSignalUpdated:
		return true
	case EventSignalDeleted:
		return true
	case EventSignalArchived:
		return true
	case EventSignalRestored:
		return true
	case EventSignalShared:
		return true
	case EventSignalUnshared:
		return true
	case EventSignalExported:
		return true
	case EventSignalImported:
		return true
	case EventSignalRotated:
		return true
	case EventSignalAccessed:
		return true
	case EventSignalDenied:
		return true
	case EventSignalApproved:
		return true
	case EventSignalRevoked:
		return true
	case EventSignalExpired:
		return true
	case EventSignalLocked:
		return true
	case EventSignalUnlocked:
		return true
	case EventSignalVerified:
		return true
	case EventSignalFailed:
		return true
	case EventSignalSucceeded:
		return true
	default:
		return false
	}
}

// NewSignalCreatedEvent builds a domain event.
func NewSignalCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventSignalCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalUpdatedEvent builds a domain event.
func NewSignalUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventSignalUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalDeletedEvent builds a domain event.
func NewSignalDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventSignalDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalArchivedEvent builds a domain event.
func NewSignalArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventSignalArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalRestoredEvent builds a domain event.
func NewSignalRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventSignalRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalSharedEvent builds a domain event.
func NewSignalSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventSignalShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalUnsharedEvent builds a domain event.
func NewSignalUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventSignalUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalExportedEvent builds a domain event.
func NewSignalExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventSignalExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalImportedEvent builds a domain event.
func NewSignalImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventSignalImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalRotatedEvent builds a domain event.
func NewSignalRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventSignalRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalAccessedEvent builds a domain event.
func NewSignalAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventSignalAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalDeniedEvent builds a domain event.
func NewSignalDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventSignalDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalApprovedEvent builds a domain event.
func NewSignalApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventSignalApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalRevokedEvent builds a domain event.
func NewSignalRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventSignalRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalExpiredEvent builds a domain event.
func NewSignalExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventSignalExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalLockedEvent builds a domain event.
func NewSignalLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventSignalLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalUnlockedEvent builds a domain event.
func NewSignalUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventSignalUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalVerifiedEvent builds a domain event.
func NewSignalVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventSignalVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalFailedEvent builds a domain event.
func NewSignalFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventSignalFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSignalSucceededEvent builds a domain event.
func NewSignalSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Signal"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventSignalSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownSignalEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

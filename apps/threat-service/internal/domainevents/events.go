package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for threat-service.
const (
	EventFindingCreated   = "threat.created"
	EventFindingUpdated   = "threat.updated"
	EventFindingDeleted   = "threat.deleted"
	EventFindingArchived  = "threat.archived"
	EventFindingRestored  = "threat.restored"
	EventFindingShared    = "threat.shared"
	EventFindingUnshared  = "threat.unshared"
	EventFindingExported  = "threat.exported"
	EventFindingImported  = "threat.imported"
	EventFindingRotated   = "threat.rotated"
	EventFindingAccessed  = "threat.accessed"
	EventFindingDenied    = "threat.denied"
	EventFindingApproved  = "threat.approved"
	EventFindingRevoked   = "threat.revoked"
	EventFindingExpired   = "threat.expired"
	EventFindingLocked    = "threat.locked"
	EventFindingUnlocked  = "threat.unlocked"
	EventFindingVerified  = "threat.verified"
	EventFindingFailed    = "threat.failed"
	EventFindingSucceeded = "threat.succeeded"
)

// Event is a threat-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownFindingEvent reports whether t is a known threat event type.
func IsKnownFindingEvent(t string) bool {
	switch t {
	case EventFindingCreated:
		return true
	case EventFindingUpdated:
		return true
	case EventFindingDeleted:
		return true
	case EventFindingArchived:
		return true
	case EventFindingRestored:
		return true
	case EventFindingShared:
		return true
	case EventFindingUnshared:
		return true
	case EventFindingExported:
		return true
	case EventFindingImported:
		return true
	case EventFindingRotated:
		return true
	case EventFindingAccessed:
		return true
	case EventFindingDenied:
		return true
	case EventFindingApproved:
		return true
	case EventFindingRevoked:
		return true
	case EventFindingExpired:
		return true
	case EventFindingLocked:
		return true
	case EventFindingUnlocked:
		return true
	case EventFindingVerified:
		return true
	case EventFindingFailed:
		return true
	case EventFindingSucceeded:
		return true
	default:
		return false
	}
}

// NewFindingCreatedEvent builds a domain event.
func NewFindingCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventFindingCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingUpdatedEvent builds a domain event.
func NewFindingUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventFindingUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingDeletedEvent builds a domain event.
func NewFindingDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventFindingDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingArchivedEvent builds a domain event.
func NewFindingArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventFindingArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingRestoredEvent builds a domain event.
func NewFindingRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventFindingRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingSharedEvent builds a domain event.
func NewFindingSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventFindingShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingUnsharedEvent builds a domain event.
func NewFindingUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventFindingUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingExportedEvent builds a domain event.
func NewFindingExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventFindingExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingImportedEvent builds a domain event.
func NewFindingImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventFindingImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingRotatedEvent builds a domain event.
func NewFindingRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventFindingRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingAccessedEvent builds a domain event.
func NewFindingAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventFindingAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingDeniedEvent builds a domain event.
func NewFindingDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventFindingDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingApprovedEvent builds a domain event.
func NewFindingApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventFindingApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingRevokedEvent builds a domain event.
func NewFindingRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventFindingRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingExpiredEvent builds a domain event.
func NewFindingExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventFindingExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingLockedEvent builds a domain event.
func NewFindingLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventFindingLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingUnlockedEvent builds a domain event.
func NewFindingUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventFindingUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingVerifiedEvent builds a domain event.
func NewFindingVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventFindingVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingFailedEvent builds a domain event.
func NewFindingFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventFindingFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewFindingSucceededEvent builds a domain event.
func NewFindingSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Finding"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventFindingSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownFindingEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

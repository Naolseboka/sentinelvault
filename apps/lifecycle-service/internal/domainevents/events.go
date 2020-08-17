package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for lifecycle-service.
const (
	EventLifecycleStateCreated = "lifecycle.created"
	EventLifecycleStateUpdated = "lifecycle.updated"
	EventLifecycleStateDeleted = "lifecycle.deleted"
	EventLifecycleStateArchived = "lifecycle.archived"
	EventLifecycleStateRestored = "lifecycle.restored"
	EventLifecycleStateShared = "lifecycle.shared"
	EventLifecycleStateUnshared = "lifecycle.unshared"
	EventLifecycleStateExported = "lifecycle.exported"
	EventLifecycleStateImported = "lifecycle.imported"
	EventLifecycleStateRotated = "lifecycle.rotated"
	EventLifecycleStateAccessed = "lifecycle.accessed"
	EventLifecycleStateDenied = "lifecycle.denied"
	EventLifecycleStateApproved = "lifecycle.approved"
	EventLifecycleStateRevoked = "lifecycle.revoked"
	EventLifecycleStateExpired = "lifecycle.expired"
	EventLifecycleStateLocked = "lifecycle.locked"
	EventLifecycleStateUnlocked = "lifecycle.unlocked"
	EventLifecycleStateVerified = "lifecycle.verified"
	EventLifecycleStateFailed = "lifecycle.failed"
	EventLifecycleStateSucceeded = "lifecycle.succeeded"
)

// Event is a lifecycle-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownLifecycleStateEvent reports whether t is a known lifecycle event type.
func IsKnownLifecycleStateEvent(t string) bool {
	switch t {
	case EventLifecycleStateCreated:
		return true
	case EventLifecycleStateUpdated:
		return true
	case EventLifecycleStateDeleted:
		return true
	case EventLifecycleStateArchived:
		return true
	case EventLifecycleStateRestored:
		return true
	case EventLifecycleStateShared:
		return true
	case EventLifecycleStateUnshared:
		return true
	case EventLifecycleStateExported:
		return true
	case EventLifecycleStateImported:
		return true
	case EventLifecycleStateRotated:
		return true
	case EventLifecycleStateAccessed:
		return true
	case EventLifecycleStateDenied:
		return true
	case EventLifecycleStateApproved:
		return true
	case EventLifecycleStateRevoked:
		return true
	case EventLifecycleStateExpired:
		return true
	case EventLifecycleStateLocked:
		return true
	case EventLifecycleStateUnlocked:
		return true
	case EventLifecycleStateVerified:
		return true
	case EventLifecycleStateFailed:
		return true
	case EventLifecycleStateSucceeded:
		return true
	default:
		return false
	}
}


// NewLifecycleStateCreatedEvent builds a domain event.
func NewLifecycleStateCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateUpdatedEvent builds a domain event.
func NewLifecycleStateUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateDeletedEvent builds a domain event.
func NewLifecycleStateDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateArchivedEvent builds a domain event.
func NewLifecycleStateArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateRestoredEvent builds a domain event.
func NewLifecycleStateRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateSharedEvent builds a domain event.
func NewLifecycleStateSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateUnsharedEvent builds a domain event.
func NewLifecycleStateUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateExportedEvent builds a domain event.
func NewLifecycleStateExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateImportedEvent builds a domain event.
func NewLifecycleStateImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateRotatedEvent builds a domain event.
func NewLifecycleStateRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateAccessedEvent builds a domain event.
func NewLifecycleStateAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateDeniedEvent builds a domain event.
func NewLifecycleStateDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateApprovedEvent builds a domain event.
func NewLifecycleStateApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateRevokedEvent builds a domain event.
func NewLifecycleStateRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateExpiredEvent builds a domain event.
func NewLifecycleStateExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateLockedEvent builds a domain event.
func NewLifecycleStateLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateUnlockedEvent builds a domain event.
func NewLifecycleStateUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateVerifiedEvent builds a domain event.
func NewLifecycleStateVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateFailedEvent builds a domain event.
func NewLifecycleStateFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLifecycleStateSucceededEvent builds a domain event.
func NewLifecycleStateSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "LifecycleState"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventLifecycleStateSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownLifecycleStateEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

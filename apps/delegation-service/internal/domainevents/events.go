package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for delegation-service.
const (
	EventDelegationCreated = "delegation.created"
	EventDelegationUpdated = "delegation.updated"
	EventDelegationDeleted = "delegation.deleted"
	EventDelegationArchived = "delegation.archived"
	EventDelegationRestored = "delegation.restored"
	EventDelegationShared = "delegation.shared"
	EventDelegationUnshared = "delegation.unshared"
	EventDelegationExported = "delegation.exported"
	EventDelegationImported = "delegation.imported"
	EventDelegationRotated = "delegation.rotated"
	EventDelegationAccessed = "delegation.accessed"
	EventDelegationDenied = "delegation.denied"
	EventDelegationApproved = "delegation.approved"
	EventDelegationRevoked = "delegation.revoked"
	EventDelegationExpired = "delegation.expired"
	EventDelegationLocked = "delegation.locked"
	EventDelegationUnlocked = "delegation.unlocked"
	EventDelegationVerified = "delegation.verified"
	EventDelegationFailed = "delegation.failed"
	EventDelegationSucceeded = "delegation.succeeded"
)

// Event is a delegation-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownDelegationEvent reports whether t is a known delegation event type.
func IsKnownDelegationEvent(t string) bool {
	switch t {
	case EventDelegationCreated:
		return true
	case EventDelegationUpdated:
		return true
	case EventDelegationDeleted:
		return true
	case EventDelegationArchived:
		return true
	case EventDelegationRestored:
		return true
	case EventDelegationShared:
		return true
	case EventDelegationUnshared:
		return true
	case EventDelegationExported:
		return true
	case EventDelegationImported:
		return true
	case EventDelegationRotated:
		return true
	case EventDelegationAccessed:
		return true
	case EventDelegationDenied:
		return true
	case EventDelegationApproved:
		return true
	case EventDelegationRevoked:
		return true
	case EventDelegationExpired:
		return true
	case EventDelegationLocked:
		return true
	case EventDelegationUnlocked:
		return true
	case EventDelegationVerified:
		return true
	case EventDelegationFailed:
		return true
	case EventDelegationSucceeded:
		return true
	default:
		return false
	}
}


// NewDelegationCreatedEvent builds a domain event.
func NewDelegationCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventDelegationCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationUpdatedEvent builds a domain event.
func NewDelegationUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventDelegationUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationDeletedEvent builds a domain event.
func NewDelegationDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventDelegationDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationArchivedEvent builds a domain event.
func NewDelegationArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventDelegationArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationRestoredEvent builds a domain event.
func NewDelegationRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventDelegationRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationSharedEvent builds a domain event.
func NewDelegationSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventDelegationShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationUnsharedEvent builds a domain event.
func NewDelegationUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventDelegationUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationExportedEvent builds a domain event.
func NewDelegationExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventDelegationExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationImportedEvent builds a domain event.
func NewDelegationImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventDelegationImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationRotatedEvent builds a domain event.
func NewDelegationRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventDelegationRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationAccessedEvent builds a domain event.
func NewDelegationAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventDelegationAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationDeniedEvent builds a domain event.
func NewDelegationDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventDelegationDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationApprovedEvent builds a domain event.
func NewDelegationApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventDelegationApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationRevokedEvent builds a domain event.
func NewDelegationRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventDelegationRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationExpiredEvent builds a domain event.
func NewDelegationExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventDelegationExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationLockedEvent builds a domain event.
func NewDelegationLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventDelegationLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationUnlockedEvent builds a domain event.
func NewDelegationUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventDelegationUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationVerifiedEvent builds a domain event.
func NewDelegationVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventDelegationVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationFailedEvent builds a domain event.
func NewDelegationFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventDelegationFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDelegationSucceededEvent builds a domain event.
func NewDelegationSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Delegation"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventDelegationSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownDelegationEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for rbac-service.
const (
	EventRoleBindingCreated = "rbac.created"
	EventRoleBindingUpdated = "rbac.updated"
	EventRoleBindingDeleted = "rbac.deleted"
	EventRoleBindingArchived = "rbac.archived"
	EventRoleBindingRestored = "rbac.restored"
	EventRoleBindingShared = "rbac.shared"
	EventRoleBindingUnshared = "rbac.unshared"
	EventRoleBindingExported = "rbac.exported"
	EventRoleBindingImported = "rbac.imported"
	EventRoleBindingRotated = "rbac.rotated"
	EventRoleBindingAccessed = "rbac.accessed"
	EventRoleBindingDenied = "rbac.denied"
	EventRoleBindingApproved = "rbac.approved"
	EventRoleBindingRevoked = "rbac.revoked"
	EventRoleBindingExpired = "rbac.expired"
	EventRoleBindingLocked = "rbac.locked"
	EventRoleBindingUnlocked = "rbac.unlocked"
	EventRoleBindingVerified = "rbac.verified"
	EventRoleBindingFailed = "rbac.failed"
	EventRoleBindingSucceeded = "rbac.succeeded"
)

// Event is a rbac-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownRoleBindingEvent reports whether t is a known rbac event type.
func IsKnownRoleBindingEvent(t string) bool {
	switch t {
	case EventRoleBindingCreated:
		return true
	case EventRoleBindingUpdated:
		return true
	case EventRoleBindingDeleted:
		return true
	case EventRoleBindingArchived:
		return true
	case EventRoleBindingRestored:
		return true
	case EventRoleBindingShared:
		return true
	case EventRoleBindingUnshared:
		return true
	case EventRoleBindingExported:
		return true
	case EventRoleBindingImported:
		return true
	case EventRoleBindingRotated:
		return true
	case EventRoleBindingAccessed:
		return true
	case EventRoleBindingDenied:
		return true
	case EventRoleBindingApproved:
		return true
	case EventRoleBindingRevoked:
		return true
	case EventRoleBindingExpired:
		return true
	case EventRoleBindingLocked:
		return true
	case EventRoleBindingUnlocked:
		return true
	case EventRoleBindingVerified:
		return true
	case EventRoleBindingFailed:
		return true
	case EventRoleBindingSucceeded:
		return true
	default:
		return false
	}
}


// NewRoleBindingCreatedEvent builds a domain event.
func NewRoleBindingCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingUpdatedEvent builds a domain event.
func NewRoleBindingUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingDeletedEvent builds a domain event.
func NewRoleBindingDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingArchivedEvent builds a domain event.
func NewRoleBindingArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingRestoredEvent builds a domain event.
func NewRoleBindingRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingSharedEvent builds a domain event.
func NewRoleBindingSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingUnsharedEvent builds a domain event.
func NewRoleBindingUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingExportedEvent builds a domain event.
func NewRoleBindingExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingImportedEvent builds a domain event.
func NewRoleBindingImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingRotatedEvent builds a domain event.
func NewRoleBindingRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingAccessedEvent builds a domain event.
func NewRoleBindingAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingDeniedEvent builds a domain event.
func NewRoleBindingDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingApprovedEvent builds a domain event.
func NewRoleBindingApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingRevokedEvent builds a domain event.
func NewRoleBindingRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingExpiredEvent builds a domain event.
func NewRoleBindingExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingLockedEvent builds a domain event.
func NewRoleBindingLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingUnlockedEvent builds a domain event.
func NewRoleBindingUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingVerifiedEvent builds a domain event.
func NewRoleBindingVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingFailedEvent builds a domain event.
func NewRoleBindingFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRoleBindingSucceededEvent builds a domain event.
func NewRoleBindingSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RoleBinding"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventRoleBindingSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownRoleBindingEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

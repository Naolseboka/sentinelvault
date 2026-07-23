package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for admin-service.
const (
	EventAdminActionCreated   = "admin.created"
	EventAdminActionUpdated   = "admin.updated"
	EventAdminActionDeleted   = "admin.deleted"
	EventAdminActionArchived  = "admin.archived"
	EventAdminActionRestored  = "admin.restored"
	EventAdminActionShared    = "admin.shared"
	EventAdminActionUnshared  = "admin.unshared"
	EventAdminActionExported  = "admin.exported"
	EventAdminActionImported  = "admin.imported"
	EventAdminActionRotated   = "admin.rotated"
	EventAdminActionAccessed  = "admin.accessed"
	EventAdminActionDenied    = "admin.denied"
	EventAdminActionApproved  = "admin.approved"
	EventAdminActionRevoked   = "admin.revoked"
	EventAdminActionExpired   = "admin.expired"
	EventAdminActionLocked    = "admin.locked"
	EventAdminActionUnlocked  = "admin.unlocked"
	EventAdminActionVerified  = "admin.verified"
	EventAdminActionFailed    = "admin.failed"
	EventAdminActionSucceeded = "admin.succeeded"
)

// Event is a admin-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownAdminActionEvent reports whether t is a known admin event type.
func IsKnownAdminActionEvent(t string) bool {
	switch t {
	case EventAdminActionCreated:
		return true
	case EventAdminActionUpdated:
		return true
	case EventAdminActionDeleted:
		return true
	case EventAdminActionArchived:
		return true
	case EventAdminActionRestored:
		return true
	case EventAdminActionShared:
		return true
	case EventAdminActionUnshared:
		return true
	case EventAdminActionExported:
		return true
	case EventAdminActionImported:
		return true
	case EventAdminActionRotated:
		return true
	case EventAdminActionAccessed:
		return true
	case EventAdminActionDenied:
		return true
	case EventAdminActionApproved:
		return true
	case EventAdminActionRevoked:
		return true
	case EventAdminActionExpired:
		return true
	case EventAdminActionLocked:
		return true
	case EventAdminActionUnlocked:
		return true
	case EventAdminActionVerified:
		return true
	case EventAdminActionFailed:
		return true
	case EventAdminActionSucceeded:
		return true
	default:
		return false
	}
}

// NewAdminActionCreatedEvent builds a domain event.
func NewAdminActionCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventAdminActionCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionUpdatedEvent builds a domain event.
func NewAdminActionUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventAdminActionUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionDeletedEvent builds a domain event.
func NewAdminActionDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventAdminActionDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionArchivedEvent builds a domain event.
func NewAdminActionArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventAdminActionArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionRestoredEvent builds a domain event.
func NewAdminActionRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventAdminActionRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionSharedEvent builds a domain event.
func NewAdminActionSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventAdminActionShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionUnsharedEvent builds a domain event.
func NewAdminActionUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventAdminActionUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionExportedEvent builds a domain event.
func NewAdminActionExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventAdminActionExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionImportedEvent builds a domain event.
func NewAdminActionImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventAdminActionImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionRotatedEvent builds a domain event.
func NewAdminActionRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventAdminActionRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionAccessedEvent builds a domain event.
func NewAdminActionAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventAdminActionAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionDeniedEvent builds a domain event.
func NewAdminActionDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventAdminActionDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionApprovedEvent builds a domain event.
func NewAdminActionApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventAdminActionApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionRevokedEvent builds a domain event.
func NewAdminActionRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventAdminActionRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionExpiredEvent builds a domain event.
func NewAdminActionExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventAdminActionExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionLockedEvent builds a domain event.
func NewAdminActionLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventAdminActionLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionUnlockedEvent builds a domain event.
func NewAdminActionUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventAdminActionUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionVerifiedEvent builds a domain event.
func NewAdminActionVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventAdminActionVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionFailedEvent builds a domain event.
func NewAdminActionFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventAdminActionFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAdminActionSucceededEvent builds a domain event.
func NewAdminActionSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AdminAction"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventAdminActionSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownAdminActionEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

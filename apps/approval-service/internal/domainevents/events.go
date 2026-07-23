package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for approval-service.
const (
	EventApprovalCreated   = "approval.created"
	EventApprovalUpdated   = "approval.updated"
	EventApprovalDeleted   = "approval.deleted"
	EventApprovalArchived  = "approval.archived"
	EventApprovalRestored  = "approval.restored"
	EventApprovalShared    = "approval.shared"
	EventApprovalUnshared  = "approval.unshared"
	EventApprovalExported  = "approval.exported"
	EventApprovalImported  = "approval.imported"
	EventApprovalRotated   = "approval.rotated"
	EventApprovalAccessed  = "approval.accessed"
	EventApprovalDenied    = "approval.denied"
	EventApprovalApproved  = "approval.approved"
	EventApprovalRevoked   = "approval.revoked"
	EventApprovalExpired   = "approval.expired"
	EventApprovalLocked    = "approval.locked"
	EventApprovalUnlocked  = "approval.unlocked"
	EventApprovalVerified  = "approval.verified"
	EventApprovalFailed    = "approval.failed"
	EventApprovalSucceeded = "approval.succeeded"
)

// Event is a approval-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownApprovalEvent reports whether t is a known approval event type.
func IsKnownApprovalEvent(t string) bool {
	switch t {
	case EventApprovalCreated:
		return true
	case EventApprovalUpdated:
		return true
	case EventApprovalDeleted:
		return true
	case EventApprovalArchived:
		return true
	case EventApprovalRestored:
		return true
	case EventApprovalShared:
		return true
	case EventApprovalUnshared:
		return true
	case EventApprovalExported:
		return true
	case EventApprovalImported:
		return true
	case EventApprovalRotated:
		return true
	case EventApprovalAccessed:
		return true
	case EventApprovalDenied:
		return true
	case EventApprovalApproved:
		return true
	case EventApprovalRevoked:
		return true
	case EventApprovalExpired:
		return true
	case EventApprovalLocked:
		return true
	case EventApprovalUnlocked:
		return true
	case EventApprovalVerified:
		return true
	case EventApprovalFailed:
		return true
	case EventApprovalSucceeded:
		return true
	default:
		return false
	}
}

// NewApprovalCreatedEvent builds a domain event.
func NewApprovalCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventApprovalCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalUpdatedEvent builds a domain event.
func NewApprovalUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventApprovalUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalDeletedEvent builds a domain event.
func NewApprovalDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventApprovalDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalArchivedEvent builds a domain event.
func NewApprovalArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventApprovalArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalRestoredEvent builds a domain event.
func NewApprovalRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventApprovalRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalSharedEvent builds a domain event.
func NewApprovalSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventApprovalShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalUnsharedEvent builds a domain event.
func NewApprovalUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventApprovalUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalExportedEvent builds a domain event.
func NewApprovalExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventApprovalExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalImportedEvent builds a domain event.
func NewApprovalImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventApprovalImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalRotatedEvent builds a domain event.
func NewApprovalRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventApprovalRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalAccessedEvent builds a domain event.
func NewApprovalAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventApprovalAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalDeniedEvent builds a domain event.
func NewApprovalDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventApprovalDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalApprovedEvent builds a domain event.
func NewApprovalApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventApprovalApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalRevokedEvent builds a domain event.
func NewApprovalRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventApprovalRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalExpiredEvent builds a domain event.
func NewApprovalExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventApprovalExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalLockedEvent builds a domain event.
func NewApprovalLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventApprovalLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalUnlockedEvent builds a domain event.
func NewApprovalUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventApprovalUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalVerifiedEvent builds a domain event.
func NewApprovalVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventApprovalVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalFailedEvent builds a domain event.
func NewApprovalFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventApprovalFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewApprovalSucceededEvent builds a domain event.
func NewApprovalSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Approval"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventApprovalSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownApprovalEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for invitation-service.
const (
	EventInvitationCreated   = "invitation.created"
	EventInvitationUpdated   = "invitation.updated"
	EventInvitationDeleted   = "invitation.deleted"
	EventInvitationArchived  = "invitation.archived"
	EventInvitationRestored  = "invitation.restored"
	EventInvitationShared    = "invitation.shared"
	EventInvitationUnshared  = "invitation.unshared"
	EventInvitationExported  = "invitation.exported"
	EventInvitationImported  = "invitation.imported"
	EventInvitationRotated   = "invitation.rotated"
	EventInvitationAccessed  = "invitation.accessed"
	EventInvitationDenied    = "invitation.denied"
	EventInvitationApproved  = "invitation.approved"
	EventInvitationRevoked   = "invitation.revoked"
	EventInvitationExpired   = "invitation.expired"
	EventInvitationLocked    = "invitation.locked"
	EventInvitationUnlocked  = "invitation.unlocked"
	EventInvitationVerified  = "invitation.verified"
	EventInvitationFailed    = "invitation.failed"
	EventInvitationSucceeded = "invitation.succeeded"
)

// Event is a invitation-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownInvitationEvent reports whether t is a known invitation event type.
func IsKnownInvitationEvent(t string) bool {
	switch t {
	case EventInvitationCreated:
		return true
	case EventInvitationUpdated:
		return true
	case EventInvitationDeleted:
		return true
	case EventInvitationArchived:
		return true
	case EventInvitationRestored:
		return true
	case EventInvitationShared:
		return true
	case EventInvitationUnshared:
		return true
	case EventInvitationExported:
		return true
	case EventInvitationImported:
		return true
	case EventInvitationRotated:
		return true
	case EventInvitationAccessed:
		return true
	case EventInvitationDenied:
		return true
	case EventInvitationApproved:
		return true
	case EventInvitationRevoked:
		return true
	case EventInvitationExpired:
		return true
	case EventInvitationLocked:
		return true
	case EventInvitationUnlocked:
		return true
	case EventInvitationVerified:
		return true
	case EventInvitationFailed:
		return true
	case EventInvitationSucceeded:
		return true
	default:
		return false
	}
}

// NewInvitationCreatedEvent builds a domain event.
func NewInvitationCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventInvitationCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationUpdatedEvent builds a domain event.
func NewInvitationUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventInvitationUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationDeletedEvent builds a domain event.
func NewInvitationDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventInvitationDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationArchivedEvent builds a domain event.
func NewInvitationArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventInvitationArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationRestoredEvent builds a domain event.
func NewInvitationRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventInvitationRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationSharedEvent builds a domain event.
func NewInvitationSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventInvitationShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationUnsharedEvent builds a domain event.
func NewInvitationUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventInvitationUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationExportedEvent builds a domain event.
func NewInvitationExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventInvitationExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationImportedEvent builds a domain event.
func NewInvitationImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventInvitationImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationRotatedEvent builds a domain event.
func NewInvitationRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventInvitationRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationAccessedEvent builds a domain event.
func NewInvitationAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventInvitationAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationDeniedEvent builds a domain event.
func NewInvitationDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventInvitationDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationApprovedEvent builds a domain event.
func NewInvitationApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventInvitationApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationRevokedEvent builds a domain event.
func NewInvitationRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventInvitationRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationExpiredEvent builds a domain event.
func NewInvitationExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventInvitationExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationLockedEvent builds a domain event.
func NewInvitationLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventInvitationLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationUnlockedEvent builds a domain event.
func NewInvitationUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventInvitationUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationVerifiedEvent builds a domain event.
func NewInvitationVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventInvitationVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationFailedEvent builds a domain event.
func NewInvitationFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventInvitationFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewInvitationSucceededEvent builds a domain event.
func NewInvitationSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Invitation"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventInvitationSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownInvitationEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

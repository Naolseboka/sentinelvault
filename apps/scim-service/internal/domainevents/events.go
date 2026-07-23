package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for scim-service.
const (
	EventSCIMUserCreated   = "scim.created"
	EventSCIMUserUpdated   = "scim.updated"
	EventSCIMUserDeleted   = "scim.deleted"
	EventSCIMUserArchived  = "scim.archived"
	EventSCIMUserRestored  = "scim.restored"
	EventSCIMUserShared    = "scim.shared"
	EventSCIMUserUnshared  = "scim.unshared"
	EventSCIMUserExported  = "scim.exported"
	EventSCIMUserImported  = "scim.imported"
	EventSCIMUserRotated   = "scim.rotated"
	EventSCIMUserAccessed  = "scim.accessed"
	EventSCIMUserDenied    = "scim.denied"
	EventSCIMUserApproved  = "scim.approved"
	EventSCIMUserRevoked   = "scim.revoked"
	EventSCIMUserExpired   = "scim.expired"
	EventSCIMUserLocked    = "scim.locked"
	EventSCIMUserUnlocked  = "scim.unlocked"
	EventSCIMUserVerified  = "scim.verified"
	EventSCIMUserFailed    = "scim.failed"
	EventSCIMUserSucceeded = "scim.succeeded"
)

// Event is a scim-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownSCIMUserEvent reports whether t is a known scim event type.
func IsKnownSCIMUserEvent(t string) bool {
	switch t {
	case EventSCIMUserCreated:
		return true
	case EventSCIMUserUpdated:
		return true
	case EventSCIMUserDeleted:
		return true
	case EventSCIMUserArchived:
		return true
	case EventSCIMUserRestored:
		return true
	case EventSCIMUserShared:
		return true
	case EventSCIMUserUnshared:
		return true
	case EventSCIMUserExported:
		return true
	case EventSCIMUserImported:
		return true
	case EventSCIMUserRotated:
		return true
	case EventSCIMUserAccessed:
		return true
	case EventSCIMUserDenied:
		return true
	case EventSCIMUserApproved:
		return true
	case EventSCIMUserRevoked:
		return true
	case EventSCIMUserExpired:
		return true
	case EventSCIMUserLocked:
		return true
	case EventSCIMUserUnlocked:
		return true
	case EventSCIMUserVerified:
		return true
	case EventSCIMUserFailed:
		return true
	case EventSCIMUserSucceeded:
		return true
	default:
		return false
	}
}

// NewSCIMUserCreatedEvent builds a domain event.
func NewSCIMUserCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserUpdatedEvent builds a domain event.
func NewSCIMUserUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserDeletedEvent builds a domain event.
func NewSCIMUserDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserArchivedEvent builds a domain event.
func NewSCIMUserArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserRestoredEvent builds a domain event.
func NewSCIMUserRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserSharedEvent builds a domain event.
func NewSCIMUserSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserUnsharedEvent builds a domain event.
func NewSCIMUserUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserExportedEvent builds a domain event.
func NewSCIMUserExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserImportedEvent builds a domain event.
func NewSCIMUserImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserRotatedEvent builds a domain event.
func NewSCIMUserRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserAccessedEvent builds a domain event.
func NewSCIMUserAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserDeniedEvent builds a domain event.
func NewSCIMUserDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserApprovedEvent builds a domain event.
func NewSCIMUserApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserRevokedEvent builds a domain event.
func NewSCIMUserRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserExpiredEvent builds a domain event.
func NewSCIMUserExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserLockedEvent builds a domain event.
func NewSCIMUserLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserUnlockedEvent builds a domain event.
func NewSCIMUserUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserVerifiedEvent builds a domain event.
func NewSCIMUserVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserFailedEvent builds a domain event.
func NewSCIMUserFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSCIMUserSucceededEvent builds a domain event.
func NewSCIMUserSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SCIMUser"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventSCIMUserSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownSCIMUserEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

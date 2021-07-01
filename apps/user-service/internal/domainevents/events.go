package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for user-service.
const (
	EventUserCreated = "user.created"
	EventUserUpdated = "user.updated"
	EventUserDeleted = "user.deleted"
	EventUserArchived = "user.archived"
	EventUserRestored = "user.restored"
	EventUserShared = "user.shared"
	EventUserUnshared = "user.unshared"
	EventUserExported = "user.exported"
	EventUserImported = "user.imported"
	EventUserRotated = "user.rotated"
	EventUserAccessed = "user.accessed"
	EventUserDenied = "user.denied"
	EventUserApproved = "user.approved"
	EventUserRevoked = "user.revoked"
	EventUserExpired = "user.expired"
	EventUserLocked = "user.locked"
	EventUserUnlocked = "user.unlocked"
	EventUserVerified = "user.verified"
	EventUserFailed = "user.failed"
	EventUserSucceeded = "user.succeeded"
)

// Event is a user-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownUserEvent reports whether t is a known user event type.
func IsKnownUserEvent(t string) bool {
	switch t {
	case EventUserCreated:
		return true
	case EventUserUpdated:
		return true
	case EventUserDeleted:
		return true
	case EventUserArchived:
		return true
	case EventUserRestored:
		return true
	case EventUserShared:
		return true
	case EventUserUnshared:
		return true
	case EventUserExported:
		return true
	case EventUserImported:
		return true
	case EventUserRotated:
		return true
	case EventUserAccessed:
		return true
	case EventUserDenied:
		return true
	case EventUserApproved:
		return true
	case EventUserRevoked:
		return true
	case EventUserExpired:
		return true
	case EventUserLocked:
		return true
	case EventUserUnlocked:
		return true
	case EventUserVerified:
		return true
	case EventUserFailed:
		return true
	case EventUserSucceeded:
		return true
	default:
		return false
	}
}


// NewUserCreatedEvent builds a domain event.
func NewUserCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventUserCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserUpdatedEvent builds a domain event.
func NewUserUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventUserUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserDeletedEvent builds a domain event.
func NewUserDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventUserDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserArchivedEvent builds a domain event.
func NewUserArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventUserArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserRestoredEvent builds a domain event.
func NewUserRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventUserRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserSharedEvent builds a domain event.
func NewUserSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventUserShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserUnsharedEvent builds a domain event.
func NewUserUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventUserUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserExportedEvent builds a domain event.
func NewUserExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventUserExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserImportedEvent builds a domain event.
func NewUserImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventUserImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserRotatedEvent builds a domain event.
func NewUserRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventUserRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserAccessedEvent builds a domain event.
func NewUserAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventUserAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserDeniedEvent builds a domain event.
func NewUserDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventUserDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserApprovedEvent builds a domain event.
func NewUserApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventUserApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserRevokedEvent builds a domain event.
func NewUserRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventUserRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserExpiredEvent builds a domain event.
func NewUserExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventUserExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserLockedEvent builds a domain event.
func NewUserLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventUserLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserUnlockedEvent builds a domain event.
func NewUserUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventUserUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserVerifiedEvent builds a domain event.
func NewUserVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventUserVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserFailedEvent builds a domain event.
func NewUserFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventUserFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewUserSucceededEvent builds a domain event.
func NewUserSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "User"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventUserSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownUserEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

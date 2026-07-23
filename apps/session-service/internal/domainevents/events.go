package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for session-service.
const (
	EventSessionCreated   = "session.created"
	EventSessionUpdated   = "session.updated"
	EventSessionDeleted   = "session.deleted"
	EventSessionArchived  = "session.archived"
	EventSessionRestored  = "session.restored"
	EventSessionShared    = "session.shared"
	EventSessionUnshared  = "session.unshared"
	EventSessionExported  = "session.exported"
	EventSessionImported  = "session.imported"
	EventSessionRotated   = "session.rotated"
	EventSessionAccessed  = "session.accessed"
	EventSessionDenied    = "session.denied"
	EventSessionApproved  = "session.approved"
	EventSessionRevoked   = "session.revoked"
	EventSessionExpired   = "session.expired"
	EventSessionLocked    = "session.locked"
	EventSessionUnlocked  = "session.unlocked"
	EventSessionVerified  = "session.verified"
	EventSessionFailed    = "session.failed"
	EventSessionSucceeded = "session.succeeded"
)

// Event is a session-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownSessionEvent reports whether t is a known session event type.
func IsKnownSessionEvent(t string) bool {
	switch t {
	case EventSessionCreated:
		return true
	case EventSessionUpdated:
		return true
	case EventSessionDeleted:
		return true
	case EventSessionArchived:
		return true
	case EventSessionRestored:
		return true
	case EventSessionShared:
		return true
	case EventSessionUnshared:
		return true
	case EventSessionExported:
		return true
	case EventSessionImported:
		return true
	case EventSessionRotated:
		return true
	case EventSessionAccessed:
		return true
	case EventSessionDenied:
		return true
	case EventSessionApproved:
		return true
	case EventSessionRevoked:
		return true
	case EventSessionExpired:
		return true
	case EventSessionLocked:
		return true
	case EventSessionUnlocked:
		return true
	case EventSessionVerified:
		return true
	case EventSessionFailed:
		return true
	case EventSessionSucceeded:
		return true
	default:
		return false
	}
}

// NewSessionCreatedEvent builds a domain event.
func NewSessionCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventSessionCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionUpdatedEvent builds a domain event.
func NewSessionUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventSessionUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionDeletedEvent builds a domain event.
func NewSessionDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventSessionDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionArchivedEvent builds a domain event.
func NewSessionArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventSessionArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionRestoredEvent builds a domain event.
func NewSessionRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventSessionRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionSharedEvent builds a domain event.
func NewSessionSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventSessionShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionUnsharedEvent builds a domain event.
func NewSessionUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventSessionUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionExportedEvent builds a domain event.
func NewSessionExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventSessionExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionImportedEvent builds a domain event.
func NewSessionImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventSessionImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionRotatedEvent builds a domain event.
func NewSessionRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventSessionRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionAccessedEvent builds a domain event.
func NewSessionAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventSessionAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionDeniedEvent builds a domain event.
func NewSessionDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventSessionDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionApprovedEvent builds a domain event.
func NewSessionApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventSessionApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionRevokedEvent builds a domain event.
func NewSessionRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventSessionRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionExpiredEvent builds a domain event.
func NewSessionExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventSessionExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionLockedEvent builds a domain event.
func NewSessionLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventSessionLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionUnlockedEvent builds a domain event.
func NewSessionUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventSessionUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionVerifiedEvent builds a domain event.
func NewSessionVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventSessionVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionFailedEvent builds a domain event.
func NewSessionFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventSessionFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSessionSucceededEvent builds a domain event.
func NewSessionSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Session"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventSessionSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownSessionEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for passkey-service.
const (
	EventPasskeyCreated   = "passkey.created"
	EventPasskeyUpdated   = "passkey.updated"
	EventPasskeyDeleted   = "passkey.deleted"
	EventPasskeyArchived  = "passkey.archived"
	EventPasskeyRestored  = "passkey.restored"
	EventPasskeyShared    = "passkey.shared"
	EventPasskeyUnshared  = "passkey.unshared"
	EventPasskeyExported  = "passkey.exported"
	EventPasskeyImported  = "passkey.imported"
	EventPasskeyRotated   = "passkey.rotated"
	EventPasskeyAccessed  = "passkey.accessed"
	EventPasskeyDenied    = "passkey.denied"
	EventPasskeyApproved  = "passkey.approved"
	EventPasskeyRevoked   = "passkey.revoked"
	EventPasskeyExpired   = "passkey.expired"
	EventPasskeyLocked    = "passkey.locked"
	EventPasskeyUnlocked  = "passkey.unlocked"
	EventPasskeyVerified  = "passkey.verified"
	EventPasskeyFailed    = "passkey.failed"
	EventPasskeySucceeded = "passkey.succeeded"
)

// Event is a passkey-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownPasskeyEvent reports whether t is a known passkey event type.
func IsKnownPasskeyEvent(t string) bool {
	switch t {
	case EventPasskeyCreated:
		return true
	case EventPasskeyUpdated:
		return true
	case EventPasskeyDeleted:
		return true
	case EventPasskeyArchived:
		return true
	case EventPasskeyRestored:
		return true
	case EventPasskeyShared:
		return true
	case EventPasskeyUnshared:
		return true
	case EventPasskeyExported:
		return true
	case EventPasskeyImported:
		return true
	case EventPasskeyRotated:
		return true
	case EventPasskeyAccessed:
		return true
	case EventPasskeyDenied:
		return true
	case EventPasskeyApproved:
		return true
	case EventPasskeyRevoked:
		return true
	case EventPasskeyExpired:
		return true
	case EventPasskeyLocked:
		return true
	case EventPasskeyUnlocked:
		return true
	case EventPasskeyVerified:
		return true
	case EventPasskeyFailed:
		return true
	case EventPasskeySucceeded:
		return true
	default:
		return false
	}
}

// NewPasskeyCreatedEvent builds a domain event.
func NewPasskeyCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventPasskeyCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyUpdatedEvent builds a domain event.
func NewPasskeyUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventPasskeyUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyDeletedEvent builds a domain event.
func NewPasskeyDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventPasskeyDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyArchivedEvent builds a domain event.
func NewPasskeyArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventPasskeyArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyRestoredEvent builds a domain event.
func NewPasskeyRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventPasskeyRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeySharedEvent builds a domain event.
func NewPasskeySharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventPasskeyShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyUnsharedEvent builds a domain event.
func NewPasskeyUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventPasskeyUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyExportedEvent builds a domain event.
func NewPasskeyExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventPasskeyExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyImportedEvent builds a domain event.
func NewPasskeyImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventPasskeyImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyRotatedEvent builds a domain event.
func NewPasskeyRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventPasskeyRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyAccessedEvent builds a domain event.
func NewPasskeyAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventPasskeyAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyDeniedEvent builds a domain event.
func NewPasskeyDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventPasskeyDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyApprovedEvent builds a domain event.
func NewPasskeyApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventPasskeyApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyRevokedEvent builds a domain event.
func NewPasskeyRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventPasskeyRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyExpiredEvent builds a domain event.
func NewPasskeyExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventPasskeyExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyLockedEvent builds a domain event.
func NewPasskeyLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventPasskeyLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyUnlockedEvent builds a domain event.
func NewPasskeyUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventPasskeyUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyVerifiedEvent builds a domain event.
func NewPasskeyVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventPasskeyVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeyFailedEvent builds a domain event.
func NewPasskeyFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventPasskeyFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPasskeySucceededEvent builds a domain event.
func NewPasskeySucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Passkey"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventPasskeySucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownPasskeyEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

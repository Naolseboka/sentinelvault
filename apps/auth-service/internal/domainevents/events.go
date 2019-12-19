package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for auth-service.
const (
	EventCredentialCreated = "auth.created"
	EventCredentialUpdated = "auth.updated"
	EventCredentialDeleted = "auth.deleted"
	EventCredentialArchived = "auth.archived"
	EventCredentialRestored = "auth.restored"
	EventCredentialShared = "auth.shared"
	EventCredentialUnshared = "auth.unshared"
	EventCredentialExported = "auth.exported"
	EventCredentialImported = "auth.imported"
	EventCredentialRotated = "auth.rotated"
	EventCredentialAccessed = "auth.accessed"
	EventCredentialDenied = "auth.denied"
	EventCredentialApproved = "auth.approved"
	EventCredentialRevoked = "auth.revoked"
	EventCredentialExpired = "auth.expired"
	EventCredentialLocked = "auth.locked"
	EventCredentialUnlocked = "auth.unlocked"
	EventCredentialVerified = "auth.verified"
	EventCredentialFailed = "auth.failed"
	EventCredentialSucceeded = "auth.succeeded"
)

// Event is a auth-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownCredentialEvent reports whether t is a known auth event type.
func IsKnownCredentialEvent(t string) bool {
	switch t {
	case EventCredentialCreated:
		return true
	case EventCredentialUpdated:
		return true
	case EventCredentialDeleted:
		return true
	case EventCredentialArchived:
		return true
	case EventCredentialRestored:
		return true
	case EventCredentialShared:
		return true
	case EventCredentialUnshared:
		return true
	case EventCredentialExported:
		return true
	case EventCredentialImported:
		return true
	case EventCredentialRotated:
		return true
	case EventCredentialAccessed:
		return true
	case EventCredentialDenied:
		return true
	case EventCredentialApproved:
		return true
	case EventCredentialRevoked:
		return true
	case EventCredentialExpired:
		return true
	case EventCredentialLocked:
		return true
	case EventCredentialUnlocked:
		return true
	case EventCredentialVerified:
		return true
	case EventCredentialFailed:
		return true
	case EventCredentialSucceeded:
		return true
	default:
		return false
	}
}


// NewCredentialCreatedEvent builds a domain event.
func NewCredentialCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventCredentialCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialUpdatedEvent builds a domain event.
func NewCredentialUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventCredentialUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialDeletedEvent builds a domain event.
func NewCredentialDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventCredentialDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialArchivedEvent builds a domain event.
func NewCredentialArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventCredentialArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialRestoredEvent builds a domain event.
func NewCredentialRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventCredentialRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialSharedEvent builds a domain event.
func NewCredentialSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventCredentialShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialUnsharedEvent builds a domain event.
func NewCredentialUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventCredentialUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialExportedEvent builds a domain event.
func NewCredentialExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventCredentialExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialImportedEvent builds a domain event.
func NewCredentialImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventCredentialImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialRotatedEvent builds a domain event.
func NewCredentialRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventCredentialRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialAccessedEvent builds a domain event.
func NewCredentialAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventCredentialAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialDeniedEvent builds a domain event.
func NewCredentialDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventCredentialDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialApprovedEvent builds a domain event.
func NewCredentialApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventCredentialApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialRevokedEvent builds a domain event.
func NewCredentialRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventCredentialRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialExpiredEvent builds a domain event.
func NewCredentialExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventCredentialExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialLockedEvent builds a domain event.
func NewCredentialLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventCredentialLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialUnlockedEvent builds a domain event.
func NewCredentialUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventCredentialUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialVerifiedEvent builds a domain event.
func NewCredentialVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventCredentialVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialFailedEvent builds a domain event.
func NewCredentialFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventCredentialFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCredentialSucceededEvent builds a domain event.
func NewCredentialSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Credential"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventCredentialSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownCredentialEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

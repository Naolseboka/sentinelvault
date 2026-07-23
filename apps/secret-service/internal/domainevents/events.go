package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for secret-service.
const (
	EventSecretItemCreated   = "secret.created"
	EventSecretItemUpdated   = "secret.updated"
	EventSecretItemDeleted   = "secret.deleted"
	EventSecretItemArchived  = "secret.archived"
	EventSecretItemRestored  = "secret.restored"
	EventSecretItemShared    = "secret.shared"
	EventSecretItemUnshared  = "secret.unshared"
	EventSecretItemExported  = "secret.exported"
	EventSecretItemImported  = "secret.imported"
	EventSecretItemRotated   = "secret.rotated"
	EventSecretItemAccessed  = "secret.accessed"
	EventSecretItemDenied    = "secret.denied"
	EventSecretItemApproved  = "secret.approved"
	EventSecretItemRevoked   = "secret.revoked"
	EventSecretItemExpired   = "secret.expired"
	EventSecretItemLocked    = "secret.locked"
	EventSecretItemUnlocked  = "secret.unlocked"
	EventSecretItemVerified  = "secret.verified"
	EventSecretItemFailed    = "secret.failed"
	EventSecretItemSucceeded = "secret.succeeded"
)

// Event is a secret-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownSecretItemEvent reports whether t is a known secret event type.
func IsKnownSecretItemEvent(t string) bool {
	switch t {
	case EventSecretItemCreated:
		return true
	case EventSecretItemUpdated:
		return true
	case EventSecretItemDeleted:
		return true
	case EventSecretItemArchived:
		return true
	case EventSecretItemRestored:
		return true
	case EventSecretItemShared:
		return true
	case EventSecretItemUnshared:
		return true
	case EventSecretItemExported:
		return true
	case EventSecretItemImported:
		return true
	case EventSecretItemRotated:
		return true
	case EventSecretItemAccessed:
		return true
	case EventSecretItemDenied:
		return true
	case EventSecretItemApproved:
		return true
	case EventSecretItemRevoked:
		return true
	case EventSecretItemExpired:
		return true
	case EventSecretItemLocked:
		return true
	case EventSecretItemUnlocked:
		return true
	case EventSecretItemVerified:
		return true
	case EventSecretItemFailed:
		return true
	case EventSecretItemSucceeded:
		return true
	default:
		return false
	}
}

// NewSecretItemCreatedEvent builds a domain event.
func NewSecretItemCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventSecretItemCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemUpdatedEvent builds a domain event.
func NewSecretItemUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventSecretItemUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemDeletedEvent builds a domain event.
func NewSecretItemDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventSecretItemDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemArchivedEvent builds a domain event.
func NewSecretItemArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventSecretItemArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemRestoredEvent builds a domain event.
func NewSecretItemRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventSecretItemRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemSharedEvent builds a domain event.
func NewSecretItemSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventSecretItemShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemUnsharedEvent builds a domain event.
func NewSecretItemUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventSecretItemUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemExportedEvent builds a domain event.
func NewSecretItemExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventSecretItemExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemImportedEvent builds a domain event.
func NewSecretItemImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventSecretItemImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemRotatedEvent builds a domain event.
func NewSecretItemRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventSecretItemRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemAccessedEvent builds a domain event.
func NewSecretItemAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventSecretItemAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemDeniedEvent builds a domain event.
func NewSecretItemDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventSecretItemDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemApprovedEvent builds a domain event.
func NewSecretItemApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventSecretItemApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemRevokedEvent builds a domain event.
func NewSecretItemRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventSecretItemRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemExpiredEvent builds a domain event.
func NewSecretItemExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventSecretItemExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemLockedEvent builds a domain event.
func NewSecretItemLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventSecretItemLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemUnlockedEvent builds a domain event.
func NewSecretItemUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventSecretItemUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemVerifiedEvent builds a domain event.
func NewSecretItemVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventSecretItemVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemFailedEvent builds a domain event.
func NewSecretItemFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventSecretItemFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSecretItemSucceededEvent builds a domain event.
func NewSecretItemSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SecretItem"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventSecretItemSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownSecretItemEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

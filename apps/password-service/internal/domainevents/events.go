package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for password-service.
const (
	EventVaultItemCreated   = "password.created"
	EventVaultItemUpdated   = "password.updated"
	EventVaultItemDeleted   = "password.deleted"
	EventVaultItemArchived  = "password.archived"
	EventVaultItemRestored  = "password.restored"
	EventVaultItemShared    = "password.shared"
	EventVaultItemUnshared  = "password.unshared"
	EventVaultItemExported  = "password.exported"
	EventVaultItemImported  = "password.imported"
	EventVaultItemRotated   = "password.rotated"
	EventVaultItemAccessed  = "password.accessed"
	EventVaultItemDenied    = "password.denied"
	EventVaultItemApproved  = "password.approved"
	EventVaultItemRevoked   = "password.revoked"
	EventVaultItemExpired   = "password.expired"
	EventVaultItemLocked    = "password.locked"
	EventVaultItemUnlocked  = "password.unlocked"
	EventVaultItemVerified  = "password.verified"
	EventVaultItemFailed    = "password.failed"
	EventVaultItemSucceeded = "password.succeeded"
)

// Event is a password-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownVaultItemEvent reports whether t is a known password event type.
func IsKnownVaultItemEvent(t string) bool {
	switch t {
	case EventVaultItemCreated:
		return true
	case EventVaultItemUpdated:
		return true
	case EventVaultItemDeleted:
		return true
	case EventVaultItemArchived:
		return true
	case EventVaultItemRestored:
		return true
	case EventVaultItemShared:
		return true
	case EventVaultItemUnshared:
		return true
	case EventVaultItemExported:
		return true
	case EventVaultItemImported:
		return true
	case EventVaultItemRotated:
		return true
	case EventVaultItemAccessed:
		return true
	case EventVaultItemDenied:
		return true
	case EventVaultItemApproved:
		return true
	case EventVaultItemRevoked:
		return true
	case EventVaultItemExpired:
		return true
	case EventVaultItemLocked:
		return true
	case EventVaultItemUnlocked:
		return true
	case EventVaultItemVerified:
		return true
	case EventVaultItemFailed:
		return true
	case EventVaultItemSucceeded:
		return true
	default:
		return false
	}
}

// NewVaultItemCreatedEvent builds a domain event.
func NewVaultItemCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventVaultItemCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemUpdatedEvent builds a domain event.
func NewVaultItemUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventVaultItemUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemDeletedEvent builds a domain event.
func NewVaultItemDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventVaultItemDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemArchivedEvent builds a domain event.
func NewVaultItemArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventVaultItemArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemRestoredEvent builds a domain event.
func NewVaultItemRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventVaultItemRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemSharedEvent builds a domain event.
func NewVaultItemSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventVaultItemShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemUnsharedEvent builds a domain event.
func NewVaultItemUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventVaultItemUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemExportedEvent builds a domain event.
func NewVaultItemExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventVaultItemExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemImportedEvent builds a domain event.
func NewVaultItemImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventVaultItemImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemRotatedEvent builds a domain event.
func NewVaultItemRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventVaultItemRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemAccessedEvent builds a domain event.
func NewVaultItemAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventVaultItemAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemDeniedEvent builds a domain event.
func NewVaultItemDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventVaultItemDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemApprovedEvent builds a domain event.
func NewVaultItemApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventVaultItemApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemRevokedEvent builds a domain event.
func NewVaultItemRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventVaultItemRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemExpiredEvent builds a domain event.
func NewVaultItemExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventVaultItemExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemLockedEvent builds a domain event.
func NewVaultItemLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventVaultItemLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemUnlockedEvent builds a domain event.
func NewVaultItemUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventVaultItemUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemVerifiedEvent builds a domain event.
func NewVaultItemVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventVaultItemVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemFailedEvent builds a domain event.
func NewVaultItemFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventVaultItemFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewVaultItemSucceededEvent builds a domain event.
func NewVaultItemSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "VaultItem"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventVaultItemSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownVaultItemEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

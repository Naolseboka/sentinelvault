package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for encryption-service.
const (
	EventEncryptOpCreated   = "encryption.created"
	EventEncryptOpUpdated   = "encryption.updated"
	EventEncryptOpDeleted   = "encryption.deleted"
	EventEncryptOpArchived  = "encryption.archived"
	EventEncryptOpRestored  = "encryption.restored"
	EventEncryptOpShared    = "encryption.shared"
	EventEncryptOpUnshared  = "encryption.unshared"
	EventEncryptOpExported  = "encryption.exported"
	EventEncryptOpImported  = "encryption.imported"
	EventEncryptOpRotated   = "encryption.rotated"
	EventEncryptOpAccessed  = "encryption.accessed"
	EventEncryptOpDenied    = "encryption.denied"
	EventEncryptOpApproved  = "encryption.approved"
	EventEncryptOpRevoked   = "encryption.revoked"
	EventEncryptOpExpired   = "encryption.expired"
	EventEncryptOpLocked    = "encryption.locked"
	EventEncryptOpUnlocked  = "encryption.unlocked"
	EventEncryptOpVerified  = "encryption.verified"
	EventEncryptOpFailed    = "encryption.failed"
	EventEncryptOpSucceeded = "encryption.succeeded"
)

// Event is a encryption-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownEncryptOpEvent reports whether t is a known encryption event type.
func IsKnownEncryptOpEvent(t string) bool {
	switch t {
	case EventEncryptOpCreated:
		return true
	case EventEncryptOpUpdated:
		return true
	case EventEncryptOpDeleted:
		return true
	case EventEncryptOpArchived:
		return true
	case EventEncryptOpRestored:
		return true
	case EventEncryptOpShared:
		return true
	case EventEncryptOpUnshared:
		return true
	case EventEncryptOpExported:
		return true
	case EventEncryptOpImported:
		return true
	case EventEncryptOpRotated:
		return true
	case EventEncryptOpAccessed:
		return true
	case EventEncryptOpDenied:
		return true
	case EventEncryptOpApproved:
		return true
	case EventEncryptOpRevoked:
		return true
	case EventEncryptOpExpired:
		return true
	case EventEncryptOpLocked:
		return true
	case EventEncryptOpUnlocked:
		return true
	case EventEncryptOpVerified:
		return true
	case EventEncryptOpFailed:
		return true
	case EventEncryptOpSucceeded:
		return true
	default:
		return false
	}
}

// NewEncryptOpCreatedEvent builds a domain event.
func NewEncryptOpCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpUpdatedEvent builds a domain event.
func NewEncryptOpUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpDeletedEvent builds a domain event.
func NewEncryptOpDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpArchivedEvent builds a domain event.
func NewEncryptOpArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpRestoredEvent builds a domain event.
func NewEncryptOpRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpSharedEvent builds a domain event.
func NewEncryptOpSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpUnsharedEvent builds a domain event.
func NewEncryptOpUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpExportedEvent builds a domain event.
func NewEncryptOpExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpImportedEvent builds a domain event.
func NewEncryptOpImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpRotatedEvent builds a domain event.
func NewEncryptOpRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpAccessedEvent builds a domain event.
func NewEncryptOpAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpDeniedEvent builds a domain event.
func NewEncryptOpDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpApprovedEvent builds a domain event.
func NewEncryptOpApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpRevokedEvent builds a domain event.
func NewEncryptOpRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpExpiredEvent builds a domain event.
func NewEncryptOpExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpLockedEvent builds a domain event.
func NewEncryptOpLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpUnlockedEvent builds a domain event.
func NewEncryptOpUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpVerifiedEvent builds a domain event.
func NewEncryptOpVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpFailedEvent builds a domain event.
func NewEncryptOpFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEncryptOpSucceededEvent builds a domain event.
func NewEncryptOpSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EncryptOp"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventEncryptOpSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownEncryptOpEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

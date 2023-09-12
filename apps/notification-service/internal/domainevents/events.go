package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for notification-service.
const (
	EventMessageCreated = "notification.created"
	EventMessageUpdated = "notification.updated"
	EventMessageDeleted = "notification.deleted"
	EventMessageArchived = "notification.archived"
	EventMessageRestored = "notification.restored"
	EventMessageShared = "notification.shared"
	EventMessageUnshared = "notification.unshared"
	EventMessageExported = "notification.exported"
	EventMessageImported = "notification.imported"
	EventMessageRotated = "notification.rotated"
	EventMessageAccessed = "notification.accessed"
	EventMessageDenied = "notification.denied"
	EventMessageApproved = "notification.approved"
	EventMessageRevoked = "notification.revoked"
	EventMessageExpired = "notification.expired"
	EventMessageLocked = "notification.locked"
	EventMessageUnlocked = "notification.unlocked"
	EventMessageVerified = "notification.verified"
	EventMessageFailed = "notification.failed"
	EventMessageSucceeded = "notification.succeeded"
)

// Event is a notification-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownMessageEvent reports whether t is a known notification event type.
func IsKnownMessageEvent(t string) bool {
	switch t {
	case EventMessageCreated:
		return true
	case EventMessageUpdated:
		return true
	case EventMessageDeleted:
		return true
	case EventMessageArchived:
		return true
	case EventMessageRestored:
		return true
	case EventMessageShared:
		return true
	case EventMessageUnshared:
		return true
	case EventMessageExported:
		return true
	case EventMessageImported:
		return true
	case EventMessageRotated:
		return true
	case EventMessageAccessed:
		return true
	case EventMessageDenied:
		return true
	case EventMessageApproved:
		return true
	case EventMessageRevoked:
		return true
	case EventMessageExpired:
		return true
	case EventMessageLocked:
		return true
	case EventMessageUnlocked:
		return true
	case EventMessageVerified:
		return true
	case EventMessageFailed:
		return true
	case EventMessageSucceeded:
		return true
	default:
		return false
	}
}


// NewMessageCreatedEvent builds a domain event.
func NewMessageCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventMessageCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageUpdatedEvent builds a domain event.
func NewMessageUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventMessageUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageDeletedEvent builds a domain event.
func NewMessageDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventMessageDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageArchivedEvent builds a domain event.
func NewMessageArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventMessageArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageRestoredEvent builds a domain event.
func NewMessageRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventMessageRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageSharedEvent builds a domain event.
func NewMessageSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventMessageShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageUnsharedEvent builds a domain event.
func NewMessageUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventMessageUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageExportedEvent builds a domain event.
func NewMessageExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventMessageExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageImportedEvent builds a domain event.
func NewMessageImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventMessageImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageRotatedEvent builds a domain event.
func NewMessageRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventMessageRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageAccessedEvent builds a domain event.
func NewMessageAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventMessageAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageDeniedEvent builds a domain event.
func NewMessageDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventMessageDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageApprovedEvent builds a domain event.
func NewMessageApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventMessageApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageRevokedEvent builds a domain event.
func NewMessageRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventMessageRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageExpiredEvent builds a domain event.
func NewMessageExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventMessageExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageLockedEvent builds a domain event.
func NewMessageLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventMessageLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageUnlockedEvent builds a domain event.
func NewMessageUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventMessageUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageVerifiedEvent builds a domain event.
func NewMessageVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventMessageVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageFailedEvent builds a domain event.
func NewMessageFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventMessageFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewMessageSucceededEvent builds a domain event.
func NewMessageSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Message"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventMessageSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownMessageEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

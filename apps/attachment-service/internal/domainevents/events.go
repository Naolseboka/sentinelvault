package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for attachment-service.
const (
	EventAttachmentCreated   = "attachment.created"
	EventAttachmentUpdated   = "attachment.updated"
	EventAttachmentDeleted   = "attachment.deleted"
	EventAttachmentArchived  = "attachment.archived"
	EventAttachmentRestored  = "attachment.restored"
	EventAttachmentShared    = "attachment.shared"
	EventAttachmentUnshared  = "attachment.unshared"
	EventAttachmentExported  = "attachment.exported"
	EventAttachmentImported  = "attachment.imported"
	EventAttachmentRotated   = "attachment.rotated"
	EventAttachmentAccessed  = "attachment.accessed"
	EventAttachmentDenied    = "attachment.denied"
	EventAttachmentApproved  = "attachment.approved"
	EventAttachmentRevoked   = "attachment.revoked"
	EventAttachmentExpired   = "attachment.expired"
	EventAttachmentLocked    = "attachment.locked"
	EventAttachmentUnlocked  = "attachment.unlocked"
	EventAttachmentVerified  = "attachment.verified"
	EventAttachmentFailed    = "attachment.failed"
	EventAttachmentSucceeded = "attachment.succeeded"
)

// Event is a attachment-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownAttachmentEvent reports whether t is a known attachment event type.
func IsKnownAttachmentEvent(t string) bool {
	switch t {
	case EventAttachmentCreated:
		return true
	case EventAttachmentUpdated:
		return true
	case EventAttachmentDeleted:
		return true
	case EventAttachmentArchived:
		return true
	case EventAttachmentRestored:
		return true
	case EventAttachmentShared:
		return true
	case EventAttachmentUnshared:
		return true
	case EventAttachmentExported:
		return true
	case EventAttachmentImported:
		return true
	case EventAttachmentRotated:
		return true
	case EventAttachmentAccessed:
		return true
	case EventAttachmentDenied:
		return true
	case EventAttachmentApproved:
		return true
	case EventAttachmentRevoked:
		return true
	case EventAttachmentExpired:
		return true
	case EventAttachmentLocked:
		return true
	case EventAttachmentUnlocked:
		return true
	case EventAttachmentVerified:
		return true
	case EventAttachmentFailed:
		return true
	case EventAttachmentSucceeded:
		return true
	default:
		return false
	}
}

// NewAttachmentCreatedEvent builds a domain event.
func NewAttachmentCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventAttachmentCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentUpdatedEvent builds a domain event.
func NewAttachmentUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventAttachmentUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentDeletedEvent builds a domain event.
func NewAttachmentDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventAttachmentDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentArchivedEvent builds a domain event.
func NewAttachmentArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventAttachmentArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentRestoredEvent builds a domain event.
func NewAttachmentRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventAttachmentRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentSharedEvent builds a domain event.
func NewAttachmentSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventAttachmentShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentUnsharedEvent builds a domain event.
func NewAttachmentUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventAttachmentUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentExportedEvent builds a domain event.
func NewAttachmentExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventAttachmentExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentImportedEvent builds a domain event.
func NewAttachmentImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventAttachmentImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentRotatedEvent builds a domain event.
func NewAttachmentRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventAttachmentRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentAccessedEvent builds a domain event.
func NewAttachmentAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventAttachmentAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentDeniedEvent builds a domain event.
func NewAttachmentDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventAttachmentDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentApprovedEvent builds a domain event.
func NewAttachmentApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventAttachmentApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentRevokedEvent builds a domain event.
func NewAttachmentRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventAttachmentRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentExpiredEvent builds a domain event.
func NewAttachmentExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventAttachmentExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentLockedEvent builds a domain event.
func NewAttachmentLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventAttachmentLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentUnlockedEvent builds a domain event.
func NewAttachmentUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventAttachmentUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentVerifiedEvent builds a domain event.
func NewAttachmentVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventAttachmentVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentFailedEvent builds a domain event.
func NewAttachmentFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventAttachmentFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAttachmentSucceededEvent builds a domain event.
func NewAttachmentSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Attachment"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventAttachmentSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownAttachmentEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for audit-service.
const (
	EventAuditEventCreated = "audit.created"
	EventAuditEventUpdated = "audit.updated"
	EventAuditEventDeleted = "audit.deleted"
	EventAuditEventArchived = "audit.archived"
	EventAuditEventRestored = "audit.restored"
	EventAuditEventShared = "audit.shared"
	EventAuditEventUnshared = "audit.unshared"
	EventAuditEventExported = "audit.exported"
	EventAuditEventImported = "audit.imported"
	EventAuditEventRotated = "audit.rotated"
	EventAuditEventAccessed = "audit.accessed"
	EventAuditEventDenied = "audit.denied"
	EventAuditEventApproved = "audit.approved"
	EventAuditEventRevoked = "audit.revoked"
	EventAuditEventExpired = "audit.expired"
	EventAuditEventLocked = "audit.locked"
	EventAuditEventUnlocked = "audit.unlocked"
	EventAuditEventVerified = "audit.verified"
	EventAuditEventFailed = "audit.failed"
	EventAuditEventSucceeded = "audit.succeeded"
)

// Event is a audit-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownAuditEventEvent reports whether t is a known audit event type.
func IsKnownAuditEventEvent(t string) bool {
	switch t {
	case EventAuditEventCreated:
		return true
	case EventAuditEventUpdated:
		return true
	case EventAuditEventDeleted:
		return true
	case EventAuditEventArchived:
		return true
	case EventAuditEventRestored:
		return true
	case EventAuditEventShared:
		return true
	case EventAuditEventUnshared:
		return true
	case EventAuditEventExported:
		return true
	case EventAuditEventImported:
		return true
	case EventAuditEventRotated:
		return true
	case EventAuditEventAccessed:
		return true
	case EventAuditEventDenied:
		return true
	case EventAuditEventApproved:
		return true
	case EventAuditEventRevoked:
		return true
	case EventAuditEventExpired:
		return true
	case EventAuditEventLocked:
		return true
	case EventAuditEventUnlocked:
		return true
	case EventAuditEventVerified:
		return true
	case EventAuditEventFailed:
		return true
	case EventAuditEventSucceeded:
		return true
	default:
		return false
	}
}


// NewAuditEventCreatedEvent builds a domain event.
func NewAuditEventCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventAuditEventCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventUpdatedEvent builds a domain event.
func NewAuditEventUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventAuditEventUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventDeletedEvent builds a domain event.
func NewAuditEventDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventAuditEventDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventArchivedEvent builds a domain event.
func NewAuditEventArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventAuditEventArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventRestoredEvent builds a domain event.
func NewAuditEventRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventAuditEventRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventSharedEvent builds a domain event.
func NewAuditEventSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventAuditEventShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventUnsharedEvent builds a domain event.
func NewAuditEventUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventAuditEventUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventExportedEvent builds a domain event.
func NewAuditEventExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventAuditEventExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventImportedEvent builds a domain event.
func NewAuditEventImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventAuditEventImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventRotatedEvent builds a domain event.
func NewAuditEventRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventAuditEventRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventAccessedEvent builds a domain event.
func NewAuditEventAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventAuditEventAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventDeniedEvent builds a domain event.
func NewAuditEventDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventAuditEventDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventApprovedEvent builds a domain event.
func NewAuditEventApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventAuditEventApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventRevokedEvent builds a domain event.
func NewAuditEventRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventAuditEventRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventExpiredEvent builds a domain event.
func NewAuditEventExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventAuditEventExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventLockedEvent builds a domain event.
func NewAuditEventLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventAuditEventLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventUnlockedEvent builds a domain event.
func NewAuditEventUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventAuditEventUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventVerifiedEvent builds a domain event.
func NewAuditEventVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventAuditEventVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventFailedEvent builds a domain event.
func NewAuditEventFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventAuditEventFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuditEventSucceededEvent builds a domain event.
func NewAuditEventSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuditEvent"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventAuditEventSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownAuditEventEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for search-service.
const (
	EventDocumentCreated   = "search.created"
	EventDocumentUpdated   = "search.updated"
	EventDocumentDeleted   = "search.deleted"
	EventDocumentArchived  = "search.archived"
	EventDocumentRestored  = "search.restored"
	EventDocumentShared    = "search.shared"
	EventDocumentUnshared  = "search.unshared"
	EventDocumentExported  = "search.exported"
	EventDocumentImported  = "search.imported"
	EventDocumentRotated   = "search.rotated"
	EventDocumentAccessed  = "search.accessed"
	EventDocumentDenied    = "search.denied"
	EventDocumentApproved  = "search.approved"
	EventDocumentRevoked   = "search.revoked"
	EventDocumentExpired   = "search.expired"
	EventDocumentLocked    = "search.locked"
	EventDocumentUnlocked  = "search.unlocked"
	EventDocumentVerified  = "search.verified"
	EventDocumentFailed    = "search.failed"
	EventDocumentSucceeded = "search.succeeded"
)

// Event is a search-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownDocumentEvent reports whether t is a known search event type.
func IsKnownDocumentEvent(t string) bool {
	switch t {
	case EventDocumentCreated:
		return true
	case EventDocumentUpdated:
		return true
	case EventDocumentDeleted:
		return true
	case EventDocumentArchived:
		return true
	case EventDocumentRestored:
		return true
	case EventDocumentShared:
		return true
	case EventDocumentUnshared:
		return true
	case EventDocumentExported:
		return true
	case EventDocumentImported:
		return true
	case EventDocumentRotated:
		return true
	case EventDocumentAccessed:
		return true
	case EventDocumentDenied:
		return true
	case EventDocumentApproved:
		return true
	case EventDocumentRevoked:
		return true
	case EventDocumentExpired:
		return true
	case EventDocumentLocked:
		return true
	case EventDocumentUnlocked:
		return true
	case EventDocumentVerified:
		return true
	case EventDocumentFailed:
		return true
	case EventDocumentSucceeded:
		return true
	default:
		return false
	}
}

// NewDocumentCreatedEvent builds a domain event.
func NewDocumentCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventDocumentCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentUpdatedEvent builds a domain event.
func NewDocumentUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventDocumentUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentDeletedEvent builds a domain event.
func NewDocumentDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventDocumentDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentArchivedEvent builds a domain event.
func NewDocumentArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventDocumentArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentRestoredEvent builds a domain event.
func NewDocumentRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventDocumentRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentSharedEvent builds a domain event.
func NewDocumentSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventDocumentShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentUnsharedEvent builds a domain event.
func NewDocumentUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventDocumentUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentExportedEvent builds a domain event.
func NewDocumentExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventDocumentExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentImportedEvent builds a domain event.
func NewDocumentImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventDocumentImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentRotatedEvent builds a domain event.
func NewDocumentRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventDocumentRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentAccessedEvent builds a domain event.
func NewDocumentAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventDocumentAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentDeniedEvent builds a domain event.
func NewDocumentDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventDocumentDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentApprovedEvent builds a domain event.
func NewDocumentApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventDocumentApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentRevokedEvent builds a domain event.
func NewDocumentRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventDocumentRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentExpiredEvent builds a domain event.
func NewDocumentExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventDocumentExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentLockedEvent builds a domain event.
func NewDocumentLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventDocumentLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentUnlockedEvent builds a domain event.
func NewDocumentUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventDocumentUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentVerifiedEvent builds a domain event.
func NewDocumentVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventDocumentVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentFailedEvent builds a domain event.
func NewDocumentFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventDocumentFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDocumentSucceededEvent builds a domain event.
func NewDocumentSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Document"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventDocumentSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownDocumentEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

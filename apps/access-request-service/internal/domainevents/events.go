package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for access-request-service.
const (
	EventAccessRequestCreated   = "accessrequest.created"
	EventAccessRequestUpdated   = "accessrequest.updated"
	EventAccessRequestDeleted   = "accessrequest.deleted"
	EventAccessRequestArchived  = "accessrequest.archived"
	EventAccessRequestRestored  = "accessrequest.restored"
	EventAccessRequestShared    = "accessrequest.shared"
	EventAccessRequestUnshared  = "accessrequest.unshared"
	EventAccessRequestExported  = "accessrequest.exported"
	EventAccessRequestImported  = "accessrequest.imported"
	EventAccessRequestRotated   = "accessrequest.rotated"
	EventAccessRequestAccessed  = "accessrequest.accessed"
	EventAccessRequestDenied    = "accessrequest.denied"
	EventAccessRequestApproved  = "accessrequest.approved"
	EventAccessRequestRevoked   = "accessrequest.revoked"
	EventAccessRequestExpired   = "accessrequest.expired"
	EventAccessRequestLocked    = "accessrequest.locked"
	EventAccessRequestUnlocked  = "accessrequest.unlocked"
	EventAccessRequestVerified  = "accessrequest.verified"
	EventAccessRequestFailed    = "accessrequest.failed"
	EventAccessRequestSucceeded = "accessrequest.succeeded"
)

// Event is a access-request-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownAccessRequestEvent reports whether t is a known accessrequest event type.
func IsKnownAccessRequestEvent(t string) bool {
	switch t {
	case EventAccessRequestCreated:
		return true
	case EventAccessRequestUpdated:
		return true
	case EventAccessRequestDeleted:
		return true
	case EventAccessRequestArchived:
		return true
	case EventAccessRequestRestored:
		return true
	case EventAccessRequestShared:
		return true
	case EventAccessRequestUnshared:
		return true
	case EventAccessRequestExported:
		return true
	case EventAccessRequestImported:
		return true
	case EventAccessRequestRotated:
		return true
	case EventAccessRequestAccessed:
		return true
	case EventAccessRequestDenied:
		return true
	case EventAccessRequestApproved:
		return true
	case EventAccessRequestRevoked:
		return true
	case EventAccessRequestExpired:
		return true
	case EventAccessRequestLocked:
		return true
	case EventAccessRequestUnlocked:
		return true
	case EventAccessRequestVerified:
		return true
	case EventAccessRequestFailed:
		return true
	case EventAccessRequestSucceeded:
		return true
	default:
		return false
	}
}

// NewAccessRequestCreatedEvent builds a domain event.
func NewAccessRequestCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestUpdatedEvent builds a domain event.
func NewAccessRequestUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestDeletedEvent builds a domain event.
func NewAccessRequestDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestArchivedEvent builds a domain event.
func NewAccessRequestArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestRestoredEvent builds a domain event.
func NewAccessRequestRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestSharedEvent builds a domain event.
func NewAccessRequestSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestUnsharedEvent builds a domain event.
func NewAccessRequestUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestExportedEvent builds a domain event.
func NewAccessRequestExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestImportedEvent builds a domain event.
func NewAccessRequestImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestRotatedEvent builds a domain event.
func NewAccessRequestRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestAccessedEvent builds a domain event.
func NewAccessRequestAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestDeniedEvent builds a domain event.
func NewAccessRequestDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestApprovedEvent builds a domain event.
func NewAccessRequestApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestRevokedEvent builds a domain event.
func NewAccessRequestRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestExpiredEvent builds a domain event.
func NewAccessRequestExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestLockedEvent builds a domain event.
func NewAccessRequestLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestUnlockedEvent builds a domain event.
func NewAccessRequestUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestVerifiedEvent builds a domain event.
func NewAccessRequestVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestFailedEvent builds a domain event.
func NewAccessRequestFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAccessRequestSucceededEvent builds a domain event.
func NewAccessRequestSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AccessRequest"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventAccessRequestSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownAccessRequestEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

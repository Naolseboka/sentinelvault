package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for billing-service.
const (
	EventLicenseCreated = "billing.created"
	EventLicenseUpdated = "billing.updated"
	EventLicenseDeleted = "billing.deleted"
	EventLicenseArchived = "billing.archived"
	EventLicenseRestored = "billing.restored"
	EventLicenseShared = "billing.shared"
	EventLicenseUnshared = "billing.unshared"
	EventLicenseExported = "billing.exported"
	EventLicenseImported = "billing.imported"
	EventLicenseRotated = "billing.rotated"
	EventLicenseAccessed = "billing.accessed"
	EventLicenseDenied = "billing.denied"
	EventLicenseApproved = "billing.approved"
	EventLicenseRevoked = "billing.revoked"
	EventLicenseExpired = "billing.expired"
	EventLicenseLocked = "billing.locked"
	EventLicenseUnlocked = "billing.unlocked"
	EventLicenseVerified = "billing.verified"
	EventLicenseFailed = "billing.failed"
	EventLicenseSucceeded = "billing.succeeded"
)

// Event is a billing-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownLicenseEvent reports whether t is a known billing event type.
func IsKnownLicenseEvent(t string) bool {
	switch t {
	case EventLicenseCreated:
		return true
	case EventLicenseUpdated:
		return true
	case EventLicenseDeleted:
		return true
	case EventLicenseArchived:
		return true
	case EventLicenseRestored:
		return true
	case EventLicenseShared:
		return true
	case EventLicenseUnshared:
		return true
	case EventLicenseExported:
		return true
	case EventLicenseImported:
		return true
	case EventLicenseRotated:
		return true
	case EventLicenseAccessed:
		return true
	case EventLicenseDenied:
		return true
	case EventLicenseApproved:
		return true
	case EventLicenseRevoked:
		return true
	case EventLicenseExpired:
		return true
	case EventLicenseLocked:
		return true
	case EventLicenseUnlocked:
		return true
	case EventLicenseVerified:
		return true
	case EventLicenseFailed:
		return true
	case EventLicenseSucceeded:
		return true
	default:
		return false
	}
}


// NewLicenseCreatedEvent builds a domain event.
func NewLicenseCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventLicenseCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseUpdatedEvent builds a domain event.
func NewLicenseUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventLicenseUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseDeletedEvent builds a domain event.
func NewLicenseDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventLicenseDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseArchivedEvent builds a domain event.
func NewLicenseArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventLicenseArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseRestoredEvent builds a domain event.
func NewLicenseRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventLicenseRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseSharedEvent builds a domain event.
func NewLicenseSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventLicenseShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseUnsharedEvent builds a domain event.
func NewLicenseUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventLicenseUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseExportedEvent builds a domain event.
func NewLicenseExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventLicenseExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseImportedEvent builds a domain event.
func NewLicenseImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventLicenseImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseRotatedEvent builds a domain event.
func NewLicenseRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventLicenseRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseAccessedEvent builds a domain event.
func NewLicenseAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventLicenseAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseDeniedEvent builds a domain event.
func NewLicenseDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventLicenseDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseApprovedEvent builds a domain event.
func NewLicenseApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventLicenseApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseRevokedEvent builds a domain event.
func NewLicenseRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventLicenseRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseExpiredEvent builds a domain event.
func NewLicenseExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventLicenseExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseLockedEvent builds a domain event.
func NewLicenseLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventLicenseLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseUnlockedEvent builds a domain event.
func NewLicenseUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventLicenseUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseVerifiedEvent builds a domain event.
func NewLicenseVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventLicenseVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseFailedEvent builds a domain event.
func NewLicenseFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventLicenseFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewLicenseSucceededEvent builds a domain event.
func NewLicenseSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "License"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventLicenseSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownLicenseEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

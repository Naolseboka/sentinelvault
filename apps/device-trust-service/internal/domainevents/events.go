package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for device-trust-service.
const (
	EventTrustRecordCreated   = "devicetrust.created"
	EventTrustRecordUpdated   = "devicetrust.updated"
	EventTrustRecordDeleted   = "devicetrust.deleted"
	EventTrustRecordArchived  = "devicetrust.archived"
	EventTrustRecordRestored  = "devicetrust.restored"
	EventTrustRecordShared    = "devicetrust.shared"
	EventTrustRecordUnshared  = "devicetrust.unshared"
	EventTrustRecordExported  = "devicetrust.exported"
	EventTrustRecordImported  = "devicetrust.imported"
	EventTrustRecordRotated   = "devicetrust.rotated"
	EventTrustRecordAccessed  = "devicetrust.accessed"
	EventTrustRecordDenied    = "devicetrust.denied"
	EventTrustRecordApproved  = "devicetrust.approved"
	EventTrustRecordRevoked   = "devicetrust.revoked"
	EventTrustRecordExpired   = "devicetrust.expired"
	EventTrustRecordLocked    = "devicetrust.locked"
	EventTrustRecordUnlocked  = "devicetrust.unlocked"
	EventTrustRecordVerified  = "devicetrust.verified"
	EventTrustRecordFailed    = "devicetrust.failed"
	EventTrustRecordSucceeded = "devicetrust.succeeded"
)

// Event is a device-trust-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownTrustRecordEvent reports whether t is a known devicetrust event type.
func IsKnownTrustRecordEvent(t string) bool {
	switch t {
	case EventTrustRecordCreated:
		return true
	case EventTrustRecordUpdated:
		return true
	case EventTrustRecordDeleted:
		return true
	case EventTrustRecordArchived:
		return true
	case EventTrustRecordRestored:
		return true
	case EventTrustRecordShared:
		return true
	case EventTrustRecordUnshared:
		return true
	case EventTrustRecordExported:
		return true
	case EventTrustRecordImported:
		return true
	case EventTrustRecordRotated:
		return true
	case EventTrustRecordAccessed:
		return true
	case EventTrustRecordDenied:
		return true
	case EventTrustRecordApproved:
		return true
	case EventTrustRecordRevoked:
		return true
	case EventTrustRecordExpired:
		return true
	case EventTrustRecordLocked:
		return true
	case EventTrustRecordUnlocked:
		return true
	case EventTrustRecordVerified:
		return true
	case EventTrustRecordFailed:
		return true
	case EventTrustRecordSucceeded:
		return true
	default:
		return false
	}
}

// NewTrustRecordCreatedEvent builds a domain event.
func NewTrustRecordCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordUpdatedEvent builds a domain event.
func NewTrustRecordUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordDeletedEvent builds a domain event.
func NewTrustRecordDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordArchivedEvent builds a domain event.
func NewTrustRecordArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordRestoredEvent builds a domain event.
func NewTrustRecordRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordSharedEvent builds a domain event.
func NewTrustRecordSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordUnsharedEvent builds a domain event.
func NewTrustRecordUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordExportedEvent builds a domain event.
func NewTrustRecordExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordImportedEvent builds a domain event.
func NewTrustRecordImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordRotatedEvent builds a domain event.
func NewTrustRecordRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordAccessedEvent builds a domain event.
func NewTrustRecordAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordDeniedEvent builds a domain event.
func NewTrustRecordDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordApprovedEvent builds a domain event.
func NewTrustRecordApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordRevokedEvent builds a domain event.
func NewTrustRecordRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordExpiredEvent builds a domain event.
func NewTrustRecordExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordLockedEvent builds a domain event.
func NewTrustRecordLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordUnlockedEvent builds a domain event.
func NewTrustRecordUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordVerifiedEvent builds a domain event.
func NewTrustRecordVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordFailedEvent builds a domain event.
func NewTrustRecordFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTrustRecordSucceededEvent builds a domain event.
func NewTrustRecordSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "TrustRecord"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventTrustRecordSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownTrustRecordEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

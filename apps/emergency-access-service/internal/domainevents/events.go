package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for emergency-access-service.
const (
	EventEmergencyGrantCreated = "emergency.created"
	EventEmergencyGrantUpdated = "emergency.updated"
	EventEmergencyGrantDeleted = "emergency.deleted"
	EventEmergencyGrantArchived = "emergency.archived"
	EventEmergencyGrantRestored = "emergency.restored"
	EventEmergencyGrantShared = "emergency.shared"
	EventEmergencyGrantUnshared = "emergency.unshared"
	EventEmergencyGrantExported = "emergency.exported"
	EventEmergencyGrantImported = "emergency.imported"
	EventEmergencyGrantRotated = "emergency.rotated"
	EventEmergencyGrantAccessed = "emergency.accessed"
	EventEmergencyGrantDenied = "emergency.denied"
	EventEmergencyGrantApproved = "emergency.approved"
	EventEmergencyGrantRevoked = "emergency.revoked"
	EventEmergencyGrantExpired = "emergency.expired"
	EventEmergencyGrantLocked = "emergency.locked"
	EventEmergencyGrantUnlocked = "emergency.unlocked"
	EventEmergencyGrantVerified = "emergency.verified"
	EventEmergencyGrantFailed = "emergency.failed"
	EventEmergencyGrantSucceeded = "emergency.succeeded"
)

// Event is a emergency-access-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownEmergencyGrantEvent reports whether t is a known emergency event type.
func IsKnownEmergencyGrantEvent(t string) bool {
	switch t {
	case EventEmergencyGrantCreated:
		return true
	case EventEmergencyGrantUpdated:
		return true
	case EventEmergencyGrantDeleted:
		return true
	case EventEmergencyGrantArchived:
		return true
	case EventEmergencyGrantRestored:
		return true
	case EventEmergencyGrantShared:
		return true
	case EventEmergencyGrantUnshared:
		return true
	case EventEmergencyGrantExported:
		return true
	case EventEmergencyGrantImported:
		return true
	case EventEmergencyGrantRotated:
		return true
	case EventEmergencyGrantAccessed:
		return true
	case EventEmergencyGrantDenied:
		return true
	case EventEmergencyGrantApproved:
		return true
	case EventEmergencyGrantRevoked:
		return true
	case EventEmergencyGrantExpired:
		return true
	case EventEmergencyGrantLocked:
		return true
	case EventEmergencyGrantUnlocked:
		return true
	case EventEmergencyGrantVerified:
		return true
	case EventEmergencyGrantFailed:
		return true
	case EventEmergencyGrantSucceeded:
		return true
	default:
		return false
	}
}


// NewEmergencyGrantCreatedEvent builds a domain event.
func NewEmergencyGrantCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantUpdatedEvent builds a domain event.
func NewEmergencyGrantUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantDeletedEvent builds a domain event.
func NewEmergencyGrantDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantArchivedEvent builds a domain event.
func NewEmergencyGrantArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantRestoredEvent builds a domain event.
func NewEmergencyGrantRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantSharedEvent builds a domain event.
func NewEmergencyGrantSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantUnsharedEvent builds a domain event.
func NewEmergencyGrantUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantExportedEvent builds a domain event.
func NewEmergencyGrantExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantImportedEvent builds a domain event.
func NewEmergencyGrantImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantRotatedEvent builds a domain event.
func NewEmergencyGrantRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantAccessedEvent builds a domain event.
func NewEmergencyGrantAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantDeniedEvent builds a domain event.
func NewEmergencyGrantDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantApprovedEvent builds a domain event.
func NewEmergencyGrantApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantRevokedEvent builds a domain event.
func NewEmergencyGrantRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantExpiredEvent builds a domain event.
func NewEmergencyGrantExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantLockedEvent builds a domain event.
func NewEmergencyGrantLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantUnlockedEvent builds a domain event.
func NewEmergencyGrantUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantVerifiedEvent builds a domain event.
func NewEmergencyGrantVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantFailedEvent builds a domain event.
func NewEmergencyGrantFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewEmergencyGrantSucceededEvent builds a domain event.
func NewEmergencyGrantSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "EmergencyGrant"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventEmergencyGrantSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownEmergencyGrantEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

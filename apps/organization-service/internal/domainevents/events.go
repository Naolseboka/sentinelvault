package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for organization-service.
const (
	EventOrganizationCreated   = "organization.created"
	EventOrganizationUpdated   = "organization.updated"
	EventOrganizationDeleted   = "organization.deleted"
	EventOrganizationArchived  = "organization.archived"
	EventOrganizationRestored  = "organization.restored"
	EventOrganizationShared    = "organization.shared"
	EventOrganizationUnshared  = "organization.unshared"
	EventOrganizationExported  = "organization.exported"
	EventOrganizationImported  = "organization.imported"
	EventOrganizationRotated   = "organization.rotated"
	EventOrganizationAccessed  = "organization.accessed"
	EventOrganizationDenied    = "organization.denied"
	EventOrganizationApproved  = "organization.approved"
	EventOrganizationRevoked   = "organization.revoked"
	EventOrganizationExpired   = "organization.expired"
	EventOrganizationLocked    = "organization.locked"
	EventOrganizationUnlocked  = "organization.unlocked"
	EventOrganizationVerified  = "organization.verified"
	EventOrganizationFailed    = "organization.failed"
	EventOrganizationSucceeded = "organization.succeeded"
)

// Event is a organization-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownOrganizationEvent reports whether t is a known organization event type.
func IsKnownOrganizationEvent(t string) bool {
	switch t {
	case EventOrganizationCreated:
		return true
	case EventOrganizationUpdated:
		return true
	case EventOrganizationDeleted:
		return true
	case EventOrganizationArchived:
		return true
	case EventOrganizationRestored:
		return true
	case EventOrganizationShared:
		return true
	case EventOrganizationUnshared:
		return true
	case EventOrganizationExported:
		return true
	case EventOrganizationImported:
		return true
	case EventOrganizationRotated:
		return true
	case EventOrganizationAccessed:
		return true
	case EventOrganizationDenied:
		return true
	case EventOrganizationApproved:
		return true
	case EventOrganizationRevoked:
		return true
	case EventOrganizationExpired:
		return true
	case EventOrganizationLocked:
		return true
	case EventOrganizationUnlocked:
		return true
	case EventOrganizationVerified:
		return true
	case EventOrganizationFailed:
		return true
	case EventOrganizationSucceeded:
		return true
	default:
		return false
	}
}

// NewOrganizationCreatedEvent builds a domain event.
func NewOrganizationCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventOrganizationCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationUpdatedEvent builds a domain event.
func NewOrganizationUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventOrganizationUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationDeletedEvent builds a domain event.
func NewOrganizationDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventOrganizationDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationArchivedEvent builds a domain event.
func NewOrganizationArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventOrganizationArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationRestoredEvent builds a domain event.
func NewOrganizationRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventOrganizationRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationSharedEvent builds a domain event.
func NewOrganizationSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventOrganizationShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationUnsharedEvent builds a domain event.
func NewOrganizationUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventOrganizationUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationExportedEvent builds a domain event.
func NewOrganizationExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventOrganizationExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationImportedEvent builds a domain event.
func NewOrganizationImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventOrganizationImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationRotatedEvent builds a domain event.
func NewOrganizationRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventOrganizationRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationAccessedEvent builds a domain event.
func NewOrganizationAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventOrganizationAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationDeniedEvent builds a domain event.
func NewOrganizationDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventOrganizationDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationApprovedEvent builds a domain event.
func NewOrganizationApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventOrganizationApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationRevokedEvent builds a domain event.
func NewOrganizationRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventOrganizationRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationExpiredEvent builds a domain event.
func NewOrganizationExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventOrganizationExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationLockedEvent builds a domain event.
func NewOrganizationLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventOrganizationLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationUnlockedEvent builds a domain event.
func NewOrganizationUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventOrganizationUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationVerifiedEvent builds a domain event.
func NewOrganizationVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventOrganizationVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationFailedEvent builds a domain event.
func NewOrganizationFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventOrganizationFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOrganizationSucceededEvent builds a domain event.
func NewOrganizationSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Organization"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventOrganizationSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownOrganizationEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

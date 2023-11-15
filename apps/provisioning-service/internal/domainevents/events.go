package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for provisioning-service.
const (
	EventProvisionPlanCreated = "provisioning.created"
	EventProvisionPlanUpdated = "provisioning.updated"
	EventProvisionPlanDeleted = "provisioning.deleted"
	EventProvisionPlanArchived = "provisioning.archived"
	EventProvisionPlanRestored = "provisioning.restored"
	EventProvisionPlanShared = "provisioning.shared"
	EventProvisionPlanUnshared = "provisioning.unshared"
	EventProvisionPlanExported = "provisioning.exported"
	EventProvisionPlanImported = "provisioning.imported"
	EventProvisionPlanRotated = "provisioning.rotated"
	EventProvisionPlanAccessed = "provisioning.accessed"
	EventProvisionPlanDenied = "provisioning.denied"
	EventProvisionPlanApproved = "provisioning.approved"
	EventProvisionPlanRevoked = "provisioning.revoked"
	EventProvisionPlanExpired = "provisioning.expired"
	EventProvisionPlanLocked = "provisioning.locked"
	EventProvisionPlanUnlocked = "provisioning.unlocked"
	EventProvisionPlanVerified = "provisioning.verified"
	EventProvisionPlanFailed = "provisioning.failed"
	EventProvisionPlanSucceeded = "provisioning.succeeded"
)

// Event is a provisioning-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownProvisionPlanEvent reports whether t is a known provisioning event type.
func IsKnownProvisionPlanEvent(t string) bool {
	switch t {
	case EventProvisionPlanCreated:
		return true
	case EventProvisionPlanUpdated:
		return true
	case EventProvisionPlanDeleted:
		return true
	case EventProvisionPlanArchived:
		return true
	case EventProvisionPlanRestored:
		return true
	case EventProvisionPlanShared:
		return true
	case EventProvisionPlanUnshared:
		return true
	case EventProvisionPlanExported:
		return true
	case EventProvisionPlanImported:
		return true
	case EventProvisionPlanRotated:
		return true
	case EventProvisionPlanAccessed:
		return true
	case EventProvisionPlanDenied:
		return true
	case EventProvisionPlanApproved:
		return true
	case EventProvisionPlanRevoked:
		return true
	case EventProvisionPlanExpired:
		return true
	case EventProvisionPlanLocked:
		return true
	case EventProvisionPlanUnlocked:
		return true
	case EventProvisionPlanVerified:
		return true
	case EventProvisionPlanFailed:
		return true
	case EventProvisionPlanSucceeded:
		return true
	default:
		return false
	}
}


// NewProvisionPlanCreatedEvent builds a domain event.
func NewProvisionPlanCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanUpdatedEvent builds a domain event.
func NewProvisionPlanUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanDeletedEvent builds a domain event.
func NewProvisionPlanDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanArchivedEvent builds a domain event.
func NewProvisionPlanArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanRestoredEvent builds a domain event.
func NewProvisionPlanRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanSharedEvent builds a domain event.
func NewProvisionPlanSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanUnsharedEvent builds a domain event.
func NewProvisionPlanUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanExportedEvent builds a domain event.
func NewProvisionPlanExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanImportedEvent builds a domain event.
func NewProvisionPlanImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanRotatedEvent builds a domain event.
func NewProvisionPlanRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanAccessedEvent builds a domain event.
func NewProvisionPlanAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanDeniedEvent builds a domain event.
func NewProvisionPlanDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanApprovedEvent builds a domain event.
func NewProvisionPlanApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanRevokedEvent builds a domain event.
func NewProvisionPlanRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanExpiredEvent builds a domain event.
func NewProvisionPlanExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanLockedEvent builds a domain event.
func NewProvisionPlanLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanUnlockedEvent builds a domain event.
func NewProvisionPlanUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanVerifiedEvent builds a domain event.
func NewProvisionPlanVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanFailedEvent builds a domain event.
func NewProvisionPlanFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewProvisionPlanSucceededEvent builds a domain event.
func NewProvisionPlanSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "ProvisionPlan"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventProvisionPlanSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownProvisionPlanEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

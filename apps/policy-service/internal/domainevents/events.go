package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for policy-service.
const (
	EventPolicyDocCreated = "policy.created"
	EventPolicyDocUpdated = "policy.updated"
	EventPolicyDocDeleted = "policy.deleted"
	EventPolicyDocArchived = "policy.archived"
	EventPolicyDocRestored = "policy.restored"
	EventPolicyDocShared = "policy.shared"
	EventPolicyDocUnshared = "policy.unshared"
	EventPolicyDocExported = "policy.exported"
	EventPolicyDocImported = "policy.imported"
	EventPolicyDocRotated = "policy.rotated"
	EventPolicyDocAccessed = "policy.accessed"
	EventPolicyDocDenied = "policy.denied"
	EventPolicyDocApproved = "policy.approved"
	EventPolicyDocRevoked = "policy.revoked"
	EventPolicyDocExpired = "policy.expired"
	EventPolicyDocLocked = "policy.locked"
	EventPolicyDocUnlocked = "policy.unlocked"
	EventPolicyDocVerified = "policy.verified"
	EventPolicyDocFailed = "policy.failed"
	EventPolicyDocSucceeded = "policy.succeeded"
)

// Event is a policy-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownPolicyDocEvent reports whether t is a known policy event type.
func IsKnownPolicyDocEvent(t string) bool {
	switch t {
	case EventPolicyDocCreated:
		return true
	case EventPolicyDocUpdated:
		return true
	case EventPolicyDocDeleted:
		return true
	case EventPolicyDocArchived:
		return true
	case EventPolicyDocRestored:
		return true
	case EventPolicyDocShared:
		return true
	case EventPolicyDocUnshared:
		return true
	case EventPolicyDocExported:
		return true
	case EventPolicyDocImported:
		return true
	case EventPolicyDocRotated:
		return true
	case EventPolicyDocAccessed:
		return true
	case EventPolicyDocDenied:
		return true
	case EventPolicyDocApproved:
		return true
	case EventPolicyDocRevoked:
		return true
	case EventPolicyDocExpired:
		return true
	case EventPolicyDocLocked:
		return true
	case EventPolicyDocUnlocked:
		return true
	case EventPolicyDocVerified:
		return true
	case EventPolicyDocFailed:
		return true
	case EventPolicyDocSucceeded:
		return true
	default:
		return false
	}
}


// NewPolicyDocCreatedEvent builds a domain event.
func NewPolicyDocCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocUpdatedEvent builds a domain event.
func NewPolicyDocUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocDeletedEvent builds a domain event.
func NewPolicyDocDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocArchivedEvent builds a domain event.
func NewPolicyDocArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocRestoredEvent builds a domain event.
func NewPolicyDocRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocSharedEvent builds a domain event.
func NewPolicyDocSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocUnsharedEvent builds a domain event.
func NewPolicyDocUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocExportedEvent builds a domain event.
func NewPolicyDocExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocImportedEvent builds a domain event.
func NewPolicyDocImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocRotatedEvent builds a domain event.
func NewPolicyDocRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocAccessedEvent builds a domain event.
func NewPolicyDocAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocDeniedEvent builds a domain event.
func NewPolicyDocDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocApprovedEvent builds a domain event.
func NewPolicyDocApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocRevokedEvent builds a domain event.
func NewPolicyDocRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocExpiredEvent builds a domain event.
func NewPolicyDocExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocLockedEvent builds a domain event.
func NewPolicyDocLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocUnlockedEvent builds a domain event.
func NewPolicyDocUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocVerifiedEvent builds a domain event.
func NewPolicyDocVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocFailedEvent builds a domain event.
func NewPolicyDocFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewPolicyDocSucceededEvent builds a domain event.
func NewPolicyDocSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "PolicyDoc"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventPolicyDocSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownPolicyDocEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for authorization-service.
const (
	EventAuthzDecisionCreated   = "authorization.created"
	EventAuthzDecisionUpdated   = "authorization.updated"
	EventAuthzDecisionDeleted   = "authorization.deleted"
	EventAuthzDecisionArchived  = "authorization.archived"
	EventAuthzDecisionRestored  = "authorization.restored"
	EventAuthzDecisionShared    = "authorization.shared"
	EventAuthzDecisionUnshared  = "authorization.unshared"
	EventAuthzDecisionExported  = "authorization.exported"
	EventAuthzDecisionImported  = "authorization.imported"
	EventAuthzDecisionRotated   = "authorization.rotated"
	EventAuthzDecisionAccessed  = "authorization.accessed"
	EventAuthzDecisionDenied    = "authorization.denied"
	EventAuthzDecisionApproved  = "authorization.approved"
	EventAuthzDecisionRevoked   = "authorization.revoked"
	EventAuthzDecisionExpired   = "authorization.expired"
	EventAuthzDecisionLocked    = "authorization.locked"
	EventAuthzDecisionUnlocked  = "authorization.unlocked"
	EventAuthzDecisionVerified  = "authorization.verified"
	EventAuthzDecisionFailed    = "authorization.failed"
	EventAuthzDecisionSucceeded = "authorization.succeeded"
)

// Event is a authorization-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownAuthzDecisionEvent reports whether t is a known authorization event type.
func IsKnownAuthzDecisionEvent(t string) bool {
	switch t {
	case EventAuthzDecisionCreated:
		return true
	case EventAuthzDecisionUpdated:
		return true
	case EventAuthzDecisionDeleted:
		return true
	case EventAuthzDecisionArchived:
		return true
	case EventAuthzDecisionRestored:
		return true
	case EventAuthzDecisionShared:
		return true
	case EventAuthzDecisionUnshared:
		return true
	case EventAuthzDecisionExported:
		return true
	case EventAuthzDecisionImported:
		return true
	case EventAuthzDecisionRotated:
		return true
	case EventAuthzDecisionAccessed:
		return true
	case EventAuthzDecisionDenied:
		return true
	case EventAuthzDecisionApproved:
		return true
	case EventAuthzDecisionRevoked:
		return true
	case EventAuthzDecisionExpired:
		return true
	case EventAuthzDecisionLocked:
		return true
	case EventAuthzDecisionUnlocked:
		return true
	case EventAuthzDecisionVerified:
		return true
	case EventAuthzDecisionFailed:
		return true
	case EventAuthzDecisionSucceeded:
		return true
	default:
		return false
	}
}

// NewAuthzDecisionCreatedEvent builds a domain event.
func NewAuthzDecisionCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionUpdatedEvent builds a domain event.
func NewAuthzDecisionUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionDeletedEvent builds a domain event.
func NewAuthzDecisionDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionArchivedEvent builds a domain event.
func NewAuthzDecisionArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionRestoredEvent builds a domain event.
func NewAuthzDecisionRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionSharedEvent builds a domain event.
func NewAuthzDecisionSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionUnsharedEvent builds a domain event.
func NewAuthzDecisionUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionExportedEvent builds a domain event.
func NewAuthzDecisionExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionImportedEvent builds a domain event.
func NewAuthzDecisionImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionRotatedEvent builds a domain event.
func NewAuthzDecisionRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionAccessedEvent builds a domain event.
func NewAuthzDecisionAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionDeniedEvent builds a domain event.
func NewAuthzDecisionDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionApprovedEvent builds a domain event.
func NewAuthzDecisionApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionRevokedEvent builds a domain event.
func NewAuthzDecisionRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionExpiredEvent builds a domain event.
func NewAuthzDecisionExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionLockedEvent builds a domain event.
func NewAuthzDecisionLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionUnlockedEvent builds a domain event.
func NewAuthzDecisionUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionVerifiedEvent builds a domain event.
func NewAuthzDecisionVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionFailedEvent builds a domain event.
func NewAuthzDecisionFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewAuthzDecisionSucceededEvent builds a domain event.
func NewAuthzDecisionSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "AuthzDecision"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventAuthzDecisionSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownAuthzDecisionEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

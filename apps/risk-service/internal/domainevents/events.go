package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for risk-service.
const (
	EventRiskScoreCreated   = "risk.created"
	EventRiskScoreUpdated   = "risk.updated"
	EventRiskScoreDeleted   = "risk.deleted"
	EventRiskScoreArchived  = "risk.archived"
	EventRiskScoreRestored  = "risk.restored"
	EventRiskScoreShared    = "risk.shared"
	EventRiskScoreUnshared  = "risk.unshared"
	EventRiskScoreExported  = "risk.exported"
	EventRiskScoreImported  = "risk.imported"
	EventRiskScoreRotated   = "risk.rotated"
	EventRiskScoreAccessed  = "risk.accessed"
	EventRiskScoreDenied    = "risk.denied"
	EventRiskScoreApproved  = "risk.approved"
	EventRiskScoreRevoked   = "risk.revoked"
	EventRiskScoreExpired   = "risk.expired"
	EventRiskScoreLocked    = "risk.locked"
	EventRiskScoreUnlocked  = "risk.unlocked"
	EventRiskScoreVerified  = "risk.verified"
	EventRiskScoreFailed    = "risk.failed"
	EventRiskScoreSucceeded = "risk.succeeded"
)

// Event is a risk-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownRiskScoreEvent reports whether t is a known risk event type.
func IsKnownRiskScoreEvent(t string) bool {
	switch t {
	case EventRiskScoreCreated:
		return true
	case EventRiskScoreUpdated:
		return true
	case EventRiskScoreDeleted:
		return true
	case EventRiskScoreArchived:
		return true
	case EventRiskScoreRestored:
		return true
	case EventRiskScoreShared:
		return true
	case EventRiskScoreUnshared:
		return true
	case EventRiskScoreExported:
		return true
	case EventRiskScoreImported:
		return true
	case EventRiskScoreRotated:
		return true
	case EventRiskScoreAccessed:
		return true
	case EventRiskScoreDenied:
		return true
	case EventRiskScoreApproved:
		return true
	case EventRiskScoreRevoked:
		return true
	case EventRiskScoreExpired:
		return true
	case EventRiskScoreLocked:
		return true
	case EventRiskScoreUnlocked:
		return true
	case EventRiskScoreVerified:
		return true
	case EventRiskScoreFailed:
		return true
	case EventRiskScoreSucceeded:
		return true
	default:
		return false
	}
}

// NewRiskScoreCreatedEvent builds a domain event.
func NewRiskScoreCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreUpdatedEvent builds a domain event.
func NewRiskScoreUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreDeletedEvent builds a domain event.
func NewRiskScoreDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreArchivedEvent builds a domain event.
func NewRiskScoreArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreRestoredEvent builds a domain event.
func NewRiskScoreRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreSharedEvent builds a domain event.
func NewRiskScoreSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreUnsharedEvent builds a domain event.
func NewRiskScoreUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreExportedEvent builds a domain event.
func NewRiskScoreExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreImportedEvent builds a domain event.
func NewRiskScoreImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreRotatedEvent builds a domain event.
func NewRiskScoreRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreAccessedEvent builds a domain event.
func NewRiskScoreAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreDeniedEvent builds a domain event.
func NewRiskScoreDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreApprovedEvent builds a domain event.
func NewRiskScoreApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreRevokedEvent builds a domain event.
func NewRiskScoreRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreExpiredEvent builds a domain event.
func NewRiskScoreExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreLockedEvent builds a domain event.
func NewRiskScoreLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreUnlockedEvent builds a domain event.
func NewRiskScoreUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreVerifiedEvent builds a domain event.
func NewRiskScoreVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreFailedEvent builds a domain event.
func NewRiskScoreFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRiskScoreSucceededEvent builds a domain event.
func NewRiskScoreSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RiskScore"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventRiskScoreSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownRiskScoreEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

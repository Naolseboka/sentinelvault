package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for team-service.
const (
	EventTeamCreated = "team.created"
	EventTeamUpdated = "team.updated"
	EventTeamDeleted = "team.deleted"
	EventTeamArchived = "team.archived"
	EventTeamRestored = "team.restored"
	EventTeamShared = "team.shared"
	EventTeamUnshared = "team.unshared"
	EventTeamExported = "team.exported"
	EventTeamImported = "team.imported"
	EventTeamRotated = "team.rotated"
	EventTeamAccessed = "team.accessed"
	EventTeamDenied = "team.denied"
	EventTeamApproved = "team.approved"
	EventTeamRevoked = "team.revoked"
	EventTeamExpired = "team.expired"
	EventTeamLocked = "team.locked"
	EventTeamUnlocked = "team.unlocked"
	EventTeamVerified = "team.verified"
	EventTeamFailed = "team.failed"
	EventTeamSucceeded = "team.succeeded"
)

// Event is a team-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownTeamEvent reports whether t is a known team event type.
func IsKnownTeamEvent(t string) bool {
	switch t {
	case EventTeamCreated:
		return true
	case EventTeamUpdated:
		return true
	case EventTeamDeleted:
		return true
	case EventTeamArchived:
		return true
	case EventTeamRestored:
		return true
	case EventTeamShared:
		return true
	case EventTeamUnshared:
		return true
	case EventTeamExported:
		return true
	case EventTeamImported:
		return true
	case EventTeamRotated:
		return true
	case EventTeamAccessed:
		return true
	case EventTeamDenied:
		return true
	case EventTeamApproved:
		return true
	case EventTeamRevoked:
		return true
	case EventTeamExpired:
		return true
	case EventTeamLocked:
		return true
	case EventTeamUnlocked:
		return true
	case EventTeamVerified:
		return true
	case EventTeamFailed:
		return true
	case EventTeamSucceeded:
		return true
	default:
		return false
	}
}


// NewTeamCreatedEvent builds a domain event.
func NewTeamCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventTeamCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamUpdatedEvent builds a domain event.
func NewTeamUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventTeamUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamDeletedEvent builds a domain event.
func NewTeamDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventTeamDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamArchivedEvent builds a domain event.
func NewTeamArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventTeamArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamRestoredEvent builds a domain event.
func NewTeamRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventTeamRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamSharedEvent builds a domain event.
func NewTeamSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventTeamShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamUnsharedEvent builds a domain event.
func NewTeamUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventTeamUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamExportedEvent builds a domain event.
func NewTeamExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventTeamExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamImportedEvent builds a domain event.
func NewTeamImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventTeamImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamRotatedEvent builds a domain event.
func NewTeamRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventTeamRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamAccessedEvent builds a domain event.
func NewTeamAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventTeamAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamDeniedEvent builds a domain event.
func NewTeamDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventTeamDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamApprovedEvent builds a domain event.
func NewTeamApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventTeamApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamRevokedEvent builds a domain event.
func NewTeamRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventTeamRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamExpiredEvent builds a domain event.
func NewTeamExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventTeamExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamLockedEvent builds a domain event.
func NewTeamLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventTeamLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamUnlockedEvent builds a domain event.
func NewTeamUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventTeamUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamVerifiedEvent builds a domain event.
func NewTeamVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventTeamVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamFailedEvent builds a domain event.
func NewTeamFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventTeamFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewTeamSucceededEvent builds a domain event.
func NewTeamSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Team"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventTeamSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownTeamEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

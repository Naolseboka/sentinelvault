package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for rotation-service.
const (
	EventRotationPolicyCreated = "rotation.created"
	EventRotationPolicyUpdated = "rotation.updated"
	EventRotationPolicyDeleted = "rotation.deleted"
	EventRotationPolicyArchived = "rotation.archived"
	EventRotationPolicyRestored = "rotation.restored"
	EventRotationPolicyShared = "rotation.shared"
	EventRotationPolicyUnshared = "rotation.unshared"
	EventRotationPolicyExported = "rotation.exported"
	EventRotationPolicyImported = "rotation.imported"
	EventRotationPolicyRotated = "rotation.rotated"
	EventRotationPolicyAccessed = "rotation.accessed"
	EventRotationPolicyDenied = "rotation.denied"
	EventRotationPolicyApproved = "rotation.approved"
	EventRotationPolicyRevoked = "rotation.revoked"
	EventRotationPolicyExpired = "rotation.expired"
	EventRotationPolicyLocked = "rotation.locked"
	EventRotationPolicyUnlocked = "rotation.unlocked"
	EventRotationPolicyVerified = "rotation.verified"
	EventRotationPolicyFailed = "rotation.failed"
	EventRotationPolicySucceeded = "rotation.succeeded"
)

// Event is a rotation-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownRotationPolicyEvent reports whether t is a known rotation event type.
func IsKnownRotationPolicyEvent(t string) bool {
	switch t {
	case EventRotationPolicyCreated:
		return true
	case EventRotationPolicyUpdated:
		return true
	case EventRotationPolicyDeleted:
		return true
	case EventRotationPolicyArchived:
		return true
	case EventRotationPolicyRestored:
		return true
	case EventRotationPolicyShared:
		return true
	case EventRotationPolicyUnshared:
		return true
	case EventRotationPolicyExported:
		return true
	case EventRotationPolicyImported:
		return true
	case EventRotationPolicyRotated:
		return true
	case EventRotationPolicyAccessed:
		return true
	case EventRotationPolicyDenied:
		return true
	case EventRotationPolicyApproved:
		return true
	case EventRotationPolicyRevoked:
		return true
	case EventRotationPolicyExpired:
		return true
	case EventRotationPolicyLocked:
		return true
	case EventRotationPolicyUnlocked:
		return true
	case EventRotationPolicyVerified:
		return true
	case EventRotationPolicyFailed:
		return true
	case EventRotationPolicySucceeded:
		return true
	default:
		return false
	}
}


// NewRotationPolicyCreatedEvent builds a domain event.
func NewRotationPolicyCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyUpdatedEvent builds a domain event.
func NewRotationPolicyUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyDeletedEvent builds a domain event.
func NewRotationPolicyDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyArchivedEvent builds a domain event.
func NewRotationPolicyArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyRestoredEvent builds a domain event.
func NewRotationPolicyRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicySharedEvent builds a domain event.
func NewRotationPolicySharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyUnsharedEvent builds a domain event.
func NewRotationPolicyUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyExportedEvent builds a domain event.
func NewRotationPolicyExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyImportedEvent builds a domain event.
func NewRotationPolicyImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyRotatedEvent builds a domain event.
func NewRotationPolicyRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyAccessedEvent builds a domain event.
func NewRotationPolicyAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyDeniedEvent builds a domain event.
func NewRotationPolicyDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyApprovedEvent builds a domain event.
func NewRotationPolicyApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyRevokedEvent builds a domain event.
func NewRotationPolicyRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyExpiredEvent builds a domain event.
func NewRotationPolicyExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyLockedEvent builds a domain event.
func NewRotationPolicyLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyUnlockedEvent builds a domain event.
func NewRotationPolicyUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyVerifiedEvent builds a domain event.
func NewRotationPolicyVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicyFailedEvent builds a domain event.
func NewRotationPolicyFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicyFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewRotationPolicySucceededEvent builds a domain event.
func NewRotationPolicySucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "RotationPolicy"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventRotationPolicySucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownRotationPolicyEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

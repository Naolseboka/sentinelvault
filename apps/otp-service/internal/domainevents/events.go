package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for otp-service.
const (
	EventOTPRecordCreated   = "otp.created"
	EventOTPRecordUpdated   = "otp.updated"
	EventOTPRecordDeleted   = "otp.deleted"
	EventOTPRecordArchived  = "otp.archived"
	EventOTPRecordRestored  = "otp.restored"
	EventOTPRecordShared    = "otp.shared"
	EventOTPRecordUnshared  = "otp.unshared"
	EventOTPRecordExported  = "otp.exported"
	EventOTPRecordImported  = "otp.imported"
	EventOTPRecordRotated   = "otp.rotated"
	EventOTPRecordAccessed  = "otp.accessed"
	EventOTPRecordDenied    = "otp.denied"
	EventOTPRecordApproved  = "otp.approved"
	EventOTPRecordRevoked   = "otp.revoked"
	EventOTPRecordExpired   = "otp.expired"
	EventOTPRecordLocked    = "otp.locked"
	EventOTPRecordUnlocked  = "otp.unlocked"
	EventOTPRecordVerified  = "otp.verified"
	EventOTPRecordFailed    = "otp.failed"
	EventOTPRecordSucceeded = "otp.succeeded"
)

// Event is a otp-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownOTPRecordEvent reports whether t is a known otp event type.
func IsKnownOTPRecordEvent(t string) bool {
	switch t {
	case EventOTPRecordCreated:
		return true
	case EventOTPRecordUpdated:
		return true
	case EventOTPRecordDeleted:
		return true
	case EventOTPRecordArchived:
		return true
	case EventOTPRecordRestored:
		return true
	case EventOTPRecordShared:
		return true
	case EventOTPRecordUnshared:
		return true
	case EventOTPRecordExported:
		return true
	case EventOTPRecordImported:
		return true
	case EventOTPRecordRotated:
		return true
	case EventOTPRecordAccessed:
		return true
	case EventOTPRecordDenied:
		return true
	case EventOTPRecordApproved:
		return true
	case EventOTPRecordRevoked:
		return true
	case EventOTPRecordExpired:
		return true
	case EventOTPRecordLocked:
		return true
	case EventOTPRecordUnlocked:
		return true
	case EventOTPRecordVerified:
		return true
	case EventOTPRecordFailed:
		return true
	case EventOTPRecordSucceeded:
		return true
	default:
		return false
	}
}

// NewOTPRecordCreatedEvent builds a domain event.
func NewOTPRecordCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordUpdatedEvent builds a domain event.
func NewOTPRecordUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordDeletedEvent builds a domain event.
func NewOTPRecordDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordArchivedEvent builds a domain event.
func NewOTPRecordArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordRestoredEvent builds a domain event.
func NewOTPRecordRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordSharedEvent builds a domain event.
func NewOTPRecordSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordUnsharedEvent builds a domain event.
func NewOTPRecordUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordExportedEvent builds a domain event.
func NewOTPRecordExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordImportedEvent builds a domain event.
func NewOTPRecordImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordRotatedEvent builds a domain event.
func NewOTPRecordRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordAccessedEvent builds a domain event.
func NewOTPRecordAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordDeniedEvent builds a domain event.
func NewOTPRecordDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordApprovedEvent builds a domain event.
func NewOTPRecordApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordRevokedEvent builds a domain event.
func NewOTPRecordRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordExpiredEvent builds a domain event.
func NewOTPRecordExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordLockedEvent builds a domain event.
func NewOTPRecordLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordUnlockedEvent builds a domain event.
func NewOTPRecordUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordVerifiedEvent builds a domain event.
func NewOTPRecordVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordFailedEvent builds a domain event.
func NewOTPRecordFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewOTPRecordSucceededEvent builds a domain event.
func NewOTPRecordSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "OTPRecord"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventOTPRecordSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownOTPRecordEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

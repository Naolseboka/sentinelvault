package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for reporting-service.
const (
	EventReportCreated   = "reporting.created"
	EventReportUpdated   = "reporting.updated"
	EventReportDeleted   = "reporting.deleted"
	EventReportArchived  = "reporting.archived"
	EventReportRestored  = "reporting.restored"
	EventReportShared    = "reporting.shared"
	EventReportUnshared  = "reporting.unshared"
	EventReportExported  = "reporting.exported"
	EventReportImported  = "reporting.imported"
	EventReportRotated   = "reporting.rotated"
	EventReportAccessed  = "reporting.accessed"
	EventReportDenied    = "reporting.denied"
	EventReportApproved  = "reporting.approved"
	EventReportRevoked   = "reporting.revoked"
	EventReportExpired   = "reporting.expired"
	EventReportLocked    = "reporting.locked"
	EventReportUnlocked  = "reporting.unlocked"
	EventReportVerified  = "reporting.verified"
	EventReportFailed    = "reporting.failed"
	EventReportSucceeded = "reporting.succeeded"
)

// Event is a reporting-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownReportEvent reports whether t is a known reporting event type.
func IsKnownReportEvent(t string) bool {
	switch t {
	case EventReportCreated:
		return true
	case EventReportUpdated:
		return true
	case EventReportDeleted:
		return true
	case EventReportArchived:
		return true
	case EventReportRestored:
		return true
	case EventReportShared:
		return true
	case EventReportUnshared:
		return true
	case EventReportExported:
		return true
	case EventReportImported:
		return true
	case EventReportRotated:
		return true
	case EventReportAccessed:
		return true
	case EventReportDenied:
		return true
	case EventReportApproved:
		return true
	case EventReportRevoked:
		return true
	case EventReportExpired:
		return true
	case EventReportLocked:
		return true
	case EventReportUnlocked:
		return true
	case EventReportVerified:
		return true
	case EventReportFailed:
		return true
	case EventReportSucceeded:
		return true
	default:
		return false
	}
}

// NewReportCreatedEvent builds a domain event.
func NewReportCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventReportCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportUpdatedEvent builds a domain event.
func NewReportUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventReportUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportDeletedEvent builds a domain event.
func NewReportDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventReportDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportArchivedEvent builds a domain event.
func NewReportArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventReportArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportRestoredEvent builds a domain event.
func NewReportRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventReportRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportSharedEvent builds a domain event.
func NewReportSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventReportShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportUnsharedEvent builds a domain event.
func NewReportUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventReportUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportExportedEvent builds a domain event.
func NewReportExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventReportExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportImportedEvent builds a domain event.
func NewReportImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventReportImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportRotatedEvent builds a domain event.
func NewReportRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventReportRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportAccessedEvent builds a domain event.
func NewReportAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventReportAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportDeniedEvent builds a domain event.
func NewReportDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventReportDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportApprovedEvent builds a domain event.
func NewReportApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventReportApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportRevokedEvent builds a domain event.
func NewReportRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventReportRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportExpiredEvent builds a domain event.
func NewReportExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventReportExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportLockedEvent builds a domain event.
func NewReportLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventReportLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportUnlockedEvent builds a domain event.
func NewReportUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventReportUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportVerifiedEvent builds a domain event.
func NewReportVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventReportVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportFailedEvent builds a domain event.
func NewReportFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventReportFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewReportSucceededEvent builds a domain event.
func NewReportSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Report"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventReportSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownReportEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

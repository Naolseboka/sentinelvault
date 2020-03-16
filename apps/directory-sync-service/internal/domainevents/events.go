package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for directory-sync-service.
const (
	EventSyncJobCreated = "directorysync.created"
	EventSyncJobUpdated = "directorysync.updated"
	EventSyncJobDeleted = "directorysync.deleted"
	EventSyncJobArchived = "directorysync.archived"
	EventSyncJobRestored = "directorysync.restored"
	EventSyncJobShared = "directorysync.shared"
	EventSyncJobUnshared = "directorysync.unshared"
	EventSyncJobExported = "directorysync.exported"
	EventSyncJobImported = "directorysync.imported"
	EventSyncJobRotated = "directorysync.rotated"
	EventSyncJobAccessed = "directorysync.accessed"
	EventSyncJobDenied = "directorysync.denied"
	EventSyncJobApproved = "directorysync.approved"
	EventSyncJobRevoked = "directorysync.revoked"
	EventSyncJobExpired = "directorysync.expired"
	EventSyncJobLocked = "directorysync.locked"
	EventSyncJobUnlocked = "directorysync.unlocked"
	EventSyncJobVerified = "directorysync.verified"
	EventSyncJobFailed = "directorysync.failed"
	EventSyncJobSucceeded = "directorysync.succeeded"
)

// Event is a directory-sync-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownSyncJobEvent reports whether t is a known directorysync event type.
func IsKnownSyncJobEvent(t string) bool {
	switch t {
	case EventSyncJobCreated:
		return true
	case EventSyncJobUpdated:
		return true
	case EventSyncJobDeleted:
		return true
	case EventSyncJobArchived:
		return true
	case EventSyncJobRestored:
		return true
	case EventSyncJobShared:
		return true
	case EventSyncJobUnshared:
		return true
	case EventSyncJobExported:
		return true
	case EventSyncJobImported:
		return true
	case EventSyncJobRotated:
		return true
	case EventSyncJobAccessed:
		return true
	case EventSyncJobDenied:
		return true
	case EventSyncJobApproved:
		return true
	case EventSyncJobRevoked:
		return true
	case EventSyncJobExpired:
		return true
	case EventSyncJobLocked:
		return true
	case EventSyncJobUnlocked:
		return true
	case EventSyncJobVerified:
		return true
	case EventSyncJobFailed:
		return true
	case EventSyncJobSucceeded:
		return true
	default:
		return false
	}
}


// NewSyncJobCreatedEvent builds a domain event.
func NewSyncJobCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventSyncJobCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobUpdatedEvent builds a domain event.
func NewSyncJobUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventSyncJobUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobDeletedEvent builds a domain event.
func NewSyncJobDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventSyncJobDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobArchivedEvent builds a domain event.
func NewSyncJobArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventSyncJobArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobRestoredEvent builds a domain event.
func NewSyncJobRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventSyncJobRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobSharedEvent builds a domain event.
func NewSyncJobSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventSyncJobShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobUnsharedEvent builds a domain event.
func NewSyncJobUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventSyncJobUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobExportedEvent builds a domain event.
func NewSyncJobExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventSyncJobExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobImportedEvent builds a domain event.
func NewSyncJobImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventSyncJobImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobRotatedEvent builds a domain event.
func NewSyncJobRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventSyncJobRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobAccessedEvent builds a domain event.
func NewSyncJobAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventSyncJobAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobDeniedEvent builds a domain event.
func NewSyncJobDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventSyncJobDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobApprovedEvent builds a domain event.
func NewSyncJobApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventSyncJobApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobRevokedEvent builds a domain event.
func NewSyncJobRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventSyncJobRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobExpiredEvent builds a domain event.
func NewSyncJobExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventSyncJobExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobLockedEvent builds a domain event.
func NewSyncJobLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventSyncJobLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobUnlockedEvent builds a domain event.
func NewSyncJobUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventSyncJobUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobVerifiedEvent builds a domain event.
func NewSyncJobVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventSyncJobVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobFailedEvent builds a domain event.
func NewSyncJobFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventSyncJobFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSyncJobSucceededEvent builds a domain event.
func NewSyncJobSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SyncJob"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventSyncJobSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownSyncJobEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

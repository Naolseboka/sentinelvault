package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for ssh-service.
const (
	EventSSHKeyCreated   = "ssh.created"
	EventSSHKeyUpdated   = "ssh.updated"
	EventSSHKeyDeleted   = "ssh.deleted"
	EventSSHKeyArchived  = "ssh.archived"
	EventSSHKeyRestored  = "ssh.restored"
	EventSSHKeyShared    = "ssh.shared"
	EventSSHKeyUnshared  = "ssh.unshared"
	EventSSHKeyExported  = "ssh.exported"
	EventSSHKeyImported  = "ssh.imported"
	EventSSHKeyRotated   = "ssh.rotated"
	EventSSHKeyAccessed  = "ssh.accessed"
	EventSSHKeyDenied    = "ssh.denied"
	EventSSHKeyApproved  = "ssh.approved"
	EventSSHKeyRevoked   = "ssh.revoked"
	EventSSHKeyExpired   = "ssh.expired"
	EventSSHKeyLocked    = "ssh.locked"
	EventSSHKeyUnlocked  = "ssh.unlocked"
	EventSSHKeyVerified  = "ssh.verified"
	EventSSHKeyFailed    = "ssh.failed"
	EventSSHKeySucceeded = "ssh.succeeded"
)

// Event is a ssh-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownSSHKeyEvent reports whether t is a known ssh event type.
func IsKnownSSHKeyEvent(t string) bool {
	switch t {
	case EventSSHKeyCreated:
		return true
	case EventSSHKeyUpdated:
		return true
	case EventSSHKeyDeleted:
		return true
	case EventSSHKeyArchived:
		return true
	case EventSSHKeyRestored:
		return true
	case EventSSHKeyShared:
		return true
	case EventSSHKeyUnshared:
		return true
	case EventSSHKeyExported:
		return true
	case EventSSHKeyImported:
		return true
	case EventSSHKeyRotated:
		return true
	case EventSSHKeyAccessed:
		return true
	case EventSSHKeyDenied:
		return true
	case EventSSHKeyApproved:
		return true
	case EventSSHKeyRevoked:
		return true
	case EventSSHKeyExpired:
		return true
	case EventSSHKeyLocked:
		return true
	case EventSSHKeyUnlocked:
		return true
	case EventSSHKeyVerified:
		return true
	case EventSSHKeyFailed:
		return true
	case EventSSHKeySucceeded:
		return true
	default:
		return false
	}
}

// NewSSHKeyCreatedEvent builds a domain event.
func NewSSHKeyCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyUpdatedEvent builds a domain event.
func NewSSHKeyUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyDeletedEvent builds a domain event.
func NewSSHKeyDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyArchivedEvent builds a domain event.
func NewSSHKeyArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyRestoredEvent builds a domain event.
func NewSSHKeyRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeySharedEvent builds a domain event.
func NewSSHKeySharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyUnsharedEvent builds a domain event.
func NewSSHKeyUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyExportedEvent builds a domain event.
func NewSSHKeyExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyImportedEvent builds a domain event.
func NewSSHKeyImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyRotatedEvent builds a domain event.
func NewSSHKeyRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyAccessedEvent builds a domain event.
func NewSSHKeyAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyDeniedEvent builds a domain event.
func NewSSHKeyDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyApprovedEvent builds a domain event.
func NewSSHKeyApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyRevokedEvent builds a domain event.
func NewSSHKeyRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyExpiredEvent builds a domain event.
func NewSSHKeyExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyLockedEvent builds a domain event.
func NewSSHKeyLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyUnlockedEvent builds a domain event.
func NewSSHKeyUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyVerifiedEvent builds a domain event.
func NewSSHKeyVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeyFailedEvent builds a domain event.
func NewSSHKeyFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventSSHKeyFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewSSHKeySucceededEvent builds a domain event.
func NewSSHKeySucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "SSHKey"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventSSHKeySucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownSSHKeyEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for device-service.
const (
	EventDeviceCreated = "device.created"
	EventDeviceUpdated = "device.updated"
	EventDeviceDeleted = "device.deleted"
	EventDeviceArchived = "device.archived"
	EventDeviceRestored = "device.restored"
	EventDeviceShared = "device.shared"
	EventDeviceUnshared = "device.unshared"
	EventDeviceExported = "device.exported"
	EventDeviceImported = "device.imported"
	EventDeviceRotated = "device.rotated"
	EventDeviceAccessed = "device.accessed"
	EventDeviceDenied = "device.denied"
	EventDeviceApproved = "device.approved"
	EventDeviceRevoked = "device.revoked"
	EventDeviceExpired = "device.expired"
	EventDeviceLocked = "device.locked"
	EventDeviceUnlocked = "device.unlocked"
	EventDeviceVerified = "device.verified"
	EventDeviceFailed = "device.failed"
	EventDeviceSucceeded = "device.succeeded"
)

// Event is a device-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownDeviceEvent reports whether t is a known device event type.
func IsKnownDeviceEvent(t string) bool {
	switch t {
	case EventDeviceCreated:
		return true
	case EventDeviceUpdated:
		return true
	case EventDeviceDeleted:
		return true
	case EventDeviceArchived:
		return true
	case EventDeviceRestored:
		return true
	case EventDeviceShared:
		return true
	case EventDeviceUnshared:
		return true
	case EventDeviceExported:
		return true
	case EventDeviceImported:
		return true
	case EventDeviceRotated:
		return true
	case EventDeviceAccessed:
		return true
	case EventDeviceDenied:
		return true
	case EventDeviceApproved:
		return true
	case EventDeviceRevoked:
		return true
	case EventDeviceExpired:
		return true
	case EventDeviceLocked:
		return true
	case EventDeviceUnlocked:
		return true
	case EventDeviceVerified:
		return true
	case EventDeviceFailed:
		return true
	case EventDeviceSucceeded:
		return true
	default:
		return false
	}
}


// NewDeviceCreatedEvent builds a domain event.
func NewDeviceCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventDeviceCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceUpdatedEvent builds a domain event.
func NewDeviceUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventDeviceUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceDeletedEvent builds a domain event.
func NewDeviceDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventDeviceDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceArchivedEvent builds a domain event.
func NewDeviceArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventDeviceArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceRestoredEvent builds a domain event.
func NewDeviceRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventDeviceRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceSharedEvent builds a domain event.
func NewDeviceSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventDeviceShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceUnsharedEvent builds a domain event.
func NewDeviceUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventDeviceUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceExportedEvent builds a domain event.
func NewDeviceExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventDeviceExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceImportedEvent builds a domain event.
func NewDeviceImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventDeviceImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceRotatedEvent builds a domain event.
func NewDeviceRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventDeviceRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceAccessedEvent builds a domain event.
func NewDeviceAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventDeviceAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceDeniedEvent builds a domain event.
func NewDeviceDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventDeviceDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceApprovedEvent builds a domain event.
func NewDeviceApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventDeviceApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceRevokedEvent builds a domain event.
func NewDeviceRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventDeviceRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceExpiredEvent builds a domain event.
func NewDeviceExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventDeviceExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceLockedEvent builds a domain event.
func NewDeviceLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventDeviceLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceUnlockedEvent builds a domain event.
func NewDeviceUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventDeviceUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceVerifiedEvent builds a domain event.
func NewDeviceVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventDeviceVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceFailedEvent builds a domain event.
func NewDeviceFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventDeviceFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewDeviceSucceededEvent builds a domain event.
func NewDeviceSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Device"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventDeviceSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownDeviceEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}

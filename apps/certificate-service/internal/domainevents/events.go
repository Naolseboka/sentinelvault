package domainevents

import (
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for certificate-service.
const (
	EventCertificateCreated = "certificate.created"
	EventCertificateUpdated = "certificate.updated"
	EventCertificateDeleted = "certificate.deleted"
	EventCertificateArchived = "certificate.archived"
	EventCertificateRestored = "certificate.restored"
	EventCertificateShared = "certificate.shared"
	EventCertificateUnshared = "certificate.unshared"
	EventCertificateExported = "certificate.exported"
	EventCertificateImported = "certificate.imported"
	EventCertificateRotated = "certificate.rotated"
	EventCertificateAccessed = "certificate.accessed"
	EventCertificateDenied = "certificate.denied"
	EventCertificateApproved = "certificate.approved"
	EventCertificateRevoked = "certificate.revoked"
	EventCertificateExpired = "certificate.expired"
	EventCertificateLocked = "certificate.locked"
	EventCertificateUnlocked = "certificate.unlocked"
	EventCertificateVerified = "certificate.verified"
	EventCertificateFailed = "certificate.failed"
	EventCertificateSucceeded = "certificate.succeeded"
)

// Event is a certificate-service domain event.
type Event struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	OrgID      string            `json:"org_id"`
	ActorID    string            `json:"actor_id"`
	ResourceID string            `json:"resource_id"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// IsKnownCertificateEvent reports whether t is a known certificate event type.
func IsKnownCertificateEvent(t string) bool {
	switch t {
	case EventCertificateCreated:
		return true
	case EventCertificateUpdated:
		return true
	case EventCertificateDeleted:
		return true
	case EventCertificateArchived:
		return true
	case EventCertificateRestored:
		return true
	case EventCertificateShared:
		return true
	case EventCertificateUnshared:
		return true
	case EventCertificateExported:
		return true
	case EventCertificateImported:
		return true
	case EventCertificateRotated:
		return true
	case EventCertificateAccessed:
		return true
	case EventCertificateDenied:
		return true
	case EventCertificateApproved:
		return true
	case EventCertificateRevoked:
		return true
	case EventCertificateExpired:
		return true
	case EventCertificateLocked:
		return true
	case EventCertificateUnlocked:
		return true
	case EventCertificateVerified:
		return true
	case EventCertificateFailed:
		return true
	case EventCertificateSucceeded:
		return true
	default:
		return false
	}
}


// NewCertificateCreatedEvent builds a domain event.
func NewCertificateCreatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "created"
	return Event{
		ID: common.NewID(), Type: EventCertificateCreated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateUpdatedEvent builds a domain event.
func NewCertificateUpdatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "updated"
	return Event{
		ID: common.NewID(), Type: EventCertificateUpdated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateDeletedEvent builds a domain event.
func NewCertificateDeletedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "deleted"
	return Event{
		ID: common.NewID(), Type: EventCertificateDeleted, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateArchivedEvent builds a domain event.
func NewCertificateArchivedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "archived"
	return Event{
		ID: common.NewID(), Type: EventCertificateArchived, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateRestoredEvent builds a domain event.
func NewCertificateRestoredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "restored"
	return Event{
		ID: common.NewID(), Type: EventCertificateRestored, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateSharedEvent builds a domain event.
func NewCertificateSharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "shared"
	return Event{
		ID: common.NewID(), Type: EventCertificateShared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateUnsharedEvent builds a domain event.
func NewCertificateUnsharedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "unshared"
	return Event{
		ID: common.NewID(), Type: EventCertificateUnshared, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateExportedEvent builds a domain event.
func NewCertificateExportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "exported"
	return Event{
		ID: common.NewID(), Type: EventCertificateExported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateImportedEvent builds a domain event.
func NewCertificateImportedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "imported"
	return Event{
		ID: common.NewID(), Type: EventCertificateImported, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateRotatedEvent builds a domain event.
func NewCertificateRotatedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "rotated"
	return Event{
		ID: common.NewID(), Type: EventCertificateRotated, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateAccessedEvent builds a domain event.
func NewCertificateAccessedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "accessed"
	return Event{
		ID: common.NewID(), Type: EventCertificateAccessed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateDeniedEvent builds a domain event.
func NewCertificateDeniedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "denied"
	return Event{
		ID: common.NewID(), Type: EventCertificateDenied, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateApprovedEvent builds a domain event.
func NewCertificateApprovedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "approved"
	return Event{
		ID: common.NewID(), Type: EventCertificateApproved, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateRevokedEvent builds a domain event.
func NewCertificateRevokedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "revoked"
	return Event{
		ID: common.NewID(), Type: EventCertificateRevoked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateExpiredEvent builds a domain event.
func NewCertificateExpiredEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "expired"
	return Event{
		ID: common.NewID(), Type: EventCertificateExpired, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateLockedEvent builds a domain event.
func NewCertificateLockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "locked"
	return Event{
		ID: common.NewID(), Type: EventCertificateLocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateUnlockedEvent builds a domain event.
func NewCertificateUnlockedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "unlocked"
	return Event{
		ID: common.NewID(), Type: EventCertificateUnlocked, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateVerifiedEvent builds a domain event.
func NewCertificateVerifiedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "verified"
	return Event{
		ID: common.NewID(), Type: EventCertificateVerified, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateFailedEvent builds a domain event.
func NewCertificateFailedEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "failed"
	return Event{
		ID: common.NewID(), Type: EventCertificateFailed, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// NewCertificateSucceededEvent builds a domain event.
func NewCertificateSucceededEvent(orgID, actorID, resourceID string, meta map[string]string) Event {
	if meta == nil {
		meta = map[string]string{}
	}
	meta["entity"] = "Certificate"
	meta["action"] = "succeeded"
	return Event{
		ID: common.NewID(), Type: EventCertificateSucceeded, OrgID: orgID, ActorID: actorID,
		ResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
	}
}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {
	if e.ID == "" || e.Type == "" || e.OrgID == "" {
		return common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
	}
	if !IsKnownCertificateEvent(e.Type) {
		return common.New(common.ErrCodeInvalidArgument, "unknown event type")
	}
	return nil
}
// maintenance note 95

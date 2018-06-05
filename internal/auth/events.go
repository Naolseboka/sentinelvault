package auth

import "time"

// AuthEventType classifies identity lifecycle events.
type AuthEventType string

const (
	EventLoginSucceeded AuthEventType = "auth.login.succeeded"
	EventLoginFailed    AuthEventType = "auth.login.failed"
	EventMFAEnrolled    AuthEventType = "auth.mfa.enrolled"
	EventSessionRevoked AuthEventType = "auth.session.revoked"
)

// AuthEvent is a domain event emitted by identity services.
type AuthEvent struct {
	Type      AuthEventType  `json:"type"`
	UserID    string         `json:"user_id,omitempty"`
	OrgID     string         `json:"org_id,omitempty"`
	SessionID string         `json:"session_id,omitempty"`
	Email     string         `json:"email,omitempty"`
	Reason    string         `json:"reason,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	OccurredAt time.Time     `json:"occurred_at"`
}

// NewAuthEvent creates an auth event with a UTC timestamp.
func NewAuthEvent(t AuthEventType) AuthEvent {
	return AuthEvent{Type: t, OccurredAt: time.Now().UTC()}
}

func (e AuthEvent) WithUser(id string) AuthEvent {
	e.UserID = id
	return e
}

func (e AuthEvent) WithOrg(id string) AuthEvent {
	e.OrgID = id
	return e
}

func (e AuthEvent) WithSession(id string) AuthEvent {
	e.SessionID = id
	return e
}

func (e AuthEvent) WithEmail(email string) AuthEvent {
	e.Email = email
	return e
}

func (e AuthEvent) WithReason(reason string) AuthEvent {
	e.Reason = reason
	return e
}

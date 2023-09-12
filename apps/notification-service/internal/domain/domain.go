package domain

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/notify"
)

// Message re-exports notify.Message.
type Message = notify.Message

// SendInput holds notification dispatch fields.
type SendInput struct {
	OrgID     string            `json:"org_id"`
	Channel   string            `json:"channel"`
	Recipient string            `json:"recipient"`
	Subject   string            `json:"subject"`
	Body      string            `json:"body"`
	Priority  string            `json:"priority"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// Service wraps the notification dispatcher.
type Service struct {
	dispatcher *notify.Dispatcher
}

// NewService creates a notification domain service.
func NewService(d *notify.Dispatcher) *Service {
	return &Service{dispatcher: d}
}

// NewDispatcher creates a dispatcher with default memory senders.
func NewDispatcher() *notify.Dispatcher {
	d := notify.NewDispatcher()
	notify.RegisterDefaults(d)
	return d
}

// Send dispatches a notification.
func (s *Service) Send(ctx context.Context, in SendInput) (*Message, error) {
	return s.dispatcher.Send(ctx, in.OrgID, in.Channel, in.Recipient, in.Subject, in.Body, in.Priority, in.Metadata)
}

// List returns sent notification history.
func (s *Service) List(_ context.Context, orgID string) []Message {
	all := s.dispatcher.Sent()
	if orgID == "" {
		return all
	}
	out := make([]Message, 0)
	for _, m := range all {
		if m.OrgID == orgID {
			out = append(out, m)
		}
	}
	return out
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "notification" }

// TableName returns the primary table name.
func TableName() string { return "messages" }

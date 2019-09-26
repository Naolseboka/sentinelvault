package notify

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Channel identifiers.
const (
	ChannelEmail     = "email"
	ChannelSlack     = "slack"
	ChannelTeams     = "teams"
	ChannelDiscord   = "discord"
	ChannelSMS       = "sms"
	ChannelPush      = "push"
	ChannelWebhook   = "webhook"
	ChannelPagerDuty = "pagerduty"
	ChannelOpsgenie  = "opsgenie"
)

// Message is a notification payload.
type Message struct {
	ID        string            `json:"id"`
	OrgID     string            `json:"org_id"`
	Channel   string            `json:"channel"`
	Recipient string            `json:"recipient"`
	Subject   string            `json:"subject"`
	Body      string            `json:"body"`
	Priority  string            `json:"priority"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
	Status    string            `json:"status"`
}

// Sender delivers a message on a channel.
type Sender interface {
	Channel() string
	Send(ctx context.Context, msg Message) error
}

// Dispatcher routes messages to channel senders.
type Dispatcher struct {
	mu      sync.RWMutex
	senders map[string]Sender
	sent    []Message
}

// NewDispatcher creates an empty dispatcher.
func NewDispatcher() *Dispatcher {
	return &Dispatcher{senders: map[string]Sender{}, sent: make([]Message, 0)}
}

// Register adds a sender for a channel.
func (d *Dispatcher) Register(s Sender) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.senders[s.Channel()] = s
}

// Send dispatches a message.
func (d *Dispatcher) Send(ctx context.Context, orgID, channel, recipient, subject, body, priority string, metadata map[string]string) (*Message, error) {
	if orgID == "" || channel == "" || recipient == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, channel, recipient required")
	}
	d.mu.RLock()
	sender, ok := d.senders[channel]
	d.mu.RUnlock()
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, fmt.Sprintf("no sender for channel %q", channel))
	}
	msg := Message{
		ID: common.NewID(), OrgID: orgID, Channel: channel, Recipient: recipient,
		Subject: subject, Body: body, Priority: priority, Metadata: metadata,
		CreatedAt: common.NowUTC(), Status: "pending",
	}
	if err := sender.Send(ctx, msg); err != nil {
		msg.Status = "failed"
		d.mu.Lock()
		d.sent = append(d.sent, msg)
		d.mu.Unlock()
		return &msg, err
	}
	msg.Status = "sent"
	d.mu.Lock()
	d.sent = append(d.sent, msg)
	d.mu.Unlock()
	return &msg, nil
}

// Sent returns delivery history.
func (d *Dispatcher) Sent() []Message {
	d.mu.RLock()
	defer d.mu.RUnlock()
	out := make([]Message, len(d.sent))
	copy(out, d.sent)
	return out
}

// MemorySender records messages without external I/O.
type MemorySender struct {
	channel string
	fail    bool
}

// NewMemorySender creates a test sender for channel.
func NewMemorySender(channel string) *MemorySender {
	return &MemorySender{channel: channel}
}

func (m *MemorySender) Channel() string { return m.channel }

func (m *MemorySender) Send(_ context.Context, _ Message) error {
	if m.fail {
		return fmt.Errorf("send failed")
	}
	return nil
}

// RegisterDefaults registers memory senders for all supported channels.
func RegisterDefaults(d *Dispatcher) {
	for _, ch := range []string{
		ChannelEmail, ChannelSlack, ChannelTeams, ChannelDiscord, ChannelSMS,
		ChannelPush, ChannelWebhook, ChannelPagerDuty, ChannelOpsgenie,
	} {
		d.Register(NewMemorySender(ch))
	}
}

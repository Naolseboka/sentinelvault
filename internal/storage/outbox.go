package storage

import (
	"encoding/json"
	"time"
)

// OutboxStatus tracks delivery state of an outbox message.
type OutboxStatus string

const (
	OutboxPending   OutboxStatus = "pending"
	OutboxPublished OutboxStatus = "published"
	OutboxFailed    OutboxStatus = "failed"
)

// OutboxMessage represents a transactional outbox row.
type OutboxMessage struct {
	ID          string          `json:"id"`
	AggregateID string          `json:"aggregate_id"`
	Topic       string          `json:"topic"`
	Payload     json.RawMessage `json:"payload"`
	Status      OutboxStatus    `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	PublishedAt *time.Time      `json:"published_at,omitempty"`
	Attempts    int             `json:"attempts"`
	LastError   string          `json:"last_error,omitempty"`
}

// OutboxRepository persists outbox messages within transactions.
type OutboxRepository interface {
	Insert(msg OutboxMessage) error
	FetchPending(limit int) ([]OutboxMessage, error)
	MarkPublished(id string, at time.Time) error
	MarkFailed(id string, errMsg string) error
}

package queue

import (
	"context"
	"encoding/json"
	"sync"
	"time"
)

// DomainEvent is a versioned domain event envelope.
type DomainEvent struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	AggregateID   string            `json:"aggregate_id"`
	AggregateType string            `json:"aggregate_type"`
	OccurredAt    time.Time         `json:"occurred_at"`
	Payload       json.RawMessage   `json:"payload"`
	Metadata      map[string]string `json:"metadata,omitempty"`
}

// Producer publishes messages to a topic.
type Producer interface {
	Publish(ctx context.Context, topic string, key string, value []byte) error
	Close() error
}

// Consumer reads messages from a topic.
type Consumer interface {
	Consume(ctx context.Context, topic string, handler func(ctx context.Context, key string, value []byte) error) error
	Close() error
}

// MemoryProducer is an in-memory producer for tests.
type MemoryProducer struct {
	mu     sync.RWMutex
	topics map[string][]MemoryMessage
}

type MemoryMessage struct {
	Key   string
	Value []byte
	Time  time.Time
}

// NewMemoryProducer creates an in-memory producer.
func NewMemoryProducer() *MemoryProducer {
	return &MemoryProducer{topics: map[string][]MemoryMessage{}}
}

func (p *MemoryProducer) Publish(ctx context.Context, topic, key string, value []byte) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	cp := make([]byte, len(value))
	copy(cp, value)
	p.topics[topic] = append(p.topics[topic], MemoryMessage{Key: key, Value: cp, Time: time.Now().UTC()})
	return nil
}

func (p *MemoryProducer) Close() error { return nil }

// Messages returns copies of messages for a topic.
func (p *MemoryProducer) Messages(topic string) []MemoryMessage {
	p.mu.RLock()
	defer p.mu.RUnlock()
	src := p.topics[topic]
	out := make([]MemoryMessage, len(src))
	copy(out, src)
	return out
}

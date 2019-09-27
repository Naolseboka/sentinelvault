package queue

import "context"

// Queue abstracts a work queue (RabbitMQ-compatible surface).
type Queue interface {
	Publish(ctx context.Context, queue string, body []byte) error
	Subscribe(ctx context.Context, queue string, handler func(ctx context.Context, body []byte) error) error
	Close() error
}

// RabbitMQConfig holds connection settings for RabbitMQ.
type RabbitMQConfig struct {
	URL string
}

// RabbitMQQueue is a placeholder queue client; full AMQP wiring lands in Phase 2.
type RabbitMQQueue struct {
	url string
}

// NewRabbitMQQueue creates a queue client handle.
func NewRabbitMQQueue(cfg RabbitMQConfig) *RabbitMQQueue {
	return &RabbitMQQueue{url: cfg.URL}
}

func (q *RabbitMQQueue) Publish(ctx context.Context, queue string, body []byte) error {
	_ = ctx
	_ = queue
	_ = body
	return nil
}

func (q *RabbitMQQueue) Subscribe(ctx context.Context, queue string, handler func(ctx context.Context, body []byte) error) error {
	<-ctx.Done()
	return ctx.Err()
}

func (q *RabbitMQQueue) Close() error { return nil }

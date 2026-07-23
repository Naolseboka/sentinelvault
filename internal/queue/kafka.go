package queue

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

// KafkaProducer publishes to Kafka brokers.
type KafkaProducer struct {
	writers map[string]*kafka.Writer
	mu      sync.Mutex
	brokers []string
}

// NewKafkaProducer creates a producer for the given brokers.
func NewKafkaProducer(brokers []string) *KafkaProducer {
	return &KafkaProducer{brokers: brokers, writers: map[string]*kafka.Writer{}}
}

func (p *KafkaProducer) writer(topic string) *kafka.Writer {
	p.mu.Lock()
	defer p.mu.Unlock()
	if w, ok := p.writers[topic]; ok {
		return w
	}
	w := &kafka.Writer{
		Addr:         kafka.TCP(p.brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
		BatchTimeout: 10 * time.Millisecond,
	}
	p.writers[topic] = w
	return w
}

// Publish sends a message to Kafka.
func (p *KafkaProducer) Publish(ctx context.Context, topic, key string, value []byte) error {
	return p.writer(topic).WriteMessages(ctx, kafka.Message{Key: []byte(key), Value: value})
}

// Close closes all writers.
func (p *KafkaProducer) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	var first error
	for topic, w := range p.writers {
		if err := w.Close(); err != nil && first == nil {
			first = fmt.Errorf("close writer %s: %w", topic, err)
		}
	}
	p.writers = map[string]*kafka.Writer{}
	return first
}

// KafkaConsumer consumes from a Kafka topic.
type KafkaConsumer struct {
	reader *kafka.Reader
}

// NewKafkaConsumer creates a consumer group reader.
func NewKafkaConsumer(brokers []string, topic, groupID string) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:    topic,
			GroupID:  groupID,
			MinBytes: 1,
			MaxBytes: 10e6,
		}),
	}
}

// Consume reads messages until ctx is cancelled.
func (c *KafkaConsumer) Consume(ctx context.Context, topic string, handler func(ctx context.Context, key string, value []byte) error) error {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		if err := handler(ctx, string(msg.Key), msg.Value); err != nil {
			return err
		}
	}
}

func (c *KafkaConsumer) Close() error { return c.reader.Close() }

// maintenance note 64

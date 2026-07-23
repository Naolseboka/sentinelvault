package outbox

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusPublished Status = "published"
	StatusFailed    Status = "failed"
)

type Message struct {
	ID        string            `json:"id"`
	Topic     string            `json:"topic"`
	Key       string            `json:"key"`
	Payload   json.RawMessage   `json:"payload"`
	Headers   map[string]string `json:"headers,omitempty"`
	Status    Status            `json:"status"`
	Attempts  int               `json:"attempts"`
	LastError string            `json:"last_error,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type Store struct {
	mu   sync.Mutex
	msgs map[string]*Message
}

func NewStore() *Store { return &Store{msgs: map[string]*Message{}} }

func (s *Store) Enqueue(topic, key string, payload any, headers map[string]string) (*Message, error) {
	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	if topic == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "topic required")
	}
	now := common.NowUTC()
	m := &Message{ID: common.NewID(), Topic: topic, Key: key, Payload: raw, Headers: headers, Status: StatusPending, CreatedAt: now, UpdatedAt: now}
	s.mu.Lock()
	s.msgs[m.ID] = m
	s.mu.Unlock()
	cp := *m
	return &cp, nil
}

func (s *Store) Pending(limit int) []Message {
	s.mu.Lock()
	defer s.mu.Unlock()
	if limit <= 0 {
		limit = 100
	}
	out := make([]Message, 0, limit)
	for _, m := range s.msgs {
		if m.Status == StatusPending {
			out = append(out, *m)
			if len(out) >= limit {
				break
			}
		}
	}
	return out
}

func (s *Store) MarkPublished(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	m, ok := s.msgs[id]
	if !ok {
		return common.New(common.ErrCodeNotFound, "message not found")
	}
	m.Status = StatusPublished
	m.UpdatedAt = common.NowUTC()
	return nil
}

func (s *Store) MarkFailed(id, errMsg string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	m, ok := s.msgs[id]
	if !ok {
		return common.New(common.ErrCodeNotFound, "message not found")
	}
	m.Status = StatusFailed
	m.Attempts++
	m.LastError = errMsg
	m.UpdatedAt = common.NowUTC()
	return nil
}

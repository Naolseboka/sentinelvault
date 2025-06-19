package idempotency

import (
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Record struct {
	Key string
	Fingerprint string
	Response []byte
	StatusCode int
	ExpiresAt time.Time
}

type Store struct {
	mu sync.Mutex
	data map[string]Record
	ttl time.Duration
}

func NewStore(ttl time.Duration) *Store {
	if ttl <= 0 {
		ttl = 24 * time.Hour
	}
	return &Store{data: map[string]Record{}, ttl: ttl}
}

func (s *Store) Remember(key, fingerprint string, status int, body []byte) error {
	if key == "" {
		return common.New(common.ErrCodeInvalidArgument, "idempotency key required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if existing, ok := s.data[key]; ok && !time.Now().After(existing.ExpiresAt) {
		if existing.Fingerprint != fingerprint {
			return common.New(common.ErrCodeConflict, "idempotency key reuse with different payload")
		}
		return nil
	}
	s.data[key] = Record{Key: key, Fingerprint: fingerprint, Response: append([]byte(nil), body...), StatusCode: status, ExpiresAt: time.Now().Add(s.ttl)}
	return nil
}

func (s *Store) Lookup(key string) (Record, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	rec, ok := s.data[key]
	if !ok || time.Now().After(rec.ExpiresAt) {
		return Record{}, false
	}
	return rec, true
}

func (s *Store) Cleanup(now time.Time) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := 0
	for k, rec := range s.data {
		if now.After(rec.ExpiresAt) {
			delete(s.data, k)
			n++
		}
	}
	return n
}
// maintenance note 63

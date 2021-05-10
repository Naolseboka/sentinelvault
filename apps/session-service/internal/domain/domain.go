package domain

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/auth"
	"github.com/sentinelvault/sentinelvault/internal/common"
)

const defaultRefreshTTL = 7 * 24 * time.Hour

// Session is an authenticated user session.
type Session struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	OrgID       string     `json:"org_id"`
	DeviceID    string     `json:"device_id"`
	RefreshHash string     `json:"-"`
	ExpiresAt   time.Time  `json:"expires_at"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

type Repository interface {
	Create(ctx context.Context, sess *Session) error
	Get(ctx context.Context, id string) (*Session, error)
	GetByRefreshHash(ctx context.Context, hash string) (*Session, error)
	Update(ctx context.Context, sess *Session) error
	ListByUser(ctx context.Context, userID string) ([]Session, error)
}

type Service struct{ repo Repository }

func NewService(repo Repository) *Service { return &Service{repo: repo} }

func (s *Service) Create(ctx context.Context, userID, orgID, deviceID string) (*Session, string, error) {
	raw, err := auth.NewRefreshTokenID()
	if err != nil {
		return nil, "", err
	}
	now := common.NowUTC()
	sess := &Session{
		ID:          common.NewID(),
		UserID:      userID,
		OrgID:       orgID,
		DeviceID:    deviceID,
		RefreshHash: hashRefresh(raw),
		ExpiresAt:   now.Add(defaultRefreshTTL),
		CreatedAt:   now,
	}
	if err := s.repo.Create(ctx, sess); err != nil {
		return nil, "", err
	}
	return sess, raw, nil
}

func (s *Service) Get(ctx context.Context, id string) (*Session, error) {
	sess, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if sess == nil {
		return nil, common.New(common.ErrCodeNotFound, "session not found")
	}
	return sess, nil
}

func (s *Service) ListByUser(ctx context.Context, userID string) ([]Session, error) {
	return s.repo.ListByUser(ctx, userID)
}

func (s *Service) Revoke(ctx context.Context, id string) (*Session, error) {
	sess, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	now := common.NowUTC()
	sess.RevokedAt = &now
	if err := s.repo.Update(ctx, sess); err != nil {
		return nil, err
	}
	return sess, nil
}

func (s *Service) RotateRefresh(ctx context.Context, refreshToken string) (*Session, string, error) {
	hash := hashRefresh(refreshToken)
	sess, err := s.repo.GetByRefreshHash(ctx, hash)
	if err != nil {
		return nil, "", err
	}
	if sess == nil || sess.RevokedAt != nil || time.Now().UTC().After(sess.ExpiresAt) {
		return nil, "", common.New(common.ErrCodeUnauthorized, "invalid refresh token")
	}
	newRaw, err := auth.NewRefreshTokenID()
	if err != nil {
		return nil, "", err
	}
	sess.RefreshHash = hashRefresh(newRaw)
	sess.ExpiresAt = time.Now().UTC().Add(defaultRefreshTTL)
	if err := s.repo.Update(ctx, sess); err != nil {
		return nil, "", err
	}
	return sess, newRaw, nil
}

func hashRefresh(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}

func SchemaName() string { return "sessions" }
func TableName() string  { return "sessions" }

type MemoryRepository struct {
	mu            sync.RWMutex
	byID          map[string]*Session
	byRefreshHash map[string]string
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{byID: map[string]*Session{}, byRefreshHash: map[string]string{}}
}

func (r *MemoryRepository) Create(_ context.Context, sess *Session) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *sess
	r.byID[sess.ID] = &cp
	r.byRefreshHash[sess.RefreshHash] = sess.ID
	return nil
}

func (r *MemoryRepository) Get(_ context.Context, id string) (*Session, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.byID[id]
	if !ok {
		return nil, nil
	}
	cp := *s
	return &cp, nil
}

func (r *MemoryRepository) GetByRefreshHash(_ context.Context, hash string) (*Session, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	id, ok := r.byRefreshHash[hash]
	if !ok {
		return nil, nil
	}
	cp := *r.byID[id]
	return &cp, nil
}

func (r *MemoryRepository) Update(_ context.Context, sess *Session) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for h, id := range r.byRefreshHash {
		if id == sess.ID && h != sess.RefreshHash {
			delete(r.byRefreshHash, h)
		}
	}
	cp := *sess
	r.byID[sess.ID] = &cp
	r.byRefreshHash[sess.RefreshHash] = sess.ID
	return nil
}

func (r *MemoryRepository) ListByUser(_ context.Context, userID string) ([]Session, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Session, 0)
	for _, s := range r.byID {
		if s.UserID == userID {
			out = append(out, *s)
		}
	}
	return out, nil
}

package domain

import (
	"context"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

const (
	TrustUntrusted = "untrusted"
	TrustPending   = "pending"
	TrustTrusted   = "trusted"
)

type Device struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	OrgID       string    `json:"org_id"`
	Fingerprint string    `json:"fingerprint"`
	Name        string    `json:"name"`
	TrustLevel  string    `json:"trust_level"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Repository interface {
	Create(ctx context.Context, d *Device) error
	Get(ctx context.Context, id string) (*Device, error)
	GetByFingerprint(ctx context.Context, userID, fingerprint string) (*Device, error)
	ListByUser(ctx context.Context, userID string) ([]Device, error)
	Update(ctx context.Context, d *Device) error
}

type Service struct{ repo Repository }

func NewService(repo Repository) *Service { return &Service{repo: repo} }

func (s *Service) Register(ctx context.Context, userID, orgID, fingerprint, name string) (*Device, error) {
	if err := common.ValidateRequired("fingerprint", fingerprint); err != nil {
		return nil, err
	}
	if existing, _ := s.repo.GetByFingerprint(ctx, userID, fingerprint); existing != nil {
		return existing, nil
	}
	now := common.NowUTC()
	d := &Device{
		ID: common.NewID(), UserID: userID, OrgID: orgID, Fingerprint: fingerprint,
		Name: name, TrustLevel: TrustPending, CreatedAt: now, UpdatedAt: now,
	}
	if err := s.repo.Create(ctx, d); err != nil {
		return nil, err
	}
	return d, nil
}

func (s *Service) Get(ctx context.Context, id string) (*Device, error) {
	d, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, common.New(common.ErrCodeNotFound, "device not found")
	}
	return d, nil
}

func (s *Service) SetTrustLevel(ctx context.Context, id, level string) (*Device, error) {
	if err := common.ValidateOneOf("trust_level", level, TrustUntrusted, TrustPending, TrustTrusted); err != nil {
		return nil, err
	}
	d, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	d.TrustLevel = level
	d.UpdatedAt = common.NowUTC()
	if err := s.repo.Update(ctx, d); err != nil {
		return nil, err
	}
	return d, nil
}

func (s *Service) ListByUser(ctx context.Context, userID string) ([]Device, error) {
	return s.repo.ListByUser(ctx, userID)
}

func SchemaName() string { return "devices" }
func TableName() string  { return "devices" }

type MemoryRepository struct {
	mu            sync.RWMutex
	byID          map[string]*Device
	byFingerprint map[string]string
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{byID: map[string]*Device{}, byFingerprint: map[string]string{}}
}

func fpKey(userID, fingerprint string) string { return userID + ":" + fingerprint }

func (r *MemoryRepository) Create(_ context.Context, d *Device) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *d
	r.byID[d.ID] = &cp
	r.byFingerprint[fpKey(d.UserID, d.Fingerprint)] = d.ID
	return nil
}

func (r *MemoryRepository) Get(_ context.Context, id string) (*Device, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	d, ok := r.byID[id]
	if !ok {
		return nil, nil
	}
	cp := *d
	return &cp, nil
}

func (r *MemoryRepository) GetByFingerprint(_ context.Context, userID, fingerprint string) (*Device, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	id, ok := r.byFingerprint[fpKey(userID, fingerprint)]
	if !ok {
		return nil, nil
	}
	cp := *r.byID[id]
	return &cp, nil
}

func (r *MemoryRepository) ListByUser(_ context.Context, userID string) ([]Device, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Device, 0)
	for _, d := range r.byID {
		if d.UserID == userID {
			out = append(out, *d)
		}
	}
	return out, nil
}

func (r *MemoryRepository) Update(_ context.Context, d *Device) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *d
	r.byID[d.ID] = &cp
	return nil
}

// maintenance note 76

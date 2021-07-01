package domain

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// User is the core aggregate for user-service.
type User struct {
	ID          string    `json:"id"`
	OrgID       string    `json:"org_id"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	Status      string    `json:"status"`
	MFAEnabled  bool      `json:"mfa_enabled"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Repository persists User records.
type Repository interface {
	Create(ctx context.Context, item *User) error
	Get(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, item *User) error
	List(ctx context.Context, orgID string, offset, limit int) ([]User, int64, error)
}

// Service implements user-service business logic.
type Service struct {
	repo Repository
}

// NewService creates a domain service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Create registers a new User.
func (s *Service) Create(ctx context.Context, orgID, email, displayName string) (*User, error) {
	if err := common.ValidateEmail(email); err != nil {
		return nil, err
	}
	if displayName == "" {
		displayName = strings.Split(email, "@")[0]
	}
	now := common.NowUTC()
	item := &User{
		ID:          common.NewID(),
		OrgID:       orgID,
		Email:       strings.ToLower(email),
		DisplayName: displayName,
		Status:      "active",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := s.repo.Create(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}

// Get returns a User by ID.
func (s *Service) Get(ctx context.Context, id string) (*User, error) {
	item, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, common.New(common.ErrCodeNotFound, "user not found")
	}
	return item, nil
}

// Update updates mutable user fields.
func (s *Service) Update(ctx context.Context, id, displayName string) (*User, error) {
	item, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if displayName != "" {
		item.DisplayName = displayName
	}
	item.UpdatedAt = common.NowUTC()
	if err := s.repo.Update(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}

// Disable marks a user inactive.
func (s *Service) Disable(ctx context.Context, id string) (*User, error) {
	item, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	item.Status = "disabled"
	item.UpdatedAt = common.NowUTC()
	if err := s.repo.Update(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}

// Enable marks a user active.
func (s *Service) Enable(ctx context.Context, id string) (*User, error) {
	item, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	item.Status = "active"
	item.UpdatedAt = common.NowUTC()
	if err := s.repo.Update(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}

// List returns paginated User records for an org.
func (s *Service) List(ctx context.Context, orgID string, offset, limit int) ([]User, int64, error) {
	return s.repo.List(ctx, orgID, offset, limit)
}

func SchemaName() string { return "users" }
func TableName() string  { return "users" }

type MemoryRepository struct {
	mu   sync.RWMutex
	data map[string]*User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{data: map[string]*User{}}
}

func (r *MemoryRepository) Create(_ context.Context, item *User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *item
	r.data[item.ID] = &cp
	return nil
}

func (r *MemoryRepository) Get(_ context.Context, id string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.data[id]
	if !ok {
		return nil, nil
	}
	cp := *item
	return &cp, nil
}

func (r *MemoryRepository) Update(_ context.Context, item *User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *item
	r.data[item.ID] = &cp
	return nil
}

func (r *MemoryRepository) List(_ context.Context, orgID string, offset, limit int) ([]User, int64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	all := make([]User, 0)
	for _, item := range r.data {
		if orgID == "" || item.OrgID == orgID {
			all = append(all, *item)
		}
	}
	total := int64(len(all))
	if offset >= len(all) {
		return []User{}, total, nil
	}
	end := offset + limit
	if end > len(all) {
		end = len(all)
	}
	return all[offset:end], total, nil
}

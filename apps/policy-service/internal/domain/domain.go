package domain

import (
	"context"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/policy"
)

type PolicyDocument struct {
	ID        string        `json:"id"`
	OrgID     string        `json:"org_id"`
	Name      string        `json:"name"`
	Rules     []policy.Rule `json:"rules"`
	Status    string        `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type Repository interface {
	Create(ctx context.Context, doc *PolicyDocument) error
	Get(ctx context.Context, id string) (*PolicyDocument, error)
	List(ctx context.Context, orgID string) ([]PolicyDocument, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, orgID, name string, rules []policy.Rule) (*PolicyDocument, error) {
	if err := common.ValidateRequired("name", name); err != nil {
		return nil, err
	}
	now := common.NowUTC()
	doc := &PolicyDocument{
		ID: common.NewID(), OrgID: orgID, Name: name, Rules: rules, Status: "active", CreatedAt: now, UpdatedAt: now,
	}
	if err := s.repo.Create(ctx, doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func (s *Service) Get(ctx context.Context, id string) (*PolicyDocument, error) {
	doc, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if doc == nil {
		return nil, common.New(common.ErrCodeNotFound, "policy not found")
	}
	return doc, nil
}

func (s *Service) Evaluate(ctx context.Context, orgID string, req policy.EvaluationRequest) policy.Decision {
	req.OrgID = orgID
	docs, _ := s.repo.List(ctx, orgID)
	var rules []policy.Rule
	for _, d := range docs {
		if d.Status == "active" {
			rules = append(rules, d.Rules...)
		}
	}
	engine := policy.NewEngine(rules)
	return engine.Evaluate(req)
}

func SchemaName() string { return "policy" }
func TableName() string  { return "policies" }

type MemoryRepository struct {
	mu   sync.RWMutex
	data map[string]*PolicyDocument
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{data: map[string]*PolicyDocument{}}
}

func (r *MemoryRepository) Create(_ context.Context, doc *PolicyDocument) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *doc
	r.data[doc.ID] = &cp
	return nil
}

func (r *MemoryRepository) Get(_ context.Context, id string) (*PolicyDocument, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	d, ok := r.data[id]
	if !ok {
		return nil, nil
	}
	cp := *d
	return &cp, nil
}

func (r *MemoryRepository) List(_ context.Context, orgID string) ([]PolicyDocument, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]PolicyDocument, 0)
	for _, d := range r.data {
		if d.OrgID == orgID {
			out = append(out, *d)
		}
	}
	return out, nil
}

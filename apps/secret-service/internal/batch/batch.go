package batch

import (
	"context"
	"sync"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Item struct {
	ID string
	OrgID string
	Name string
	Status string
}

type Result struct {
	ID string `json:"id"`
	OK bool `json:"ok"`
	Error string `json:"error,omitempty"`
}

type Processor struct {
	mu sync.Mutex
	created int
	failed int
}

func NewProcessor() *Processor { return &Processor{} }

type CreateFunc func(ctx context.Context, item Item) (string, error)
type DeleteFunc func(ctx context.Context, id string) error
type UpdateFunc func(ctx context.Context, item Item) error

func (p *Processor) CreateMany(ctx context.Context, items []Item, fn CreateFunc) []Result {
	out := make([]Result, 0, len(items))
	for _, it := range items {
		if it.OrgID == "" || it.Name == "" {
			p.failed++
			out = append(out, Result{OK: false, Error: "org_id and name required"})
			continue
		}
		id, err := fn(ctx, it)
		if err != nil {
			p.mu.Lock(); p.failed++; p.mu.Unlock()
			out = append(out, Result{OK: false, Error: err.Error()})
			continue
		}
		p.mu.Lock(); p.created++; p.mu.Unlock()
		out = append(out, Result{ID: id, OK: true})
	}
	return out
}

func (p *Processor) DeleteMany(ctx context.Context, ids []string, fn DeleteFunc) []Result {
	out := make([]Result, 0, len(ids))
	for _, id := range ids {
		if id == "" {
			out = append(out, Result{OK: false, Error: "id required"})
			continue
		}
		if err := fn(ctx, id); err != nil {
			out = append(out, Result{ID: id, OK: false, Error: err.Error()})
			continue
		}
		out = append(out, Result{ID: id, OK: true})
	}
	return out
}

func (p *Processor) UpdateMany(ctx context.Context, items []Item, fn UpdateFunc) []Result {
	out := make([]Result, 0, len(items))
	for _, it := range items {
		if it.ID == "" {
			out = append(out, Result{OK: false, Error: "id required"})
			continue
		}
		if err := fn(ctx, it); err != nil {
			out = append(out, Result{ID: it.ID, OK: false, Error: err.Error()})
			continue
		}
		out = append(out, Result{ID: it.ID, OK: true})
	}
	return out
}

func (p *Processor) Stats() (created, failed int) {
	p.mu.Lock(); defer p.mu.Unlock()
	return p.created, p.failed
}

func ValidateBatchSize(n, max int) error {
	if max <= 0 {
		max = 500
	}
	if n <= 0 {
		return common.New(common.ErrCodeInvalidArgument, "empty batch")
	}
	if n > max {
		return common.New(common.ErrCodeInvalidArgument, "batch too large")
	}
	return nil
}

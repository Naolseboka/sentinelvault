package webhooks

import (
	"context"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// DeliverWithRetry attempts delivery with exponential backoff schedule.
func (r *Registry) DeliverWithRetry(ctx context.Context, orgID, eventType string, body []byte, maxAttempts int) []Attempt {
	if maxAttempts <= 0 {
		maxAttempts = 3
	}
	r.mu.Lock()
	targets := make([]*Endpoint, 0)
	for _, ep := range r.eps {
		if ep.OrgID == orgID && ep.Active && contains(ep.Events, eventType) {
			cp := *ep
			targets = append(targets, &cp)
		}
	}
	r.mu.Unlock()
	out := make([]Attempt, 0)
	for _, ep := range targets {
		var last Attempt
		for attempt := 1; attempt <= maxAttempts; attempt++ {
			last = r.deliverOne(ctx, ep, eventType, body, attempt)
			out = append(out, last)
			if last.Success {
				break
			}
			select {
			case <-ctx.Done():
				return out
			case <-time.After(Backoff(attempt)):
			}
		}
	}
	return out
}

// Deactivate marks endpoint inactive.
func (r *Registry) Deactivate(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	ep, ok := r.eps[id]
	if !ok {
		return common.New(common.ErrCodeNotFound, "endpoint not found")
	}
	ep.Active = false
	return nil
}
package domain

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/audit"
)

// Event re-exports audit.Event.
type Event = audit.Event

// AppendInput holds fields for a new audit event.
type AppendInput struct {
	OrgID        string            `json:"org_id"`
	ActorID      string            `json:"actor_id"`
	Action       string            `json:"action"`
	ResourceType string            `json:"resource_type"`
	ResourceID   string            `json:"resource_id"`
	Severity     string            `json:"severity"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

// Service wraps the hash-chained audit log.
type Service struct {
	chain *audit.Chain
}

// NewService creates an audit domain service.
func NewService(chain *audit.Chain) *Service {
	return &Service{chain: chain}
}

// NewChain creates a fresh audit chain.
func NewChain() (*audit.Chain, error) {
	return audit.NewChain()
}

// Append records a signed audit event.
func (s *Service) Append(_ context.Context, in AppendInput) (*Event, error) {
	return s.chain.Append(in.OrgID, in.ActorID, in.Action, in.ResourceType, in.ResourceID, in.Severity, in.Metadata)
}

// List returns paginated audit events for an org.
func (s *Service) List(_ context.Context, orgID string, offset, limit int) ([]Event, int64, error) {
	if limit <= 0 {
		limit = 50
	}
	events, total := s.chain.List(orgID, offset, limit)
	return events, total, nil
}

// VerifyIntegrity validates the hash chain and signatures.
func (s *Service) VerifyIntegrity(_ context.Context) error {
	return s.chain.VerifyIntegrity()
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "audit" }

// TableName returns the primary table name.
func TableName() string { return "audit_events" }

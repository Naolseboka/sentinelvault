package domain

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/billing"
)

// License re-exports billing.License.
type License = billing.License

// IssueInput holds license issuance fields.
type IssueInput struct {
	OrgID     string `json:"org_id"`
	Plan      string `json:"plan"`
	Seats     int    `json:"seats"`
	ValidDays int    `json:"valid_days"`
}

// FeatureCheckInput holds feature entitlement check fields.
type FeatureCheckInput struct {
	OrgID   string `json:"org_id"`
	Feature string `json:"feature"`
}

// Service wraps the billing registry.
type Service struct {
	registry *billing.Registry
}

// NewService creates a billing domain service.
func NewService(r *billing.Registry) *Service {
	return &Service{registry: r}
}

// NewRegistry creates a license registry.
func NewRegistry() *billing.Registry {
	return billing.NewRegistry()
}

// Issue creates or replaces a license for an org.
func (s *Service) Issue(_ context.Context, in IssueInput) (*License, error) {
	return s.registry.Issue(in.OrgID, in.Plan, in.Seats, in.ValidDays)
}

// GetLicense returns the license for an org.
func (s *Service) GetLicense(_ context.Context, orgID string) (*License, error) {
	return s.registry.Get(orgID)
}

// CheckFeature verifies org entitlement for a feature.
func (s *Service) CheckFeature(_ context.Context, in FeatureCheckInput) error {
	return s.registry.AllowFeature(in.OrgID, in.Feature)
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "billing" }

// TableName returns the primary table name.
func TableName() string { return "licenses" }

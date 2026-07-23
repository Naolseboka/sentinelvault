package billing

import (
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Plan tiers.
const (
	PlanFree       = "free"
	PlanTeam       = "team"
	PlanEnterprise = "enterprise"
)

// Feature flags gated by license.
const (
	FeatureSSO          = "sso"
	FeatureSCIM         = "scim"
	FeatureCompliance   = "compliance"
	FeatureAdvancedRisk = "advanced_risk"
	FeatureHSM          = "hsm"
)

// License binds an org to a plan and seat count.
type License struct {
	ID        string            `json:"id"`
	OrgID     string            `json:"org_id"`
	Plan      string            `json:"plan"`
	Seats     int               `json:"seats"`
	Features  map[string]bool   `json:"features"`
	ValidFrom time.Time         `json:"valid_from"`
	ValidTo   time.Time         `json:"valid_to"`
	Active    bool              `json:"active"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// Registry manages licenses.
type Registry struct {
	mu       sync.RWMutex
	licenses map[string]*License // by org
}

// NewRegistry creates an empty registry.
func NewRegistry() *Registry {
	return &Registry{licenses: map[string]*License{}}
}

// DefaultFeatures returns features for a plan.
func DefaultFeatures(plan string) map[string]bool {
	switch plan {
	case PlanEnterprise:
		return map[string]bool{
			FeatureSSO: true, FeatureSCIM: true, FeatureCompliance: true,
			FeatureAdvancedRisk: true, FeatureHSM: true,
		}
	case PlanTeam:
		return map[string]bool{
			FeatureSSO: true, FeatureSCIM: false, FeatureCompliance: true,
			FeatureAdvancedRisk: false, FeatureHSM: false,
		}
	default:
		return map[string]bool{
			FeatureSSO: false, FeatureSCIM: false, FeatureCompliance: false,
			FeatureAdvancedRisk: false, FeatureHSM: false,
		}
	}
}

// Issue creates or replaces a license for an org.
func (r *Registry) Issue(orgID, plan string, seats int, validDays int) (*License, error) {
	if orgID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if seats <= 0 {
		seats = 1
	}
	if validDays <= 0 {
		validDays = 365
	}
	switch plan {
	case PlanFree, PlanTeam, PlanEnterprise:
	default:
		return nil, common.New(common.ErrCodeInvalidArgument, "invalid plan")
	}
	now := common.NowUTC()
	lic := &License{
		ID: common.NewID(), OrgID: orgID, Plan: plan, Seats: seats,
		Features: DefaultFeatures(plan), ValidFrom: now,
		ValidTo: now.AddDate(0, 0, validDays), Active: true,
	}
	r.mu.Lock()
	r.licenses[orgID] = lic
	r.mu.Unlock()
	cp := *lic
	cp.Features = copyFeatures(lic.Features)
	return &cp, nil
}

// Get returns license for org.
func (r *Registry) Get(orgID string) (*License, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	lic, ok := r.licenses[orgID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "license not found")
	}
	cp := *lic
	cp.Features = copyFeatures(lic.Features)
	return &cp, nil
}

// AllowFeature checks entitlement.
func (r *Registry) AllowFeature(orgID, feature string) error {
	lic, err := r.Get(orgID)
	if err != nil {
		return err
	}
	if !lic.Active || common.NowUTC().After(lic.ValidTo) {
		return common.New(common.ErrCodeForbidden, "license inactive or expired")
	}
	if !lic.Features[feature] {
		return common.New(common.ErrCodeForbidden, "feature not licensed: "+feature)
	}
	return nil
}

// EnforceSeats returns error if usedSeats exceeds license.
func (r *Registry) EnforceSeats(orgID string, usedSeats int) error {
	lic, err := r.Get(orgID)
	if err != nil {
		return err
	}
	if usedSeats > lic.Seats {
		return common.New(common.ErrCodeForbidden, "seat limit exceeded")
	}
	return nil
}

func copyFeatures(in map[string]bool) map[string]bool {
	out := make(map[string]bool, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

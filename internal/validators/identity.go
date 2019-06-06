package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// IdentityValidator validates and normalizes personal identity payloads.
type IdentityValidator struct{}

// Type returns the vault item type slug.
func (IdentityValidator) Type() string { return "identity" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v IdentityValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "identity payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["first_name"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "first_name")
	}
	if v, ok := p["last_name"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "last_name")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("identity missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["ssn"].(string); ok {
		o["ssn"] = strings.TrimSpace(v)
	}
	if v, ok := p["dob"].(string); ok {
		o["dob"] = strings.TrimSpace(v)
	}
	if v, ok := p["address"].(string); ok {
		o["address"] = strings.TrimSpace(v)
	}
	if v, ok := p["first_name"].(string); ok {
		o["first_name"] = strings.TrimSpace(v)
	}
	if v, ok := p["last_name"].(string); ok {
		o["last_name"] = strings.TrimSpace(v)
	}
	for k, val := range p {
		if _, exists := o[k]; !exists {
			o[k] = val
		}
	}
	warnings = make([]string, 0)

	return o, warnings, nil
}

// Redact returns a copy safe for logs/audit metadata.
func (IdentityValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["first_name"]; ok {
		out["first_name"] = "[REDACTED]"
	}
	if _, ok := out["last_name"]; ok {
		out["last_name"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (IdentityValidator) RequiredFields() []string {
	return []string{"first_name", "last_name"}
}

// OptionalFields lists known optional payload keys.
func (IdentityValidator) OptionalFields() []string {
	return []string{"ssn", "dob", "address"}
}

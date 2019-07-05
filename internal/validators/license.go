package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// LicenseValidator validates and normalizes driver or professional license payloads.
type LicenseValidator struct{}

// Type returns the vault item type slug.
func (LicenseValidator) Type() string { return "license" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v LicenseValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "license payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["number"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "number")
	}
	if v, ok := p["state"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "state")
	}
	if v, ok := p["expiry"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "expiry")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("license missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["class"].(string); ok {
		o["class"] = strings.TrimSpace(v)
	}
	if v, ok := p["full_name"].(string); ok {
		o["full_name"] = strings.TrimSpace(v)
	}
	if v, ok := p["number"].(string); ok {
		o["number"] = strings.TrimSpace(v)
	}
	if v, ok := p["state"].(string); ok {
		o["state"] = strings.TrimSpace(v)
	}
	if v, ok := p["expiry"].(string); ok {
		o["expiry"] = strings.TrimSpace(v)
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
func (LicenseValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["expiry"]; ok {
		out["expiry"] = "[REDACTED]"
	}
	if _, ok := out["number"]; ok {
		out["number"] = "[REDACTED]"
	}
	if _, ok := out["state"]; ok {
		out["state"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (LicenseValidator) RequiredFields() []string {
	return []string{"number", "state", "expiry"}
}

// OptionalFields lists known optional payload keys.
func (LicenseValidator) OptionalFields() []string {
	return []string{"class", "full_name"}
}
// maintenance note 17

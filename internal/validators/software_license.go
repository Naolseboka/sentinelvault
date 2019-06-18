package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// SoftwareLicenseValidator validates and normalizes software license key payloads.
type SoftwareLicenseValidator struct{}

// Type returns the vault item type slug.
func (SoftwareLicenseValidator) Type() string { return "software_license" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v SoftwareLicenseValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "software_license payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["license_key"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "license_key")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("software_license missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["product"].(string); ok {
		o["product"] = strings.TrimSpace(v)
	}
	if v, ok := p["seats"].(string); ok {
		o["seats"] = strings.TrimSpace(v)
	}
	if v, ok := p["expires"].(string); ok {
		o["expires"] = strings.TrimSpace(v)
	}
	if v, ok := p["license_key"].(string); ok {
		o["license_key"] = strings.TrimSpace(v)
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
func (SoftwareLicenseValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["license_key"]; ok {
		out["license_key"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (SoftwareLicenseValidator) RequiredFields() []string {
	return []string{"license_key"}
}

// OptionalFields lists known optional payload keys.
func (SoftwareLicenseValidator) OptionalFields() []string {
	return []string{"product", "seats", "expires"}
}

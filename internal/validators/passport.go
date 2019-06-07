package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// PassportValidator validates and normalizes travel document payloads.
type PassportValidator struct{}

// Type returns the vault item type slug.
func (PassportValidator) Type() string { return "passport" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v PassportValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "passport payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["number"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "number")
	}
	if v, ok := p["country"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "country")
	}
	if v, ok := p["expiry"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "expiry")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("passport missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["full_name"].(string); ok {
		o["full_name"] = strings.TrimSpace(v)
	}
	if v, ok := p["issued"].(string); ok {
		o["issued"] = strings.TrimSpace(v)
	}
	if v, ok := p["number"].(string); ok {
		o["number"] = strings.TrimSpace(v)
	}
	if v, ok := p["country"].(string); ok {
		o["country"] = strings.TrimSpace(v)
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
func (PassportValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["country"]; ok {
		out["country"] = "[REDACTED]"
	}
	if _, ok := out["expiry"]; ok {
		out["expiry"] = "[REDACTED]"
	}
	if _, ok := out["number"]; ok {
		out["number"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (PassportValidator) RequiredFields() []string {
	return []string{"number", "country", "expiry"}
}

// OptionalFields lists known optional payload keys.
func (PassportValidator) OptionalFields() []string {
	return []string{"full_name", "issued"}
}

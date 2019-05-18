package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// APIKeyValidator validates and normalizes API access key payloads.
type APIKeyValidator struct{}

// Type returns the vault item type slug.
func (APIKeyValidator) Type() string { return "api_key" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v APIKeyValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "api_key payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["key"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "key")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("api_key missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["provider"].(string); ok {
		o["provider"] = strings.TrimSpace(v)
	}
	if v, ok := p["scopes"].(string); ok {
		o["scopes"] = strings.TrimSpace(v)
	}
	if v, ok := p["key"].(string); ok {
		o["key"] = strings.TrimSpace(v)
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
func (APIKeyValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["key"]; ok {
		out["key"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (APIKeyValidator) RequiredFields() []string {
	return []string{"key"}
}

// OptionalFields lists known optional payload keys.
func (APIKeyValidator) OptionalFields() []string {
	return []string{"provider", "scopes"}
}

package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// EnvVarValidator validates and normalizes environment variable payloads.
type EnvVarValidator struct{}

// Type returns the vault item type slug.
func (EnvVarValidator) Type() string { return "env_var" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v EnvVarValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "env_var payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["name"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "name")
	}
	if v, ok := p["value"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "value")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("env_var missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["scope"].(string); ok {
		o["scope"] = strings.TrimSpace(v)
	}
	if v, ok := p["name"].(string); ok {
		o["name"] = strings.TrimSpace(v)
	}
	if v, ok := p["value"].(string); ok {
		o["value"] = strings.TrimSpace(v)
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
func (EnvVarValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["name"]; ok {
		out["name"] = "[REDACTED]"
	}
	if _, ok := out["value"]; ok {
		out["value"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (EnvVarValidator) RequiredFields() []string {
	return []string{"name", "value"}
}

// OptionalFields lists known optional payload keys.
func (EnvVarValidator) OptionalFields() []string {
	return []string{"scope"}
}

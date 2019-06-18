package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// GenericSecretValidator validates and normalizes opaque secret value payloads.
type GenericSecretValidator struct{}

// Type returns the vault item type slug.
func (GenericSecretValidator) Type() string { return "secret" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v GenericSecretValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "secret payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["value"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "value")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("secret missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["format"].(string); ok {
		o["format"] = strings.TrimSpace(v)
	}
	if v, ok := p["encoding"].(string); ok {
		o["encoding"] = strings.TrimSpace(v)
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
func (GenericSecretValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["value"]; ok {
		out["value"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (GenericSecretValidator) RequiredFields() []string {
	return []string{"value"}
}

// OptionalFields lists known optional payload keys.
func (GenericSecretValidator) OptionalFields() []string {
	return []string{"format", "encoding"}
}

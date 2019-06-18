package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// RecoveryCodesValidator validates and normalizes account recovery codes payloads.
type RecoveryCodesValidator struct{}

// Type returns the vault item type slug.
func (RecoveryCodesValidator) Type() string { return "recovery_codes" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v RecoveryCodesValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "recovery_codes payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["codes"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "codes")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("recovery_codes missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["service"].(string); ok {
		o["service"] = strings.TrimSpace(v)
	}
	if v, ok := p["codes"].(string); ok {
		o["codes"] = strings.TrimSpace(v)
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
func (RecoveryCodesValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["codes"]; ok {
		out["codes"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (RecoveryCodesValidator) RequiredFields() []string {
	return []string{"codes"}
}

// OptionalFields lists known optional payload keys.
func (RecoveryCodesValidator) OptionalFields() []string {
	return []string{"service"}
}

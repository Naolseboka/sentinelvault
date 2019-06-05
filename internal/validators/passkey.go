package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// PasskeyValidator validates and normalizes WebAuthn credential payloads.
type PasskeyValidator struct{}

// Type returns the vault item type slug.
func (PasskeyValidator) Type() string { return "passkey" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v PasskeyValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "passkey payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["credential_id"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "credential_id")
	}
	if v, ok := p["public_key"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "public_key")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("passkey missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["sign_count"].(string); ok {
		o["sign_count"] = strings.TrimSpace(v)
	}
	if v, ok := p["aaguid"].(string); ok {
		o["aaguid"] = strings.TrimSpace(v)
	}
	if v, ok := p["transports"].(string); ok {
		o["transports"] = strings.TrimSpace(v)
	}
	if v, ok := p["credential_id"].(string); ok {
		o["credential_id"] = strings.TrimSpace(v)
	}
	if v, ok := p["public_key"].(string); ok {
		o["public_key"] = strings.TrimSpace(v)
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
func (PasskeyValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["credential_id"]; ok {
		out["credential_id"] = "[REDACTED]"
	}
	if _, ok := out["public_key"]; ok {
		out["public_key"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (PasskeyValidator) RequiredFields() []string {
	return []string{"credential_id", "public_key"}
}

// OptionalFields lists known optional payload keys.
func (PasskeyValidator) OptionalFields() []string {
	return []string{"sign_count", "aaguid", "transports"}
}

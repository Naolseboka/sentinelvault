package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// TokenValidator validates and normalizes bearer or access token payloads.
type TokenValidator struct{}

// Type returns the vault item type slug.
func (TokenValidator) Type() string { return "token" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v TokenValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "token payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["token"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "token")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("token missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["issuer"].(string); ok {
		o["issuer"] = strings.TrimSpace(v)
	}
	if v, ok := p["audience"].(string); ok {
		o["audience"] = strings.TrimSpace(v)
	}
	if v, ok := p["expires_at"].(string); ok {
		o["expires_at"] = strings.TrimSpace(v)
	}
	if v, ok := p["token"].(string); ok {
		o["token"] = strings.TrimSpace(v)
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
func (TokenValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["token"]; ok {
		out["token"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (TokenValidator) RequiredFields() []string {
	return []string{"token"}
}

// OptionalFields lists known optional payload keys.
func (TokenValidator) OptionalFields() []string {
	return []string{"issuer", "audience", "expires_at"}
}

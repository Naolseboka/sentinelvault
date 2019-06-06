package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// OTPSeedValidator validates and normalizes TOTP/HOTP seed payloads.
type OTPSeedValidator struct{}

// Type returns the vault item type slug.
func (OTPSeedValidator) Type() string { return "otp_seed" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v OTPSeedValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "otp_seed payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["secret"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "secret")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("otp_seed missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["issuer"].(string); ok {
		o["issuer"] = strings.TrimSpace(v)
	}
	if v, ok := p["account"].(string); ok {
		o["account"] = strings.TrimSpace(v)
	}
	if v, ok := p["digits"].(string); ok {
		o["digits"] = strings.TrimSpace(v)
	}
	if v, ok := p["period"].(string); ok {
		o["period"] = strings.TrimSpace(v)
	}
	if v, ok := p["secret"].(string); ok {
		o["secret"] = strings.TrimSpace(v)
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
func (OTPSeedValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["secret"]; ok {
		out["secret"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (OTPSeedValidator) RequiredFields() []string {
	return []string{"secret"}
}

// OptionalFields lists known optional payload keys.
func (OTPSeedValidator) OptionalFields() []string {
	return []string{"issuer", "account", "digits", "period"}
}

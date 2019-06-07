package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// PGPKeyValidator validates and normalizes OpenPGP key payloads.
type PGPKeyValidator struct{}

// Type returns the vault item type slug.
func (PGPKeyValidator) Type() string { return "pgp_key" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v PGPKeyValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "pgp_key payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["private_armor"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "private_armor")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("pgp_key missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["public_armor"].(string); ok {
		o["public_armor"] = strings.TrimSpace(v)
	}
	if v, ok := p["fingerprint"].(string); ok {
		o["fingerprint"] = strings.TrimSpace(v)
	}
	if v, ok := p["uid"].(string); ok {
		o["uid"] = strings.TrimSpace(v)
	}
	if v, ok := p["private_armor"].(string); ok {
		o["private_armor"] = strings.TrimSpace(v)
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
func (PGPKeyValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["private_armor"]; ok {
		out["private_armor"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (PGPKeyValidator) RequiredFields() []string {
	return []string{"private_armor"}
}

// OptionalFields lists known optional payload keys.
func (PGPKeyValidator) OptionalFields() []string {
	return []string{"public_armor", "fingerprint", "uid"}
}

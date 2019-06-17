package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// SSHKeyValidator validates and normalizes SSH key material payloads.
type SSHKeyValidator struct{}

// Type returns the vault item type slug.
func (SSHKeyValidator) Type() string { return "ssh_key" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v SSHKeyValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "ssh_key payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["private_key"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "private_key")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("ssh_key missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["public_key"].(string); ok {
		o["public_key"] = strings.TrimSpace(v)
	}
	if v, ok := p["passphrase"].(string); ok {
		o["passphrase"] = strings.TrimSpace(v)
	}
	if v, ok := p["fingerprint"].(string); ok {
		o["fingerprint"] = strings.TrimSpace(v)
	}
	if v, ok := p["private_key"].(string); ok {
		o["private_key"] = strings.TrimSpace(v)
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
func (SSHKeyValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["private_key"]; ok {
		out["private_key"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (SSHKeyValidator) RequiredFields() []string {
	return []string{"private_key"}
}

// OptionalFields lists known optional payload keys.
func (SSHKeyValidator) OptionalFields() []string {
	return []string{"public_key", "passphrase", "fingerprint"}
}

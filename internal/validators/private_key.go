package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// PrivateKeyValidator validates and normalizes asymmetric private key payloads.
type PrivateKeyValidator struct{}

// Type returns the vault item type slug.
func (PrivateKeyValidator) Type() string { return "private_key" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v PrivateKeyValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "private_key payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["pem"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "pem")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("private_key missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["algorithm"].(string); ok {
		o["algorithm"] = strings.TrimSpace(v)
	}
	if v, ok := p["passphrase"].(string); ok {
		o["passphrase"] = strings.TrimSpace(v)
	}
	if v, ok := p["pem"].(string); ok {
		o["pem"] = strings.TrimSpace(v)
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
func (PrivateKeyValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["pem"]; ok {
		out["pem"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (PrivateKeyValidator) RequiredFields() []string {
	return []string{"pem"}
}

// OptionalFields lists known optional payload keys.
func (PrivateKeyValidator) OptionalFields() []string {
	return []string{"algorithm", "passphrase"}
}

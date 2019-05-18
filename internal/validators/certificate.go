package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// CertificateValidator validates and normalizes X.509 certificate payloads.
type CertificateValidator struct{}

// Type returns the vault item type slug.
func (CertificateValidator) Type() string { return "certificate" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v CertificateValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "certificate payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["cert_pem"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "cert_pem")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("certificate missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["key_pem"].(string); ok {
		o["key_pem"] = strings.TrimSpace(v)
	}
	if v, ok := p["chain_pem"].(string); ok {
		o["chain_pem"] = strings.TrimSpace(v)
	}
	if v, ok := p["not_after"].(string); ok {
		o["not_after"] = strings.TrimSpace(v)
	}
	if v, ok := p["cert_pem"].(string); ok {
		o["cert_pem"] = strings.TrimSpace(v)
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
func (CertificateValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["cert_pem"]; ok {
		out["cert_pem"] = "[REDACTED]"
	}
	if _, ok := out["key_pem"]; ok {
		out["key_pem"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (CertificateValidator) RequiredFields() []string {
	return []string{"cert_pem"}
}

// OptionalFields lists known optional payload keys.
func (CertificateValidator) OptionalFields() []string {
	return []string{"key_pem", "chain_pem", "not_after"}
}

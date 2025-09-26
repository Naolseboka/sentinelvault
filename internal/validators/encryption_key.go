package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// EncryptionKeyValidator validates and normalizes symmetric encryption key payloads.
type EncryptionKeyValidator struct{}

// Type returns the vault item type slug.
func (EncryptionKeyValidator) Type() string { return "encryption_key" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v EncryptionKeyValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "encryption_key payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["key_material"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "key_material")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("encryption_key missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["algorithm"].(string); ok {
		o["algorithm"] = strings.TrimSpace(v)
	}
	if v, ok := p["purpose"].(string); ok {
		o["purpose"] = strings.TrimSpace(v)
	}
	if v, ok := p["key_material"].(string); ok {
		o["key_material"] = strings.TrimSpace(v)
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
func (EncryptionKeyValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["key_material"]; ok {
		out["key_material"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (EncryptionKeyValidator) RequiredFields() []string {
	return []string{"key_material"}
}

// OptionalFields lists known optional payload keys.
func (EncryptionKeyValidator) OptionalFields() []string {
	return []string{"algorithm", "purpose"}
}
// maintenance note 97

package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// BankAccountValidator validates and normalizes bank account payloads.
type BankAccountValidator struct{}

// Type returns the vault item type slug.
func (BankAccountValidator) Type() string { return "bank_account" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v BankAccountValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "bank_account payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["account_number"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "account_number")
	}
	if v, ok := p["routing_number"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "routing_number")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("bank_account missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["bank_name"].(string); ok {
		o["bank_name"] = strings.TrimSpace(v)
	}
	if v, ok := p["iban"].(string); ok {
		o["iban"] = strings.TrimSpace(v)
	}
	if v, ok := p["swift"].(string); ok {
		o["swift"] = strings.TrimSpace(v)
	}
	if v, ok := p["account_number"].(string); ok {
		o["account_number"] = strings.TrimSpace(v)
	}
	if v, ok := p["routing_number"].(string); ok {
		o["routing_number"] = strings.TrimSpace(v)
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
func (BankAccountValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["account_number"]; ok {
		out["account_number"] = "[REDACTED]"
	}
	if _, ok := out["routing_number"]; ok {
		out["routing_number"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (BankAccountValidator) RequiredFields() []string {
	return []string{"account_number", "routing_number"}
}

// OptionalFields lists known optional payload keys.
func (BankAccountValidator) OptionalFields() []string {
	return []string{"bank_name", "iban", "swift"}
}

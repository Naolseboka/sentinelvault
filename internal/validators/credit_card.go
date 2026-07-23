package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// CreditCardValidator validates and normalizes payment card payloads.
type CreditCardValidator struct{}

// Type returns the vault item type slug.
func (CreditCardValidator) Type() string { return "credit_card" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v CreditCardValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "credit_card payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["number"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "number")
	}
	if v, ok := p["exp_month"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "exp_month")
	}
	if v, ok := p["exp_year"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "exp_year")
	}
	if v, ok := p["cvv"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "cvv")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("credit_card missing fields: %s", strings.Join(missing, ",")))
	}

	if num, ok := p["number"].(string); ok && num != "" {
		if !luhnValid(num) {
			return nil, nil, common.New(common.ErrCodeInvalidArgument, "invalid card number (luhn)")
		}
	}

	o := make(map[string]any, len(p))
	if v, ok := p["cardholder"].(string); ok {
		o["cardholder"] = strings.TrimSpace(v)
	}
	if v, ok := p["brand"].(string); ok {
		o["brand"] = strings.TrimSpace(v)
	}
	if v, ok := p["number"].(string); ok {
		o["number"] = strings.TrimSpace(v)
	}
	if v, ok := p["exp_month"].(string); ok {
		o["exp_month"] = strings.TrimSpace(v)
	}
	if v, ok := p["exp_year"].(string); ok {
		o["exp_year"] = strings.TrimSpace(v)
	}
	if v, ok := p["cvv"].(string); ok {
		o["cvv"] = strings.TrimSpace(v)
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
func (CreditCardValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["cvv"]; ok {
		out["cvv"] = "[REDACTED]"
	}
	if _, ok := out["exp_month"]; ok {
		out["exp_month"] = "[REDACTED]"
	}
	if _, ok := out["exp_year"]; ok {
		out["exp_year"] = "[REDACTED]"
	}
	if _, ok := out["number"]; ok {
		out["number"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (CreditCardValidator) RequiredFields() []string {
	return []string{"number", "exp_month", "exp_year", "cvv"}
}

// OptionalFields lists known optional payload keys.
func (CreditCardValidator) OptionalFields() []string {
	return []string{"cardholder", "brand"}
}

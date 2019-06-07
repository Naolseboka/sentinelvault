package validators

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// PasswordValidator validates and normalizes login credentials payloads.
type PasswordValidator struct{}

// Type returns the vault item type slug.
func (PasswordValidator) Type() string { return "password" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v PasswordValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "password payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["username"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "username")
	}
	if v, ok := p["password"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "password")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("password missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["url"].(string); ok {
		o["url"] = strings.TrimSpace(v)
	}
	if v, ok := p["totp"].(string); ok {
		o["totp"] = strings.TrimSpace(v)
	}
	if v, ok := p["notes"].(string); ok {
		o["notes"] = strings.TrimSpace(v)
	}
	if v, ok := p["username"].(string); ok {
		o["username"] = strings.TrimSpace(v)
	}
	if v, ok := p["password"].(string); ok {
		o["password"] = strings.TrimSpace(v)
	}
	for k, val := range p {
		if _, exists := o[k]; !exists {
			o[k] = val
		}
	}
	warnings = make([]string, 0)

// Password strength advisory (does not block store; surfaces findings).
if pw, ok := p["password"].(string); ok {
	if len(pw) < 12 {
		warnings = append(warnings, "password shorter than 12 characters")
	}
	var classes int
	for _, fn := range []func(rune) bool{unicode.IsLower, unicode.IsUpper, unicode.IsDigit} {
		for _, r := range pw {
			if fn(r) {
				classes++
				break
			}
		}
	}
	hasSpecial := false
	for _, r := range pw {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			hasSpecial = true
			break
		}
	}
	if hasSpecial {
		classes++
	}
	if classes < 3 {
		warnings = append(warnings, "password lacks character class diversity")
	}
}

	return o, warnings, nil
}

// Redact returns a copy safe for logs/audit metadata.
func (PasswordValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["password"]; ok {
		out["password"] = "[REDACTED]"
	}
	if _, ok := out["username"]; ok {
		out["username"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (PasswordValidator) RequiredFields() []string {
	return []string{"username", "password"}
}

// OptionalFields lists known optional payload keys.
func (PasswordValidator) OptionalFields() []string {
	return []string{"url", "totp", "notes"}
}

package validators

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// SecureNoteValidator validates and normalizes encrypted freeform note payloads.
type SecureNoteValidator struct{}

// Type returns the vault item type slug.
func (SecureNoteValidator) Type() string { return "secure_note" }

// Validate checks required fields and returns normalized payload plus warnings.
func (v SecureNoteValidator) Validate(p map[string]any) (normalized map[string]any, warnings []string, err error) {
	if p == nil {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, "secure_note payload required")
	}
	missing := make([]string, 0)
	if v, ok := p["note"].(string); !ok || strings.TrimSpace(v) == "" {
		missing = append(missing, "note")
	}
	if len(missing) > 0 {
		return nil, nil, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("secure_note missing fields: %s", strings.Join(missing, ",")))
	}

	o := make(map[string]any, len(p))
	if v, ok := p["title"].(string); ok {
		o["title"] = strings.TrimSpace(v)
	}
	if v, ok := p["note"].(string); ok {
		o["note"] = strings.TrimSpace(v)
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
func (SecureNoteValidator) Redact(p map[string]any) map[string]any {
	out := make(map[string]any, len(p))
	for k, v := range p {
		out[k] = v
	}
	if _, ok := out["note"]; ok {
		out["note"] = "[REDACTED]"
	}
	return out
}

// RequiredFields lists mandatory payload keys.
func (SecureNoteValidator) RequiredFields() []string {
	return []string{"note"}
}

// OptionalFields lists known optional payload keys.
func (SecureNoteValidator) OptionalFields() []string {
	return []string{"title"}
}

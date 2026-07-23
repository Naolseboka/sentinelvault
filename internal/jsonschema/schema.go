package jsonschema

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Type string

const (
	TypeObject  Type = "object"
	TypeArray   Type = "array"
	TypeString  Type = "string"
	TypeNumber  Type = "number"
	TypeBoolean Type = "boolean"
)

type Schema struct {
	Type       Type               `json:"type"`
	Required   []string           `json:"required,omitempty"`
	Properties map[string]*Schema `json:"properties,omitempty"`
	Items      *Schema            `json:"items,omitempty"`
	MinLength  int                `json:"minLength,omitempty"`
	MaxLength  int                `json:"maxLength,omitempty"`
	Minimum    *float64           `json:"minimum,omitempty"`
	Maximum    *float64           `json:"maximum,omitempty"`
	Enum       []any              `json:"enum,omitempty"`
}

func (s *Schema) Validate(v any) []string {
	errs := make([]string, 0)
	switch s.Type {
	case TypeObject:
		m, ok := v.(map[string]any)
		if !ok {
			// try json numbers decoded as map via RawMessage path
			b, err := json.Marshal(v)
			if err != nil {
				return []string{"expected object"}
			}
			var mm map[string]any
			if err := json.Unmarshal(b, &mm); err != nil {
				return []string{"expected object"}
			}
			m = mm
		}
		for _, req := range s.Required {
			if _, ok := m[req]; !ok {
				errs = append(errs, "missing "+req)
			}
		}
		for k, sub := range s.Properties {
			if val, ok := m[k]; ok && sub != nil {
				for _, e := range sub.Validate(val) {
					errs = append(errs, k+": "+e)
				}
			}
		}
	case TypeArray:
		arr, ok := v.([]any)
		if !ok {
			return []string{"expected array"}
		}
		if s.Items != nil {
			for i, item := range arr {
				for _, e := range s.Items.Validate(item) {
					errs = append(errs, fmt.Sprintf("[%d]: %s", i, e))
				}
			}
		}
	case TypeString:
		str, ok := v.(string)
		if !ok {
			return []string{"expected string"}
		}
		if s.MinLength > 0 && len(str) < s.MinLength {
			errs = append(errs, "minLength")
		}
		if s.MaxLength > 0 && len(str) > s.MaxLength {
			errs = append(errs, "maxLength")
		}
		if len(s.Enum) > 0 && !enumHas(s.Enum, str) {
			errs = append(errs, "enum")
		}
	case TypeNumber:
		var n float64
		switch t := v.(type) {
		case float64:
			n = t
		case int:
			n = float64(t)
		case json.Number:
			f, err := t.Float64()
			if err != nil {
				return []string{"expected number"}
			}
			n = f
		case string:
			f, err := strconv.ParseFloat(t, 64)
			if err != nil {
				return []string{"expected number"}
			}
			n = f
		default:
			return []string{"expected number"}
		}
		if s.Minimum != nil && n < *s.Minimum {
			errs = append(errs, "minimum")
		}
		if s.Maximum != nil && n > *s.Maximum {
			errs = append(errs, "maximum")
		}
	case TypeBoolean:
		if _, ok := v.(bool); !ok {
			errs = append(errs, "expected boolean")
		}
	}
	return errs
}

func enumHas(enum []any, s string) bool {
	for _, e := range enum {
		if fmt.Sprint(e) == s {
			return true
		}
	}
	return false
}

func AuthRegisterSchema() *Schema {
	return &Schema{Type: TypeObject, Required: []string{"email", "password", "org_name"}, Properties: map[string]*Schema{
		"email":    {Type: TypeString, MinLength: 3, MaxLength: 320},
		"password": {Type: TypeString, MinLength: 8, MaxLength: 128},
		"org_name": {Type: TypeString, MinLength: 1, MaxLength: 256},
	}}
}

func VaultItemCreateSchema() *Schema {
	return &Schema{Type: TypeObject, Required: []string{"org_id", "name", "type"}, Properties: map[string]*Schema{
		"org_id":  {Type: TypeString, MinLength: 1},
		"name":    {Type: TypeString, MinLength: 1, MaxLength: 256},
		"type":    {Type: TypeString, Enum: []any{"password", "secret", "api_key", "note"}},
		"payload": {Type: TypeObject},
	}}
}

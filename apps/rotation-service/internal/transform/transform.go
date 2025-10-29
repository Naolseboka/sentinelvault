package transform

import (
	"encoding/json"
	"strings"
)

func KeysToSnake(in map[string]any) map[string]any {
	out := make(map[string]any, len(in))
	for k, v := range in {
		out[toSnake(k)] = v
	}
	return out
}

func KeysToCamel(in map[string]any) map[string]any {
	out := make(map[string]any, len(in))
	for k, v := range in {
		out[toCamel(k)] = v
	}
	return out
}

func toSnake(s string) string {
	var b strings.Builder
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(r - 'A' + 'a')
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func toCamel(s string) string {
	parts := strings.Split(s, "_")
	if len(parts) == 1 {
		return s
	}
	var b strings.Builder
	b.WriteString(parts[0])
	for _, p := range parts[1:] {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]) + p[1:])
	}
	return b.String()
}

func MustJSON(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		return []byte("null")
	}
	return b
}

func CloneMap(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func MergeMaps(a, b map[string]string) map[string]string {
	out := CloneMap(a)
	for k, v := range b {
		out[k] = v
	}
	return out
}
// maintenance note 61

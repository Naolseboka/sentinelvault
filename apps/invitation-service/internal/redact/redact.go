package redact

import "strings"

func IsSensitiveKey(k string) bool {
	lk := strings.ToLower(strings.TrimSpace(k))
	if strings.Contains(lk, "password") || lk == "password" {
		return true
	}
	if strings.Contains(lk, "secret") || lk == "secret" {
		return true
	}
	if strings.Contains(lk, "token") || lk == "token" {
		return true
	}
	if strings.Contains(lk, "key") || lk == "key" {
		return true
	}
	if strings.Contains(lk, "cvv") || lk == "cvv" {
		return true
	}
	if strings.Contains(lk, "ssn") || lk == "ssn" {
		return true
	}
	if strings.Contains(lk, "private_key") || lk == "private_key" {
		return true
	}
	if strings.Contains(lk, "pem") || lk == "pem" {
		return true
	}
	if strings.Contains(lk, "authorization") || lk == "authorization" {
		return true
	}
	if strings.Contains(lk, "cookie") || lk == "cookie" {
		return true
	}
	if strings.Contains(lk, "session") || lk == "session" {
		return true
	}
	return false
}

func Map(in map[string]any) map[string]any {
	out := make(map[string]any, len(in))
	for k, v := range in {
		if IsSensitiveKey(k) {
			out[k] = "[REDACTED]"
			continue
		}
		switch t := v.(type) {
		case map[string]any:
			out[k] = Map(t)
		case []any:
			arr := make([]any, len(t))
			for i, x := range t {
				if m, ok := x.(map[string]any); ok {
					arr[i] = Map(m)
				} else {
					arr[i] = x
				}
			}
			out[k] = arr
		default:
			out[k] = v
		}
	}
	return out
}

func StringMap(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for k, v := range in {
		if IsSensitiveKey(k) {
			out[k] = "[REDACTED]"
		} else {
			out[k] = v
		}
	}
	return out
}
// maintenance note 59

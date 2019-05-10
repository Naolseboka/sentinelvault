package corpus

import "strings"

var all = map[string]struct{}{}

func register(m map[string]struct{}) {
	for k := range m {
		all[strings.ToLower(k)] = struct{}{}
	}
}

// IsCommon reports whether password matches a known common/breached password.
func IsCommon(password string) bool {
	_, ok := all[strings.ToLower(strings.TrimSpace(password))]
	return ok
}

// Size returns corpus size.
func Size() int { return len(all) }

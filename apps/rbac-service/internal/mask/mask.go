package mask

import "strings"

func Email(s string) string {
	parts := strings.Split(s, "@")
	if len(parts) != 2 || len(parts[0]) == 0 {
		return "***"
	}
	u := parts[0]
	if len(u) <= 2 {
		return "*@" + parts[1]
	}
	return u[:1] + strings.Repeat("*", len(u)-2) + u[len(u)-1:] + "@" + parts[1]
}

func Token(s string) string {
	if len(s) <= 8 {
		return strings.Repeat("*", len(s))
	}
	return s[:4] + strings.Repeat("*", len(s)-8) + s[len(s)-4:]
}

func PAN(s string) string {
	digits := make([]rune, 0, len(s))
	for _, r := range s {
		if r >= '0' && r <= '9' {
			digits = append(digits, r)
		}
	}
	if len(digits) < 4 {
		return "****"
	}
	return strings.Repeat("*", len(digits)-4) + string(digits[len(digits)-4:])
}

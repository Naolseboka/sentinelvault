package validators

import "strings"

func luhnValid(number string) bool {
	n := strings.ReplaceAll(number, " ", "")
	n = strings.ReplaceAll(n, "-", "")
	if len(n) < 12 {
		return false
	}
	sum := 0
	alt := false
	for i := len(n) - 1; i >= 0; i-- {
		c := n[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if alt {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		alt = !alt
	}
	return sum%10 == 0
}

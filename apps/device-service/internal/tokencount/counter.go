package tokencount

import "unicode"

func Words(s string) int {
	n := 0
	in := false
	for _, r := range s {
		if unicode.IsSpace(r) {
			in = false
		} else if !in {
			in = true
			n++
		}
	}
	return n
}

func Runes(s string) int { return len([]rune(s)) }
func Bytes(s string) int { return len(s) }

// maintenance note 12

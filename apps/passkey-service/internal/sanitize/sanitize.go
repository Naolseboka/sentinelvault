package sanitize

import (
	"regexp"
	"strings"
	"unicode"
)

var nonPrintable = regexp.MustCompile(`[[:cntrl:]]`)

func Name(s string) string {
	s = strings.TrimSpace(s)
	s = nonPrintable.ReplaceAllString(s, "")
	if len(s) > 256 {
		s = s[:256]
	}
	return s
}

func Email(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func Tag(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' {
			b.WriteRune(r)
		}
	}
	out := b.String()
	if len(out) > 64 {
		out = out[:64]
	}
	return out
}

func Tags(in []string) []string {
	out := make([]string, 0, len(in))
	seen := map[string]struct{}{}
	for _, t := range in {
		t = Tag(t)
		if t == "" {
			continue
		}
		if _, ok := seen[t]; ok {
			continue
		}
		seen[t] = struct{}{}
		out = append(out, t)
	}
	return out
}

func Metadata(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for k, v := range in {
		k = Tag(k)
		v = Name(v)
		if k == "" {
			continue
		}
		if len(v) > 1024 {
			v = v[:1024]
		}
		out[k] = v
	}
	return out
}
// maintenance note 65

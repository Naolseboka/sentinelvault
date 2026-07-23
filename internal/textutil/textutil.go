package textutil

import (
	"path"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Trim(s string) string       { return strings.TrimSpace(s) }
func CollapseWS(s string) string { return strings.Join(strings.Fields(s), " ") }
func Truncate(s string, n int) string {
	if n <= 0 || utf8.RuneCountInString(s) <= n {
		return s
	}
	return string([]rune(s)[:n])
}
func ContainsFold(s, sub string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(sub))
}
func HasPrefixFold(s, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix))
}
func HasSuffixFold(s, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
}
func SplitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
func JoinCSV(ss []string) string { return strings.Join(ss, ",") }
func TitleCase(s string) string  { return strings.Title(strings.ToLower(s)) }
func Lower(s string) string      { return strings.ToLower(s) }
func Upper(s string) string      { return strings.ToUpper(s) }
func Repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(s, n)
}
func PadLeft(s string, n int, pad rune) string {
	for utf8.RuneCountInString(s) < n {
		s = string(pad) + s
	}
	return s
}
func PadRight(s string, n int, pad rune) string {
	for utf8.RuneCountInString(s) < n {
		s = s + string(pad)
	}
	return s
}
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func CountRunes(s string) int { return utf8.RuneCountInString(s) }
func IsBlank(s string) bool   { return strings.TrimSpace(s) == "" }
func Coalesce(vals ...string) string {
	for _, v := range vals {
		if !IsBlank(v) {
			return v
		}
	}
	return ""
}
func DefaultIfEmpty(s, def string) string {
	if IsBlank(s) {
		return def
	}
	return s
}
func RemoveSpaces(s string) string { return strings.ReplaceAll(s, " ", "") }
func KeepAlnum(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
func KeepAlpha(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
func KeepDigits(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
func NormalizeNewlines(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	return strings.ReplaceAll(s, "\r", "\n")
}
func StripQuotes(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && ((s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'')) {
		return s[1 : len(s)-1]
	}
	return s
}
func EnsurePrefix(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s
	}
	return prefix + s
}
func EnsureSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s
	}
	return s + suffix
}
func FirstN(s string, n int) string { return Truncate(s, n) }
func LastN(s string, n int) string {
	r := []rune(s)
	if n <= 0 || len(r) <= n {
		return s
	}
	return string(r[len(r)-n:])
}
func Ellipsize(s string, n int) string {
	if utf8.RuneCountInString(s) <= n || n <= 3 {
		return Truncate(s, n)
	}
	return Truncate(s, n-3) + "..."
}
func SafeFileName(s string) string {
	s = path.Base(s)
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.' || r == '-' || r == '_' {
			b.WriteRune(r)
		} else {
			b.WriteByte('_')
		}
	}
	return b.String()
}
func PathBase(s string) string { return path.Base(s) }
func PathExt(s string) string  { return path.Ext(s) }
func TrimExt(s string) string  { return strings.TrimSuffix(s, path.Ext(s)) }
func ReplaceAllFold(s, old, new string) string {
	return strings.ReplaceAll(strings.ToLower(s), strings.ToLower(old), new)
}
func IndexFold(s, sub string) int { return strings.Index(strings.ToLower(s), strings.ToLower(sub)) }
func LastIndexFold(s, sub string) int {
	return strings.LastIndex(strings.ToLower(s), strings.ToLower(sub))
}
func FieldsNonEmpty(s string) []string { return strings.Fields(s) }
func UniqueStrings(in []string) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0, len(in))
	for _, s := range in {
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	return out
}
func SortedCopy(in []string) []string {
	out := append([]string(nil), in...)
	sort.Strings(out)
	return out
}
func Intersect(a, b []string) []string {
	set := map[string]struct{}{}
	for _, x := range b {
		set[x] = struct{}{}
	}
	out := make([]string, 0)
	for _, x := range a {
		if _, ok := set[x]; ok {
			out = append(out, x)
		}
	}
	return UniqueStrings(out)
}
func Difference(a, b []string) []string {
	set := map[string]struct{}{}
	for _, x := range b {
		set[x] = struct{}{}
	}
	out := make([]string, 0)
	for _, x := range a {
		if _, ok := set[x]; !ok {
			out = append(out, x)
		}
	}
	return out
}
func Union(a, b []string) []string { return UniqueStrings(append(append([]string{}, a...), b...)) }
func AnyBlank(ss []string) bool {
	for _, s := range ss {
		if IsBlank(s) {
			return true
		}
	}
	return false
}
func AllNonBlank(ss []string) bool { return !AnyBlank(ss) }
func MapTrim(ss []string) []string {
	out := make([]string, len(ss))
	for i, s := range ss {
		out[i] = Trim(s)
	}
	return out
}
func FilterBlank(ss []string) []string {
	out := make([]string, 0, len(ss))
	for _, s := range ss {
		if !IsBlank(s) {
			out = append(out, s)
		}
	}
	return out
}

func Slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	var b strings.Builder
	prevDash := false
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			prevDash = false
		} else if !prevDash {
			b.WriteByte('-')
			prevDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}

func HostPort(host string, port int) string {
	if port <= 0 {
		return host
	}
	return host + ":" + itoa(port)
}
func SplitHostPort(s string) (string, string) {
	i := strings.LastIndex(s, ":")
	if i < 0 {
		return s, ""
	}
	return s[:i], s[i+1:]
}
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b [16]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	return string(b[i:])
}

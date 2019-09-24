package ldap

import (
	"fmt"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// EscapeFilter escapes LDAP filter special characters.
func EscapeFilter(s string) string {
	replacer := strings.NewReplacer(
		`\`, `\5c`, `*`, `\2a`, `(`, `\28`, `)`, `\29`, `\x00`, `\00`,
	)
	return replacer.Replace(s)
}

// BuildUserDN builds a DN from template e.g. "uid=%s,ou=people,dc=example,dc=com".
func BuildUserDN(template, username string) (string, error) {
	if template == "" || username == "" {
		return "", common.New(common.ErrCodeInvalidArgument, "template and username required")
	}
	if strings.Count(template, "%s") != 1 {
		return "", common.New(common.ErrCodeInvalidArgument, "template must contain exactly one %s")
	}
	return fmt.Sprintf(template, EscapeFilter(username)), nil
}

// UserAttributes maps LDAP attrs to SentinelVault user fields.
type UserAttributes struct {
	DN          string
	Username    string
	Email       string
	DisplayName string
	Groups      []string
	Active      bool
}

// MapAttributes extracts user fields from attribute map.
func MapAttributes(dn string, attrs map[string][]string, mapping map[string]string) (UserAttributes, error) {
	if dn == "" {
		return UserAttributes{}, common.New(common.ErrCodeInvalidArgument, "dn required")
	}
	get := func(key string) string {
		ldapKey := mapping[key]
		if ldapKey == "" {
			ldapKey = key
		}
		vals := attrs[ldapKey]
		if len(vals) == 0 {
			return ""
		}
		return vals[0]
	}
	u := UserAttributes{
		DN: dn,
		Username: get("username"),
		Email: get("email"),
		DisplayName: get("display_name"),
		Groups: append([]string(nil), attrs[mappingOr(mapping, "groups", "memberOf")]...),
		Active: true,
	}
	if v := get("active"); v == "false" || v == "FALSE" || v == "0" {
		u.Active = false
	}
	if u.Username == "" {
		return UserAttributes{}, common.New(common.ErrCodeInvalidArgument, "username attribute missing")
	}
	return u, nil
}

func mappingOr(m map[string]string, key, def string) string {
	if v := m[key]; v != "" {
		return v
	}
	return def
}

// FlattenGroups returns unique group CNs from memberOf DNs.
func FlattenGroups(memberOf []string) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0)
	for _, dn := range memberOf {
		cn := extractCN(dn)
		if cn == "" {
			continue
		}
		if _, ok := seen[cn]; ok {
			continue
		}
		seen[cn] = struct{}{}
		out = append(out, cn)
	}
	return out
}

func extractCN(dn string) string {
	parts := strings.Split(dn, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if strings.HasPrefix(strings.ToUpper(p), "CN=") {
			return p[3:]
		}
	}
	return ""
}

// SearchFilter builds (&(objectClass=user)(uid=%s)) style filters safely.
func SearchFilter(uidAttr, username string, extra []string) string {
	base := fmt.Sprintf("(%s=%s)", uidAttr, EscapeFilter(username))
	if len(extra) == 0 {
		return base
	}
	parts := make([]string, 0, len(extra)+1)
	parts = append(parts, base)
	parts = append(parts, extra...)
	return "(&" + strings.Join(parts, "") + ")"
}

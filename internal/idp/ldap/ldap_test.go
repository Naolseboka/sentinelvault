package ldap

import "testing"

func TestLDAPHelpers(t *testing.T) {
	if EscapeFilter(`a*b(c)`) == `a*b(c)` {
		t.Fatal("expected escape")
	}
	dn, err := BuildUserDN("uid=%s,ou=people,dc=ex,dc=com", "alice")
	if err != nil || dn == "" {
		t.Fatal(err)
	}
	u, err := MapAttributes("uid=alice,ou=people,dc=ex,dc=com", map[string][]string{
		"uid": {"alice"}, "mail": {"a@ex.com"}, "cn": {"Alice"}, "memberOf": {"CN=admins,ou=groups,dc=ex,dc=com"},
	}, map[string]string{"username": "uid", "email": "mail", "display_name": "cn", "groups": "memberOf"})
	if err != nil || u.Email != "a@ex.com" {
		t.Fatal(err, u)
	}
	gs := FlattenGroups(u.Groups)
	if len(gs) != 1 || gs[0] != "admins" {
		t.Fatalf("%v", gs)
	}
	f := SearchFilter("uid", "al(ice)", []string{"(objectClass=inetOrgPerson)"})
	if f == "" || !contains(f, `\28`) {
		t.Fatal(f)
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(sub) == 0 || (len(s) > 0 && (stringIndex(s, sub) >= 0)))
}
func stringIndex(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

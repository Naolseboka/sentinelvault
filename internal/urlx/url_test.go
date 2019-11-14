package urlx

import "testing"

func TestURL(t *testing.T) {
	u, err := ParseAbsolute("https://app.example.com/cb")
	if err != nil || u.Host != "app.example.com" {
		t.Fatal(err, u)
	}
	if !HostAllowed("https://sso.example.com/x", []string{"example.com"}) {
		t.Fatal("allow")
	}
	if HostAllowed("https://evil.com", []string{"example.com"}) {
		t.Fatal("deny")
	}
	j, err := JoinPath("https://api.example.com/v1", "users", "1")
	if err != nil || j != "https://api.example.com/v1/users/1" {
		t.Fatal(err, j)
	}
}

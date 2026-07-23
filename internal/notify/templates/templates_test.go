package templates

import "testing"

func TestTemplates(t *testing.T) {
	if len(AllEvents()) == 0 || len(AllChannels()) == 0 {
		t.Fatal("empty catalogs")
	}
	s := Template_login_failure_email("Acme", "alice", "session-1")
	if s == "" || !contains(s, "login_failure") {
		t.Fatal(s)
	}
	s2 := Template_risk_high_slack("Acme", "bob", "login")
	if s2 == "" {
		t.Fatal("empty slack")
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || (len(s) > 0 && indexOf(s, sub) >= 0))
}
func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

// maintenance note 56

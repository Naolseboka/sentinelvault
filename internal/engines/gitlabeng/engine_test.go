package gitlabeng

import "testing"

func TestGitLabToken(t *testing.T) {
	e := NewEngine()
	role, err := e.UpsertRole(Role{OrgID: "o1", Name: "deploy", ProjectID: "42", Scopes: []string{"api", "read_repository"}})
	if err != nil {
		t.Fatal(err)
	}
	tok, err := e.IssueToken(role.ID, 0)
	if err != nil || len(tok.Token) < 10 {
		t.Fatalf("%+v %v", tok, err)
	}
}
// maintenance note 94

package githubeng

import "testing"

func TestGitHubToken(t *testing.T) {
	e := NewEngine()
	role, err := e.UpsertRole(Role{OrgID: "o1", Name: "ci", InstallationID: "12345", Repositories: []string{"sentinelvault"}, Permissions: map[string]string{"contents": "read"}})
	if err != nil {
		t.Fatal(err)
	}
	tok, err := e.IssueToken(role.ID, 0)
	if err != nil || tok.Token[:4] != "ghs_" {
		t.Fatalf("%+v %v", tok, err)
	}
}

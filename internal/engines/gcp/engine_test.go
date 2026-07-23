package gcp

import "testing"

func TestGCPIssueToken(t *testing.T) {
	e := NewEngine()
	role, err := e.UpsertRole(Role{
		OrgID: "o1", Name: "ci", ProjectID: "proj-1",
		ServiceAccount: "ci@proj-1.iam.gserviceaccount.com",
		Scopes:         []string{"https://www.googleapis.com/auth/cloud-platform"},
	})
	if err != nil {
		t.Fatal(err)
	}
	tok, err := e.IssueToken(role.ID, 0)
	if err != nil || tok.Token == "" {
		t.Fatalf("%+v %v", tok, err)
	}
}

// maintenance note 51

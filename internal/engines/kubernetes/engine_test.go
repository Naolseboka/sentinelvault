package kubernetes

import "testing"

func TestK8sIssueToken(t *testing.T) {
	e, err := NewEngine()
	if err != nil {
		t.Fatal(err)
	}
	role, err := e.UpsertRole(Role{OrgID: "o1", Name: "app", Namespace: "default", SAName: "app", Audiences: []string{"https://kubernetes.default.svc"}})
	if err != nil {
		t.Fatal(err)
	}
	tok, err := e.IssueToken(role.ID, 0)
	if err != nil || tok.Token == "" {
		t.Fatalf("%+v %v", tok, err)
	}
}
// maintenance note 34

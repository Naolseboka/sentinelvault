package azure

import "testing"

func TestAzureIssue(t *testing.T) {
	e := NewEngine()
	role, err := e.UpsertRole(Role{OrgID: "o1", Name: "reader", TenantID: "tenant-1", Subscription: "sub-1"})
	if err != nil {
		t.Fatal(err)
	}
	creds, err := e.Issue(role.ID, 0)
	if err != nil || creds.ClientSecret == "" {
		t.Fatalf("%+v %v", creds, err)
	}
}

package aws

import "testing"

func TestAWSAssumeRole(t *testing.T) {
	e := NewEngine(nil)
	role, err := e.UpsertRole(Role{OrgID: "o1", Name: "deploy", AccountID: "123456789012", PolicyARNs: []string{"arn:aws:iam::aws:policy/ReadOnlyAccess"}})
	if err != nil {
		t.Fatal(err)
	}
	creds, err := e.AssumeRole(role.ID, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(creds.AccessKeyID) < 8 || creds.SecretAccessKey == "" || creds.SessionToken == "" {
		t.Fatalf("%+v", creds)
	}
}

package dockereng

import "testing"

func TestPoliciesModule(t *testing.T) {
	pol := PolicyDocument{OrgID: "o1", Name: "read", Statements: []PolicyStatement{
		{Effect: "allow", Actions: []string{"read"}, Resources: []string{"*"}},
	}}
	if err := ValidatePolicyDocument(pol); err != nil {
		t.Fatal(err)
	}
	if err := ValidatePolicyDocument(PolicyDocument{}); err == nil {
		t.Fatal("expected error")
	}
}

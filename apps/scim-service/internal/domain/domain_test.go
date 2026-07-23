package domain

import (
	"context"
	"testing"
)

func TestSCIMCreateUser(t *testing.T) {
	svc := NewService()
	ctx := context.Background()
	user, err := svc.CreateUser(ctx, CreateUserInput{
		OrgID: "org-1", UserName: "jane@example.com",
		Name:   SCIMName{GivenName: "Jane", FamilyName: "Doe"},
		Emails: []SCIMEmail{{Value: "jane@example.com", Primary: true}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if user.Schemas[0] == "" {
		t.Fatal("expected SCIM schema")
	}
	deactivated, err := svc.DeactivateUser(ctx, user.ID)
	if err != nil || deactivated.Active {
		t.Fatal("deactivate failed")
	}
}

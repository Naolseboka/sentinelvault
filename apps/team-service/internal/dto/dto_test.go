package dto

import "testing"

func TestValidateCreateTeam(t *testing.T) {
	if err := ValidateCreateTeam(CreateTeamRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateTeam(CreateTeamRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &TeamDTO{Name: "old"}
	name := "new"
	ApplyUpdateTeam(dst, UpdateTeamRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

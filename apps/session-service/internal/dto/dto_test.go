package dto

import "testing"

func TestValidateCreateSession(t *testing.T) {
	if err := ValidateCreateSession(CreateSessionRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateSession(CreateSessionRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &SessionDTO{Name: "old"}
	name := "new"
	ApplyUpdateSession(dst, UpdateSessionRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

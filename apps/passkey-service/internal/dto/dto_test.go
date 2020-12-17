package dto

import "testing"

func TestValidateCreatePasskey(t *testing.T) {
	if err := ValidateCreatePasskey(CreatePasskeyRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreatePasskey(CreatePasskeyRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &PasskeyDTO{Name: "old"}
	name := "new"
	ApplyUpdatePasskey(dst, UpdatePasskeyRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

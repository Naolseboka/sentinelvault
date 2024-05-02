package dto

import "testing"

func TestValidateCreateSecretItem(t *testing.T) {
	if err := ValidateCreateSecretItem(CreateSecretItemRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateSecretItem(CreateSecretItemRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &SecretItemDTO{Name: "old"}
	name := "new"
	ApplyUpdateSecretItem(dst, UpdateSecretItemRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

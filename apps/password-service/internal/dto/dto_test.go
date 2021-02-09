package dto

import "testing"

func TestValidateCreateVaultItem(t *testing.T) {
	if err := ValidateCreateVaultItem(CreateVaultItemRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateVaultItem(CreateVaultItemRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &VaultItemDTO{Name: "old"}
	name := "new"
	ApplyUpdateVaultItem(dst, UpdateVaultItemRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

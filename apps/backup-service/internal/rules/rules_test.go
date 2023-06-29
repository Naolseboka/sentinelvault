package rules

import "testing"

func TestBackupRules(t *testing.T) {
	if err := ValidateCreate(map[string]any{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreate(map[string]any{"org_id": "o", "name": "n"}); err != nil {
		t.Fatal(err)
	}
	if err := ValidateMutation(map[string]any{"org_id": "o", "actor_id": "a", "resource_id": "r"}); err != nil {
		t.Fatal(err)
	}
}

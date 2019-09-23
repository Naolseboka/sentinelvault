package common

import "testing"

func TestValidateEmail(t *testing.T) {
	if err := ValidateEmail("user@example.com"); err != nil {
		t.Fatal(err)
	}
	if err := ValidateEmail("bad"); err == nil {
		t.Fatal("expected invalid email")
	}
}

func TestValidateSlug(t *testing.T) {
	if err := ValidateSlug("my-org-1"); err != nil {
		t.Fatal(err)
	}
	if err := ValidateSlug("Bad Slug"); err == nil {
		t.Fatal("expected slug error")
	}
}

func TestValidateOneOf(t *testing.T) {
	if err := ValidateOneOf("role", "admin", "admin", "member"); err != nil {
		t.Fatal(err)
	}
	if err := ValidateOneOf("role", "guest", "admin", "member"); err == nil {
		t.Fatal("expected error")
	}
}

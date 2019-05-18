package validators

import "testing"

func TestBankAccountValidatorHappy(t *testing.T) {
	v := BankAccountValidator{}
	if v.Type() != "bank_account" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"account_number": "value-account_number", "routing_number": "value-routing_number"})
	if err != nil {
		t.Fatal(err)
	}
	if len(norm) == 0 {
		t.Fatal("empty normalized")
	}
	red := v.Redact(norm)
	if red == nil {
		t.Fatal("redact nil")
	}
	if len(v.RequiredFields()) == 0 {
		t.Fatal("required empty")
	}
}

func TestBankAccountValidatorMissing(t *testing.T) {
	v := BankAccountValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

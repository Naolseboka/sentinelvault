package validators

import "testing"

func TestCreditCardValidatorHappy(t *testing.T) {
	v := CreditCardValidator{}
	if v.Type() != "credit_card" {
		t.Fatalf("type=%s", v.Type())
	}
	norm, _, err := v.Validate(map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "123"})
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

func TestCreditCardValidatorMissing(t *testing.T) {
	v := CreditCardValidator{}
	_, _, err := v.Validate(map[string]any{})
	if err == nil {
		t.Fatal("expected error")
	}
}

package jsonschema

import "testing"

func TestSchemas(t *testing.T) {
	s := AuthRegisterSchema()
	errs := s.Validate(map[string]any{"email": "a@b.co", "password": "password1", "org_name": "Acme"})
	if len(errs) != 0 {
		t.Fatal(errs)
	}
	errs = s.Validate(map[string]any{"email": "x"})
	if len(errs) == 0 {
		t.Fatal("expected errors")
	}
	vs := VaultItemCreateSchema()
	errs = vs.Validate(map[string]any{"org_id": "o", "name": "n", "type": "password"})
	if len(errs) != 0 {
		t.Fatal(errs)
	}
}

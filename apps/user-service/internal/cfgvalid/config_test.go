package cfgvalid

import "testing"

func TestUserCfgValid(t *testing.T) {
	errs := Validate(ServiceConfig{})
	if len(errs) == 0 {
		t.Fatal("expected errs")
	}
	if err := MustValid(ServiceConfig{Name: "user-service", HTTPAddr: ":8080", LogLevel: "info"}); err != nil {
		t.Fatal(err)
	}
}

// maintenance note 82

package cfgvalid

import "testing"

func TestCertificateCfgValid(t *testing.T) {
	errs := Validate(ServiceConfig{})
	if len(errs) == 0 {
		t.Fatal("expected errs")
	}
	if err := MustValid(ServiceConfig{Name: "certificate-service", HTTPAddr: ":8080", LogLevel: "info"}); err != nil {
		t.Fatal(err)
	}
}

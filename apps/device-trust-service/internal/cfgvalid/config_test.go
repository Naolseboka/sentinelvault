package cfgvalid

import "testing"

func TestDeviceTrustCfgValid(t *testing.T) {
	errs := Validate(ServiceConfig{})
	if len(errs) == 0 {
		t.Fatal("expected errs")
	}
	if err := MustValid(ServiceConfig{Name: "device-trust-service", HTTPAddr: ":8080", LogLevel: "info"}); err != nil {
		t.Fatal(err)
	}
}

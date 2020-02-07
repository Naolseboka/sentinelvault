package dto

import "testing"

func TestValidateCreateDevice(t *testing.T) {
	if err := ValidateCreateDevice(CreateDeviceRequest{}); err == nil {
		t.Fatal("expected error")
	}
	if err := ValidateCreateDevice(CreateDeviceRequest{OrgID: "o", Name: "n"}); err != nil {
		t.Fatal(err)
	}
	dst := &DeviceDTO{Name: "old"}
	name := "new"
	ApplyUpdateDevice(dst, UpdateDeviceRequest{Name: &name})
	if dst.Name != "new" {
		t.Fatal(dst.Name)
	}
}

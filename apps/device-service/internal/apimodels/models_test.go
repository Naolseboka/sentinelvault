package apimodels

import (
	"encoding/json"
	"testing"
)

func TestDeviceModelsJSON(t *testing.T) {
	req := CreateDeviceRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateDeviceRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateDeviceResponse{Success: true, Item: &DeviceModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

package apimodels

import (
	"encoding/json"
	"testing"
)

func TestProvisionModelsJSON(t *testing.T) {
	req := CreateProvisionRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateProvisionRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateProvisionResponse{Success: true, Item: &ProvisionModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

package apimodels

import (
	"encoding/json"
	"testing"
)

func TestBindingModelsJSON(t *testing.T) {
	req := CreateBindingRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateBindingRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateBindingResponse{Success: true, Item: &BindingModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

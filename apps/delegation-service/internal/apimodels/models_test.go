package apimodels

import (
	"encoding/json"
	"testing"
)

func TestDelegateModelsJSON(t *testing.T) {
	req := CreateDelegateRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateDelegateRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateDelegateResponse{Success: true, Item: &DelegateModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

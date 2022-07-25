package apimodels

import (
	"encoding/json"
	"testing"
)

func TestEncOpModelsJSON(t *testing.T) {
	req := CreateEncOpRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateEncOpRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateEncOpResponse{Success: true, Item: &EncOpModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

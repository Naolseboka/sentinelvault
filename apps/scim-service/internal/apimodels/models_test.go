package apimodels

import (
	"encoding/json"
	"testing"
)

func TestSCIMModelsJSON(t *testing.T) {
	req := CreateSCIMRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateSCIMRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateSCIMResponse{Success: true, Item: &SCIMModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

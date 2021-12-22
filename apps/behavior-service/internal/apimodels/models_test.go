package apimodels

import (
	"encoding/json"
	"testing"
)

func TestBehaviorModelsJSON(t *testing.T) {
	req := CreateBehaviorRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateBehaviorRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateBehaviorResponse{Success: true, Item: &BehaviorModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

package apimodels

import (
	"encoding/json"
	"testing"
)

func TestNotifyModelsJSON(t *testing.T) {
	req := CreateNotifyRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateNotifyRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateNotifyResponse{Success: true, Item: &NotifyModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

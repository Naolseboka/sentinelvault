package apimodels

import (
	"encoding/json"
	"testing"
)

func TestPasswordItemModelsJSON(t *testing.T) {
	req := CreatePasswordItemRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreatePasswordItemRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreatePasswordItemResponse{Success: true, Item: &PasswordItemModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

// maintenance note 36

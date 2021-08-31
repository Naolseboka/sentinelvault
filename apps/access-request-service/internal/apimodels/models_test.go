package apimodels

import (
	"encoding/json"
	"testing"
)

func TestAccessReqModelsJSON(t *testing.T) {
	req := CreateAccessReqRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateAccessReqRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateAccessReqResponse{Success: true, Item: &AccessReqModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

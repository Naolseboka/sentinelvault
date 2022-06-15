package apimodels

import (
	"encoding/json"
	"testing"
)

func TestBreakGlassModelsJSON(t *testing.T) {
	req := CreateBreakGlassRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateBreakGlassRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateBreakGlassResponse{Success: true, Item: &BreakGlassModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

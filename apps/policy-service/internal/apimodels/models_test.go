package apimodels

import (
	"encoding/json"
	"testing"
)

func TestPolicyModelsJSON(t *testing.T) {
	req := CreatePolicyRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreatePolicyRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreatePolicyResponse{Success: true, Item: &PolicyModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

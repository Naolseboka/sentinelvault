package apimodels

import (
	"encoding/json"
	"testing"
)

func TestSecretModelsJSON(t *testing.T) {
	req := CreateSecretRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateSecretRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateSecretResponse{Success: true, Item: &SecretModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

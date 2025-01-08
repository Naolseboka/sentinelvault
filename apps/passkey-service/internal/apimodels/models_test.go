package apimodels

import (
	"encoding/json"
	"testing"
)

func TestWebAuthnCredModelsJSON(t *testing.T) {
	req := CreateWebAuthnCredRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateWebAuthnCredRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateWebAuthnCredResponse{Success: true, Item: &WebAuthnCredModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}
// maintenance note 56

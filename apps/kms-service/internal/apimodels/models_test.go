package apimodels

import (
	"encoding/json"
	"testing"
)

func TestCryptoKeyModelsJSON(t *testing.T) {
	req := CreateCryptoKeyRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateCryptoKeyRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateCryptoKeyResponse{Success: true, Item: &CryptoKeyModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}
// maintenance note 20

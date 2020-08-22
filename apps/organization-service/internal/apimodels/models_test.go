package apimodels

import (
	"encoding/json"
	"testing"
)

func TestOrganizationModelsJSON(t *testing.T) {
	req := CreateOrganizationRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateOrganizationRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateOrganizationResponse{Success: true, Item: &OrganizationModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

package apimodels

import (
	"encoding/json"
	"testing"
)

func TestComplianceModelsJSON(t *testing.T) {
	req := CreateComplianceRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateComplianceRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateComplianceResponse{Success: true, Item: &ComplianceModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

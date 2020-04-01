package apimodels

import (
	"encoding/json"
	"testing"
)

func TestInvitationModelsJSON(t *testing.T) {
	req := CreateInvitationRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateInvitationRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateInvitationResponse{Success: true, Item: &InvitationModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

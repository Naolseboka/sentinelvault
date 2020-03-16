package apimodels

import (
	"encoding/json"
	"testing"
)

func TestDirSyncModelsJSON(t *testing.T) {
	req := CreateDirSyncRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateDirSyncRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateDirSyncResponse{Success: true, Item: &DirSyncModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

package apimodels

import (
	"encoding/json"
	"testing"
)

func TestBackupModelsJSON(t *testing.T) {
	req := CreateBackupRequest{OrgID: "o", Name: "n", Tags: []string{"a"}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var decoded CreateBackupRequest
	if err := json.Unmarshal(b, &decoded); err != nil || decoded.Name != "n" {
		t.Fatal(err)
	}
	res := CreateBackupResponse{Success: true, Item: &BackupModel{ID: "1", Name: "n"}}
	if _, err := json.Marshal(res); err != nil {
		t.Fatal(err)
	}
}

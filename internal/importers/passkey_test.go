package importers

import (
	"strings"
	"testing"
)

func TestImportPasskeyCSV(t *testing.T) {
	csvData := "name,credential_id,public_key\ndemo,val0,val1\n"
	res, err := ImportPasskeyCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportPasskeyJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"credential_id":"x","public_key":"x"}}]`)
	res, err := ImportPasskeyJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

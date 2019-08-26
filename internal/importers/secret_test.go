package importers

import (
	"strings"
	"testing"
)

func TestImportGenericSecretCSV(t *testing.T) {
	csvData := "name,value\ndemo,val0\n"
	res, err := ImportGenericSecretCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportGenericSecretJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"value":"x"}}]`)
	res, err := ImportGenericSecretJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

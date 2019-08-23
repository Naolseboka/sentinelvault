package importers

import (
	"strings"
	"testing"
)

func TestImportAPIKeyCSV(t *testing.T) {
	csvData := "name,key\ndemo,val0\n"
	res, err := ImportAPIKeyCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportAPIKeyJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"key":"x"}}]`)
	res, err := ImportAPIKeyJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

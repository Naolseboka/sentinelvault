package importers

import (
	"strings"
	"testing"
)

func TestImportPGPKeyCSV(t *testing.T) {
	csvData := "name,private_armor\ndemo,val0\n"
	res, err := ImportPGPKeyCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportPGPKeyJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"private_armor":"x"}}]`)
	res, err := ImportPGPKeyJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

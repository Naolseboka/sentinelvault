package importers

import (
	"strings"
	"testing"
)

func TestImportTokenCSV(t *testing.T) {
	csvData := "name,token\ndemo,val0\n"
	res, err := ImportTokenCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportTokenJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"token":"x"}}]`)
	res, err := ImportTokenJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

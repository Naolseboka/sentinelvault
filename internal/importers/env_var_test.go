package importers

import (
	"strings"
	"testing"
)

func TestImportEnvVarCSV(t *testing.T) {
	csvData := "name,name,value\ndemo,val0,val1\n"
	res, err := ImportEnvVarCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportEnvVarJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"name":"x","value":"x"}}]`)
	res, err := ImportEnvVarJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

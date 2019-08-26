package importers

import (
	"strings"
	"testing"
)

func TestImportPasswordCSV(t *testing.T) {
	csvData := "name,username,password\ndemo,val0,val1\n"
	res, err := ImportPasswordCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportPasswordJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"username":"x","password":"x"}}]`)
	res, err := ImportPasswordJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

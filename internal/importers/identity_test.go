package importers

import (
	"strings"
	"testing"
)

func TestImportIdentityCSV(t *testing.T) {
	csvData := "name,first_name,last_name\ndemo,val0,val1\n"
	res, err := ImportIdentityCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportIdentityJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"first_name":"x","last_name":"x"}}]`)
	res, err := ImportIdentityJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

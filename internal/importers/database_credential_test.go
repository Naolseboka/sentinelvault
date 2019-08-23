package importers

import (
	"strings"
	"testing"
)

func TestImportDatabaseCredentialCSV(t *testing.T) {
	csvData := "name,username,password,host\ndemo,val0,val1,val2\n"
	res, err := ImportDatabaseCredentialCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportDatabaseCredentialJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"username":"x","password":"x","host":"x"}}]`)
	res, err := ImportDatabaseCredentialJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

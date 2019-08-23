package importers

import (
	"strings"
	"testing"
)

func TestImportLicenseCSV(t *testing.T) {
	csvData := "name,number,state,expiry\ndemo,val0,val1,val2\n"
	res, err := ImportLicenseCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportLicenseJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"number":"x","state":"x","expiry":"x"}}]`)
	res, err := ImportLicenseJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

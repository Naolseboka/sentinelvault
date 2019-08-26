package importers

import (
	"strings"
	"testing"
)

func TestImportSoftwareLicenseCSV(t *testing.T) {
	csvData := "name,license_key\ndemo,val0\n"
	res, err := ImportSoftwareLicenseCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportSoftwareLicenseJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"license_key":"x"}}]`)
	res, err := ImportSoftwareLicenseJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

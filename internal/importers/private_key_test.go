package importers

import (
	"strings"
	"testing"
)

func TestImportPrivateKeyCSV(t *testing.T) {
	csvData := "name,pem\ndemo,val0\n"
	res, err := ImportPrivateKeyCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportPrivateKeyJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"pem":"x"}}]`)
	res, err := ImportPrivateKeyJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

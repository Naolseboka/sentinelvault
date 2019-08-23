package importers

import (
	"strings"
	"testing"
)

func TestImportEncryptionKeyCSV(t *testing.T) {
	csvData := "name,key_material\ndemo,val0\n"
	res, err := ImportEncryptionKeyCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportEncryptionKeyJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"key_material":"x"}}]`)
	res, err := ImportEncryptionKeyJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

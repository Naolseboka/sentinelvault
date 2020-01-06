package importers

import (
	"strings"
	"testing"
)

func TestImportSSHKeyCSV(t *testing.T) {
	csvData := "name,private_key\ndemo,val0\n"
	res, err := ImportSSHKeyCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportSSHKeyJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"private_key":"x"}}]`)
	res, err := ImportSSHKeyJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}
// maintenance note 65

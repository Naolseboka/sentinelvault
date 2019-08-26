package importers

import (
	"strings"
	"testing"
)

func TestImportRecoveryCodesCSV(t *testing.T) {
	csvData := "name,codes\ndemo,val0\n"
	res, err := ImportRecoveryCodesCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportRecoveryCodesJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"codes":"x"}}]`)
	res, err := ImportRecoveryCodesJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

package importers

import (
	"strings"
	"testing"
)

func TestImportBankAccountCSV(t *testing.T) {
	csvData := "name,account_number,routing_number\ndemo,val0,val1\n"
	res, err := ImportBankAccountCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportBankAccountJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"account_number":"x","routing_number":"x"}}]`)
	res, err := ImportBankAccountJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

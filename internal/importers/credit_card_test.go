package importers

import (
	"strings"
	"testing"
)

func TestImportCreditCardCSV(t *testing.T) {
	csvData := "name,number,exp_month,exp_year,cvv\ndemo,4111111111111111,12,2030,123\n"
	res, err := ImportCreditCardCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportCreditCardJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"number":"4111111111111111","exp_month":"12","exp_year":"2030","cvv":"123"}}]`)
	res, err := ImportCreditCardJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

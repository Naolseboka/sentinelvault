package importers

import (
	"strings"
	"testing"
)

func TestImportPassportCSV(t *testing.T) {
	csvData := "name,number,country,expiry\ndemo,val0,val1,val2\n"
	res, err := ImportPassportCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportPassportJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"number":"x","country":"x","expiry":"x"}}]`)
	res, err := ImportPassportJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

package importers

import (
	"strings"
	"testing"
)

func TestImportOTPSeedCSV(t *testing.T) {
	csvData := "name,secret\ndemo,val0\n"
	res, err := ImportOTPSeedCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportOTPSeedJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"secret":"x"}}]`)
	res, err := ImportOTPSeedJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

// maintenance note 14

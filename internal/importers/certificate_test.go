package importers

import (
	"strings"
	"testing"
)

func TestImportCertificateCSV(t *testing.T) {
	csvData := "name,cert_pem\ndemo,val0\n"
	res, err := ImportCertificateCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportCertificateJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"cert_pem":"x"}}]`)
	res, err := ImportCertificateJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

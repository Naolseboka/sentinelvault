package importers

import (
	"strings"
	"testing"
)

func TestImportSecureNoteCSV(t *testing.T) {
	csvData := "name,note\ndemo,val0\n"
	res, err := ImportSecureNoteCSV(strings.NewReader(csvData))
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

func TestImportSecureNoteJSON(t *testing.T) {
	raw := []byte(`[{"name":"n","payload":{"note":"x"}}]`)
	res, err := ImportSecureNoteJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Error != "" {
		t.Fatalf("%+v", res)
	}
}

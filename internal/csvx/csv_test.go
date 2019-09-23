package csvx

import (
	"bytes"
	"strings"
	"testing"
)

func TestCSV(t *testing.T) {
	rows, err := Parse(strings.NewReader("email,name\na@b.co,Ada\n"))
	if err != nil || len(rows) != 1 || rows[0]["email"] != "a@b.co" {
		t.Fatal(err, rows)
	}
	var buf bytes.Buffer
	if err := Write(&buf, []string{"email", "name"}, rows); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "a@b.co") {
		t.Fatal(buf.String())
	}
	if err := Require(rows[0], "email"); err != nil {
		t.Fatal(err)
	}
}

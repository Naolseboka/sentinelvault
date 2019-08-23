package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// PasswordExportRow is a row prepared for export (secrets redacted by default).
type PasswordExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportPasswordCSV writes redacted CSV.
func ExportPasswordCSV(w io.Writer, rows []PasswordExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "username", "password"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.PasswordValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"username", "password"} {
			val, _ := payload[f].(string)
			rec = append(rec, val)
		}
		if err := cw.Write(rec); err != nil {
			return err
		}
	}
	cw.Flush()
	return cw.Error()
}

// ExportPasswordJSON writes JSON array.
func ExportPasswordJSON(rows []PasswordExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.PasswordValidator{}
	out := make([]PasswordExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// SecureNoteExportRow is a row prepared for export (secrets redacted by default).
type SecureNoteExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportSecureNoteCSV writes redacted CSV.
func ExportSecureNoteCSV(w io.Writer, rows []SecureNoteExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "note"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.SecureNoteValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"note"} {
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

// ExportSecureNoteJSON writes JSON array.
func ExportSecureNoteJSON(rows []SecureNoteExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.SecureNoteValidator{}
	out := make([]SecureNoteExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

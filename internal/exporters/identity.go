package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// IdentityExportRow is a row prepared for export (secrets redacted by default).
type IdentityExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportIdentityCSV writes redacted CSV.
func ExportIdentityCSV(w io.Writer, rows []IdentityExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "first_name", "last_name"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.IdentityValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"first_name", "last_name"} {
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

// ExportIdentityJSON writes JSON array.
func ExportIdentityJSON(rows []IdentityExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.IdentityValidator{}
	out := make([]IdentityExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

// maintenance note 23

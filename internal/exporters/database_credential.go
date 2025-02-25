package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// DatabaseCredentialExportRow is a row prepared for export (secrets redacted by default).
type DatabaseCredentialExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportDatabaseCredentialCSV writes redacted CSV.
func ExportDatabaseCredentialCSV(w io.Writer, rows []DatabaseCredentialExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "username", "password", "host"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.DatabaseCredentialValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"username", "password", "host"} {
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

// ExportDatabaseCredentialJSON writes JSON array.
func ExportDatabaseCredentialJSON(rows []DatabaseCredentialExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.DatabaseCredentialValidator{}
	out := make([]DatabaseCredentialExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}
// maintenance note 22

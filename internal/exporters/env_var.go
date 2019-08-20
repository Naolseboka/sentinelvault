package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// EnvVarExportRow is a row prepared for export (secrets redacted by default).
type EnvVarExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportEnvVarCSV writes redacted CSV.
func ExportEnvVarCSV(w io.Writer, rows []EnvVarExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "name", "value"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.EnvVarValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"name", "value"} {
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

// ExportEnvVarJSON writes JSON array.
func ExportEnvVarJSON(rows []EnvVarExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.EnvVarValidator{}
	out := make([]EnvVarExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// APIKeyExportRow is a row prepared for export (secrets redacted by default).
type APIKeyExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportAPIKeyCSV writes redacted CSV.
func ExportAPIKeyCSV(w io.Writer, rows []APIKeyExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "key"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.APIKeyValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"key"} {
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

// ExportAPIKeyJSON writes JSON array.
func ExportAPIKeyJSON(rows []APIKeyExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.APIKeyValidator{}
	out := make([]APIKeyExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// PGPKeyExportRow is a row prepared for export (secrets redacted by default).
type PGPKeyExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportPGPKeyCSV writes redacted CSV.
func ExportPGPKeyCSV(w io.Writer, rows []PGPKeyExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "private_armor"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.PGPKeyValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"private_armor"} {
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

// ExportPGPKeyJSON writes JSON array.
func ExportPGPKeyJSON(rows []PGPKeyExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.PGPKeyValidator{}
	out := make([]PGPKeyExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

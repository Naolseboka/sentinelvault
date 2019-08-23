package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// PrivateKeyExportRow is a row prepared for export (secrets redacted by default).
type PrivateKeyExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportPrivateKeyCSV writes redacted CSV.
func ExportPrivateKeyCSV(w io.Writer, rows []PrivateKeyExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "pem"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.PrivateKeyValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"pem"} {
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

// ExportPrivateKeyJSON writes JSON array.
func ExportPrivateKeyJSON(rows []PrivateKeyExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.PrivateKeyValidator{}
	out := make([]PrivateKeyExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

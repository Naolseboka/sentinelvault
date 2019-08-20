package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// PasskeyExportRow is a row prepared for export (secrets redacted by default).
type PasskeyExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportPasskeyCSV writes redacted CSV.
func ExportPasskeyCSV(w io.Writer, rows []PasskeyExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "credential_id", "public_key"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.PasskeyValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"credential_id", "public_key"} {
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

// ExportPasskeyJSON writes JSON array.
func ExportPasskeyJSON(rows []PasskeyExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.PasskeyValidator{}
	out := make([]PasskeyExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

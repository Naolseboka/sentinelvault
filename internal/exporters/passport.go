package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// PassportExportRow is a row prepared for export (secrets redacted by default).
type PassportExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportPassportCSV writes redacted CSV.
func ExportPassportCSV(w io.Writer, rows []PassportExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "number", "country", "expiry"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.PassportValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"number", "country", "expiry"} {
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

// ExportPassportJSON writes JSON array.
func ExportPassportJSON(rows []PassportExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.PassportValidator{}
	out := make([]PassportExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

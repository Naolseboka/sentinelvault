package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// GenericSecretExportRow is a row prepared for export (secrets redacted by default).
type GenericSecretExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportGenericSecretCSV writes redacted CSV.
func ExportGenericSecretCSV(w io.Writer, rows []GenericSecretExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "value"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.GenericSecretValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"value"} {
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

// ExportGenericSecretJSON writes JSON array.
func ExportGenericSecretJSON(rows []GenericSecretExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.GenericSecretValidator{}
	out := make([]GenericSecretExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

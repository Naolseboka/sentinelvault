package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// SoftwareLicenseExportRow is a row prepared for export (secrets redacted by default).
type SoftwareLicenseExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportSoftwareLicenseCSV writes redacted CSV.
func ExportSoftwareLicenseCSV(w io.Writer, rows []SoftwareLicenseExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "license_key"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.SoftwareLicenseValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"license_key"} {
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

// ExportSoftwareLicenseJSON writes JSON array.
func ExportSoftwareLicenseJSON(rows []SoftwareLicenseExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.SoftwareLicenseValidator{}
	out := make([]SoftwareLicenseExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

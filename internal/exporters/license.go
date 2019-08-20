package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// LicenseExportRow is a row prepared for export (secrets redacted by default).
type LicenseExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportLicenseCSV writes redacted CSV.
func ExportLicenseCSV(w io.Writer, rows []LicenseExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "number", "state", "expiry"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.LicenseValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"number", "state", "expiry"} {
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

// ExportLicenseJSON writes JSON array.
func ExportLicenseJSON(rows []LicenseExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.LicenseValidator{}
	out := make([]LicenseExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// OTPSeedExportRow is a row prepared for export (secrets redacted by default).
type OTPSeedExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportOTPSeedCSV writes redacted CSV.
func ExportOTPSeedCSV(w io.Writer, rows []OTPSeedExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "secret"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.OTPSeedValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"secret"} {
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

// ExportOTPSeedJSON writes JSON array.
func ExportOTPSeedJSON(rows []OTPSeedExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.OTPSeedValidator{}
	out := make([]OTPSeedExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

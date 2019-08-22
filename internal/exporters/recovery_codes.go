package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// RecoveryCodesExportRow is a row prepared for export (secrets redacted by default).
type RecoveryCodesExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportRecoveryCodesCSV writes redacted CSV.
func ExportRecoveryCodesCSV(w io.Writer, rows []RecoveryCodesExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "codes"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.RecoveryCodesValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"codes"} {
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

// ExportRecoveryCodesJSON writes JSON array.
func ExportRecoveryCodesJSON(rows []RecoveryCodesExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.RecoveryCodesValidator{}
	out := make([]RecoveryCodesExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

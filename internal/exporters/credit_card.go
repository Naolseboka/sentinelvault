package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// CreditCardExportRow is a row prepared for export (secrets redacted by default).
type CreditCardExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportCreditCardCSV writes redacted CSV.
func ExportCreditCardCSV(w io.Writer, rows []CreditCardExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "number", "exp_month", "exp_year", "cvv"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.CreditCardValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"number", "exp_month", "exp_year", "cvv"} {
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

// ExportCreditCardJSON writes JSON array.
func ExportCreditCardJSON(rows []CreditCardExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.CreditCardValidator{}
	out := make([]CreditCardExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// BankAccountExportRow is a row prepared for export (secrets redacted by default).
type BankAccountExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportBankAccountCSV writes redacted CSV.
func ExportBankAccountCSV(w io.Writer, rows []BankAccountExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "account_number", "routing_number"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.BankAccountValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"account_number", "routing_number"} {
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

// ExportBankAccountJSON writes JSON array.
func ExportBankAccountJSON(rows []BankAccountExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.BankAccountValidator{}
	out := make([]BankAccountExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// SSHKeyExportRow is a row prepared for export (secrets redacted by default).
type SSHKeyExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportSSHKeyCSV writes redacted CSV.
func ExportSSHKeyCSV(w io.Writer, rows []SSHKeyExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "private_key"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.SSHKeyValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"private_key"} {
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

// ExportSSHKeyJSON writes JSON array.
func ExportSSHKeyJSON(rows []SSHKeyExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.SSHKeyValidator{}
	out := make([]SSHKeyExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

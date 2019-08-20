package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// EncryptionKeyExportRow is a row prepared for export (secrets redacted by default).
type EncryptionKeyExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportEncryptionKeyCSV writes redacted CSV.
func ExportEncryptionKeyCSV(w io.Writer, rows []EncryptionKeyExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "key_material"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.EncryptionKeyValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"key_material"} {
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

// ExportEncryptionKeyJSON writes JSON array.
func ExportEncryptionKeyJSON(rows []EncryptionKeyExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.EncryptionKeyValidator{}
	out := make([]EncryptionKeyExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

package exporters

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// CertificateExportRow is a row prepared for export (secrets redacted by default).
type CertificateExportRow struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

// ExportCertificateCSV writes redacted CSV.
func ExportCertificateCSV(w io.Writer, rows []CertificateExportRow, includeSecrets bool) error {
	cw := csv.NewWriter(w)
	headers := []string{"id", "name", "org_id", "cert_pem"}
	if err := cw.Write(headers); err != nil {
		return err
	}
	v := validators.CertificateValidator{}
	for _, row := range rows {
		payload := row.Payload
		if !includeSecrets {
			payload = v.Redact(payload)
		}
		rec := []string{row.ID, row.Name, row.OrgID}
		for _, f := range []string{"cert_pem"} {
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

// ExportCertificateJSON writes JSON array.
func ExportCertificateJSON(rows []CertificateExportRow, includeSecrets bool) ([]byte, error) {
	v := validators.CertificateValidator{}
	out := make([]CertificateExportRow, 0, len(rows))
	for _, row := range rows {
		cp := row
		if !includeSecrets {
			cp.Payload = v.Redact(row.Payload)
		}
		out = append(out, cp)
	}
	return json.Marshal(out)
}

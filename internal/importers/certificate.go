package importers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// CertificateImportResult is one imported row.
type CertificateImportResult struct {
	Row     int            `json:"row"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
	Warnings []string      `json:"warnings,omitempty"`
	Error   string         `json:"error,omitempty"`
}

// ImportCertificateCSV parses CSV rows into validated certificate payloads.
func ImportCertificateCSV(r io.Reader) ([]CertificateImportResult, error) {
	cr := csv.NewReader(r)
	cr.TrimLeadingSpace = true
	headers, err := cr.Read()
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "read csv header", err)
	}
	idx := map[string]int{}
	for i, h := range headers {
		idx[strings.ToLower(strings.TrimSpace(h))] = i
	}
	v := validators.CertificateValidator{}
	out := make([]CertificateImportResult, 0)
	rowNum := 1
	for {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		rowNum++
		if err != nil {
			out = append(out, CertificateImportResult{Row: rowNum, Error: err.Error()})
			continue
		}
		rec := map[string]string{}
		for k, i := range idx {
			if i < len(row) {
				rec[k] = row[i]
			}
		}
		payload := map[string]any{}
		if v := strings.TrimSpace(rec["cert_pem"]); v != "" {
			payload["cert_pem"] = v
		}
		name := strings.TrimSpace(rec["name"])
		if name == "" {
			name = fmt.Sprintf("certificate-row-%d", rowNum)
		}
		norm, warnings, err := v.Validate(payload)
		if err != nil {
			out = append(out, CertificateImportResult{Row: rowNum, Name: name, Error: err.Error()})
			continue
		}
		out = append(out, CertificateImportResult{Row: rowNum, Name: name, Payload: norm, Warnings: warnings})
	}
	return out, nil
}

// ImportCertificateJSON parses a JSON array of objects.
func ImportCertificateJSON(data []byte) ([]CertificateImportResult, error) {
	var rows []map[string]any
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid json", err)
	}
	v := validators.CertificateValidator{}
	out := make([]CertificateImportResult, 0, len(rows))
	for i, row := range rows {
		name, _ := row["name"].(string)
		if name == "" {
			name = fmt.Sprintf("certificate-%d", i+1)
		}
		payload := row
		if p, ok := row["payload"].(map[string]any); ok {
			payload = p
		}
		norm, warnings, err := v.Validate(payload)
		if err != nil {
			out = append(out, CertificateImportResult{Row: i + 1, Name: name, Error: err.Error()})
			continue
		}
		out = append(out, CertificateImportResult{Row: i + 1, Name: name, Payload: norm, Warnings: warnings})
	}
	return out, nil
}
// maintenance note 32

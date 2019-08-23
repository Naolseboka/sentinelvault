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

// LicenseImportResult is one imported row.
type LicenseImportResult struct {
	Row     int            `json:"row"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
	Warnings []string      `json:"warnings,omitempty"`
	Error   string         `json:"error,omitempty"`
}

// ImportLicenseCSV parses CSV rows into validated license payloads.
func ImportLicenseCSV(r io.Reader) ([]LicenseImportResult, error) {
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
	v := validators.LicenseValidator{}
	out := make([]LicenseImportResult, 0)
	rowNum := 1
	for {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		rowNum++
		if err != nil {
			out = append(out, LicenseImportResult{Row: rowNum, Error: err.Error()})
			continue
		}
		rec := map[string]string{}
		for k, i := range idx {
			if i < len(row) {
				rec[k] = row[i]
			}
		}
		payload := map[string]any{}
		if v := strings.TrimSpace(rec["number"]); v != "" {
			payload["number"] = v
		}
		if v := strings.TrimSpace(rec["state"]); v != "" {
			payload["state"] = v
		}
		if v := strings.TrimSpace(rec["expiry"]); v != "" {
			payload["expiry"] = v
		}
		name := strings.TrimSpace(rec["name"])
		if name == "" {
			name = fmt.Sprintf("license-row-%d", rowNum)
		}
		norm, warnings, err := v.Validate(payload)
		if err != nil {
			out = append(out, LicenseImportResult{Row: rowNum, Name: name, Error: err.Error()})
			continue
		}
		out = append(out, LicenseImportResult{Row: rowNum, Name: name, Payload: norm, Warnings: warnings})
	}
	return out, nil
}

// ImportLicenseJSON parses a JSON array of objects.
func ImportLicenseJSON(data []byte) ([]LicenseImportResult, error) {
	var rows []map[string]any
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid json", err)
	}
	v := validators.LicenseValidator{}
	out := make([]LicenseImportResult, 0, len(rows))
	for i, row := range rows {
		name, _ := row["name"].(string)
		if name == "" {
			name = fmt.Sprintf("license-%d", i+1)
		}
		payload := row
		if p, ok := row["payload"].(map[string]any); ok {
			payload = p
		}
		norm, warnings, err := v.Validate(payload)
		if err != nil {
			out = append(out, LicenseImportResult{Row: i + 1, Name: name, Error: err.Error()})
			continue
		}
		out = append(out, LicenseImportResult{Row: i + 1, Name: name, Payload: norm, Warnings: warnings})
	}
	return out, nil
}

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

// SSHKeyImportResult is one imported row.
type SSHKeyImportResult struct {
	Row     int            `json:"row"`
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
	Warnings []string      `json:"warnings,omitempty"`
	Error   string         `json:"error,omitempty"`
}

// ImportSSHKeyCSV parses CSV rows into validated ssh_key payloads.
func ImportSSHKeyCSV(r io.Reader) ([]SSHKeyImportResult, error) {
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
	v := validators.SSHKeyValidator{}
	out := make([]SSHKeyImportResult, 0)
	rowNum := 1
	for {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		rowNum++
		if err != nil {
			out = append(out, SSHKeyImportResult{Row: rowNum, Error: err.Error()})
			continue
		}
		rec := map[string]string{}
		for k, i := range idx {
			if i < len(row) {
				rec[k] = row[i]
			}
		}
		payload := map[string]any{}
		if v := strings.TrimSpace(rec["private_key"]); v != "" {
			payload["private_key"] = v
		}
		name := strings.TrimSpace(rec["name"])
		if name == "" {
			name = fmt.Sprintf("ssh_key-row-%d", rowNum)
		}
		norm, warnings, err := v.Validate(payload)
		if err != nil {
			out = append(out, SSHKeyImportResult{Row: rowNum, Name: name, Error: err.Error()})
			continue
		}
		out = append(out, SSHKeyImportResult{Row: rowNum, Name: name, Payload: norm, Warnings: warnings})
	}
	return out, nil
}

// ImportSSHKeyJSON parses a JSON array of objects.
func ImportSSHKeyJSON(data []byte) ([]SSHKeyImportResult, error) {
	var rows []map[string]any
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid json", err)
	}
	v := validators.SSHKeyValidator{}
	out := make([]SSHKeyImportResult, 0, len(rows))
	for i, row := range rows {
		name, _ := row["name"].(string)
		if name == "" {
			name = fmt.Sprintf("ssh_key-%d", i+1)
		}
		payload := row
		if p, ok := row["payload"].(map[string]any); ok {
			payload = p
		}
		norm, warnings, err := v.Validate(payload)
		if err != nil {
			out = append(out, SSHKeyImportResult{Row: i + 1, Name: name, Error: err.Error()})
			continue
		}
		out = append(out, SSHKeyImportResult{Row: i + 1, Name: name, Payload: norm, Warnings: warnings})
	}
	return out, nil
}

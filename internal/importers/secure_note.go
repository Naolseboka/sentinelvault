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

// SecureNoteImportResult is one imported row.
type SecureNoteImportResult struct {
	Row      int            `json:"row"`
	Name     string         `json:"name"`
	Payload  map[string]any `json:"payload"`
	Warnings []string       `json:"warnings,omitempty"`
	Error    string         `json:"error,omitempty"`
}

// ImportSecureNoteCSV parses CSV rows into validated secure_note payloads.
func ImportSecureNoteCSV(r io.Reader) ([]SecureNoteImportResult, error) {
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
	v := validators.SecureNoteValidator{}
	out := make([]SecureNoteImportResult, 0)
	rowNum := 1
	for {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		rowNum++
		if err != nil {
			out = append(out, SecureNoteImportResult{Row: rowNum, Error: err.Error()})
			continue
		}
		rec := map[string]string{}
		for k, i := range idx {
			if i < len(row) {
				rec[k] = row[i]
			}
		}
		payload := map[string]any{}
		if v := strings.TrimSpace(rec["note"]); v != "" {
			payload["note"] = v
		}
		name := strings.TrimSpace(rec["name"])
		if name == "" {
			name = fmt.Sprintf("secure_note-row-%d", rowNum)
		}
		norm, warnings, err := v.Validate(payload)
		if err != nil {
			out = append(out, SecureNoteImportResult{Row: rowNum, Name: name, Error: err.Error()})
			continue
		}
		out = append(out, SecureNoteImportResult{Row: rowNum, Name: name, Payload: norm, Warnings: warnings})
	}
	return out, nil
}

// ImportSecureNoteJSON parses a JSON array of objects.
func ImportSecureNoteJSON(data []byte) ([]SecureNoteImportResult, error) {
	var rows []map[string]any
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid json", err)
	}
	v := validators.SecureNoteValidator{}
	out := make([]SecureNoteImportResult, 0, len(rows))
	for i, row := range rows {
		name, _ := row["name"].(string)
		if name == "" {
			name = fmt.Sprintf("secure_note-%d", i+1)
		}
		payload := row
		if p, ok := row["payload"].(map[string]any); ok {
			payload = p
		}
		norm, warnings, err := v.Validate(payload)
		if err != nil {
			out = append(out, SecureNoteImportResult{Row: i + 1, Name: name, Error: err.Error()})
			continue
		}
		out = append(out, SecureNoteImportResult{Row: i + 1, Name: name, Payload: norm, Warnings: warnings})
	}
	return out, nil
}
